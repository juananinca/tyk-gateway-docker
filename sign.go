package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const SECRET_KEY = "posc_ndr75e4a20bbcfd4db5a9b0d2663a862e21"

func signRequest(req *http.Request) {
	refDate := "Mon, 02 Jan 2006 15:04:05 MST"

	// Prepare the request headers:
	tim := time.Now().Format(refDate)
	req.Header.Add("Date", tim)
	req.Header.Add("X-Test-1", "hello")
	req.Header.Add("X-Test-2", "world")

	// Prepare the signature to include those headers:
	signatureString := "(request-target): " + "get /your/path/goes/here"
	signatureString += "date: " + tim + "\n"
	signatureString += "x-test-1: " + "hello" + "\n"
	signatureString += "x-test-2: " + "world"

	// SHA1 Encode the signature
	HmacSecret := SECRET_KEY
	key := []byte(HmacSecret)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(signatureString))

	// Base64 and URL Encode the string
	sigString := base64.StdEncoding.EncodeToString(h.Sum(nil))
	encodedString := url.QueryEscape(sigString)

	// Add the header
	req.Header.Add("Authorization", fmt.Sprintf("Signature keyId=\"9876\",algorithm=\"hmac-sha1\",headers=\"(request-target) date x-test-1 x-test-2\",signature=\"%s\"", encodedString))
}

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080/animal/hello", nil)
	if err != nil {
		log.Fatal(err)
	}
	
	signRequest(req)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println("Body:", string(body))

}

