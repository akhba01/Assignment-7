package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

type resPost struct {
	UserID int    `json:"id"`
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	var url = "https://jsonplaceholder.typicode.com/posts"
	var method = "POST"
	for {
		water := h8HelperRand.RandomInt(0, 100)
		wind := h8HelperRand.RandomInt(0, 100)
		userID := h8HelperRand.RandomInt(0, 100)

		dataRequest := map[string]interface{}{
			"userID": userID,
			"wind":   wind,
			"water":  water,
		}

		reqJSON, err := json.MarshalIndent(dataRequest, "", " ")

		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Request :")
		fmt.Println(string(reqJSON))

		req, err := http.NewRequest(method, url, bytes.NewBuffer(reqJSON))
		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Response :")
		fmt.Println(string(body))

		var ResPost resPost
		err = json.Unmarshal(body, &ResPost)
		if err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

		// print object water dan wind
		fmt.Println("print object water dan wind")
		ObjectWaterWind(ResPost.Water, ResPost.Wind)

		// print status water dan status wind
		fmt.Println("print status water dan status wind")
		StatusWaterWind(ResPost.Water, ResPost.Wind)

		time.Sleep(15 * time.Second)
	}
}

type WaterWind struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func ObjectWaterWind(water, wind int) {
	result := WaterWind{
		Water: water,
		Wind:  wind,
	}
	json, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(string(json))

}

func StatusWaterWind(water, wind int) {

	var statusWater string
	var statusWind string

	if water < 5 {
		statusWater = "aman"
	} else if water >= 6 && water <= 8 {
		statusWater = "siaga"
	} else {
		statusWater = "bahaya"
	}

	if wind < 6 {
		statusWind = "aman"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "siaga"
	} else {
		statusWind = "bahaya"
	}

	fmt.Printf("status water : %s \n", statusWater)
	fmt.Printf("status wind : %s \n", statusWind)
}
