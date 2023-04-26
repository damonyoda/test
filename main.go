package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/hash", func(c *fiber.Ctx) error {
		time.Sleep(1 * time.Second)
		hash, err := generateRandomHash()
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error generating hash, %v", err))
		}

		return c.JSON(fiber.Map{"hash": hash})
	})

	app.Get("/check", func(c *fiber.Ctx) error {
		// var counter int
		for {
			// counter++
			// if counter > 50 {
			// 	// maximum attempts reached
			// 	fmt.Println(counter)
			// 	return c.JSON(fiber.Map{"message": "Maximum attempts reached"}) // to prevent infinite looping
			// }
			resp, err := http.Get("http://localhost:3000/hash")

			if err != nil {
				fmt.Println(err)
				continue
			}
			defer resp.Body.Close()

			var data map[string]string
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				fmt.Println(err)
				continue
			}

			hash := data["hash"]
			lastCharater := string(hash[len(hash)-1])
			isEven, err := checkLastCharacter(lastCharater)
			if err != nil {
				fmt.Println(fmt.Sprintf("%v is alphabet", string(lastCharater)))
				continue
			}
			if isEven {
				fmt.Println(fmt.Sprintf("%v is even Number. PASS", string(lastCharater)))
			} else {
				fmt.Println(fmt.Sprintf("%v is odd Number. !PASS", string(lastCharater)))
			}
			time.Sleep(1 * time.Second) // Wait for 1 second before making the next request
		}
	})

	app.Listen(":3000")
}

func generateRandomHash() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	h := sha256.New()
	_, err = h.Write(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func checkLastCharacter(last string) (bool, error) {
	i, err := strconv.Atoi(last)
	if err != nil {
		return false, err
	}
	return i%2 == 0, nil
}
