package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type Jso struct {
	Facts   []string `json:"facts"`
	Success bool     `json:"success"`
}

func getJson(url string) map[string]string {
	// Make HTTP GET request to fetch the JSON data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	var data map[string]string

	// Unmarshal the JSON into the struct
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	return data
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

}

func main() {
	fmt.Println("Hello")
	js := getJson("https://raw.githubusercontent.com/pui4/compack/main/data.json")
	fmt.Println(js["discord"])
	downloadFile("./poop.png", "https://gophercoding.com/img/logo-original.png")
}
