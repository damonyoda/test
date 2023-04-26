package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var hash string

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
