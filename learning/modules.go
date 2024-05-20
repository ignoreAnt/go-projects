package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	baseURL := "https://cat-fact.herokuapp.com"
	facts := "/facts"

	//url := "https://jsonplaceholder.typicode.com/todos/1"
	url := baseURL + facts

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-My-Client", "Learning Go")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		fmt.Println(fmt.Sprintf("unexpected status: got %v", res.Status))
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Header.Get("Content-Type"))
	fmt.Println(res.Header)
	fmt.Println(string(resp))

	var data struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"_id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", data)
	fmt.Printf("%v\n", data)
	fmt.Println(data)
}
