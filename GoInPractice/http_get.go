package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://ezequielaguilar.com.pa")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
