package main

import (
	"fmt"
	"net/http"
)

type APIClient struct {
	baseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{baseURL: baseURL}
}

func (c *APIClient) FetchData(path string) (string, error) {
	url := c.baseURL + path
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buffer)
		fmt.Println("body bytes read", n)
		if err != nil {
			break
		}
		data = append(data, buffer[:n]...)
	}

	return string(data), nil
}

func main() {
	client := NewAPIClient("https://dummyjson.com")
	data, err := client.FetchData("/products/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Data:", data)
}
