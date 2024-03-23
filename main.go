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

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
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
	jso := Jso{}
	getJson("https://dog-api.kinduff.com/api/facts", &jso)
	fmt.Println(jso.Facts[0])
	downloadFile("./poop.png", "https://gophercoding.com/img/logo-original.png")
}
