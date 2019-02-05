package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://api.airvisual.com/v2/city?city=London&state=England&country=UK&key=QXK5ELcmwFpTb3Cce"
	method := "GET"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
