package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
	}
	if body, err := ioutil.ReadAll(resp.Body); err == nil {

		var result map[string]interface{}

		json.Unmarshal([]byte(body), &result)
		switch API {
		case 1: // for the Fixer API
			if result["success"] == true {
				fmt.Println(result["rates"].(map[string]interface{})["USD"])
			} else {
				fmt.Println(result["error"].(map[string]interface{})["info"])
			}
		}
	} else {
		log.Fatal(err)
	}
}
