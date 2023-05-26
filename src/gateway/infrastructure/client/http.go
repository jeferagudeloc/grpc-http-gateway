package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jeferagudeloc/grpc-http-gateway/src/gateway/domain"
)

type HttpClient struct{}

func (c *HttpClient) GetOrdersFromBackend() ([]domain.Order, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://server:8080/orders", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var orders []domain.Order
	err = json.Unmarshal(body, &orders)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return orders, nil
}

func (c *HttpClient) GetUsersFromBackend() ([]domain.User, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://server:8080/users", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var users []domain.User
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return users, nil
}
