package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* type Result struct {
	Success   bool
	Timestamp int
	Base      string
	Date      string
	Rates     map[string]float64
}

type Error struct {
	Success bool
	Error   struct {
		Code int
		Type string
		Info string
	}
} */

var apis map[int]string

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
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
			case 2: //for openweater API
				if result["main"] != nil {
					fmt.Println(result["main"].(map[string]interface{})["temp"])
				} else {
					fmt.Println(result["message"])
				}
			}
		} else {
			log.Fatal(err)
		}
	}
}

func main() {
	/* url := "http://data.fixer.io/api/latest?access_key=49fd52932dfdb84310a3dcfa77c25beb" */
	apis = make(map[int]string)

	apis[1] = "http://data.fixer.io/api/latest?access_key=" + "49fd52932dfdb84310a3dcfa77c25beb"

	apis[2] = "http://api.openweathermap.org/data/2.5/weather?" + "q=SINGAPORE&appid=7678666e1317103c103cf81584eebb14"

	fetchData(1)
	fetchData(2)
}
