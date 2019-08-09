package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	j := 0
	for i := 1024; i <= 65535; i++ {
		// fmt.Println("W")
		go func() {
			fmt.Println("?")
			url := "http://localhost:8080/addServer"

			payload := strings.NewReader("{\"IP\":\"localhost:" + strconv.Itoa(i) + "\"}")

			req, _ := http.NewRequest("POST", url, payload)

			req.Header.Add("Content-Type", "application/json")

			res, _ := http.DefaultClient.Do(req)

			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			j++
			fmt.Println(res)
			fmt.Println(string(body))
		}()
	}
	fmt.Println(j)
	fmt.Println(65536 - 1024)
}
