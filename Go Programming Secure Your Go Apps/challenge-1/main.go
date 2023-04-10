package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Data struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	StatusWater string `json:"status_water"`
	StatusWind  string `json:"status_wind"`
}

func getStatus(value int, unit string) string {
	if strings.ToLower(unit) == "water" {
		if value < 5 {
			return "aman"
		} else if value >= 5 && value <= 8 {
			return "siaga"
		} else {
			return "bahaya"
		}
	} else if strings.ToLower(unit) == "wind" {
		if value < 6 {
			return "aman"
		} else if value >= 6 && value <= 15 {
			return "siaga"
		} else {
			return "bahaya"
		}
	} else {
		return ""
	}
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		data := Data{
			Water:       water,
			Wind:        wind,
			StatusWater: getStatus(water, "water"),
			StatusWind:  getStatus(wind, "wind"),
		}

		json_data, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
			continue
		}

		resp, err := http.Post(url, "application/json", strings.NewReader(string(json_data)))
		if err != nil {
			log.Fatalln(err)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(resp.Status)
		fmt.Println(string(body))
		fmt.Println("status water : ", data.StatusWater)
		fmt.Println("status wind : ", data.StatusWind)

		time.Sleep(15 * time.Second)
	}
}
