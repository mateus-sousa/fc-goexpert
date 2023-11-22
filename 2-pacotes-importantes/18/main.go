package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	//reqBody := bytes.NewBuffer([]byte(`{"name": "Panda"`))
	//resp, err := c.Post("http://google.com", "application/json", reqBody)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//io.CopyBuffer(os.Stdout, resp.Body, nil)
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(body)
}
