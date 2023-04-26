import http from "k6/http";
import { sleep } from "k6";

export let options = {
  vus: 1, // number of virtual users
  duration: "20s", // duration of the test
};

export default function () {
  http.get("http://localhost:3000/check");
  sleep(1000); // wait for 1 second before sending the next request
}