package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
  "github.com/gorilla/mux"
	"encoding/json"
	"net/url"
)

type OpenWeather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func GetWeather(city string) []byte{
	safeQuery := url.QueryEscape(city)

	url := fmt.Sprintf("https://community-open-weather-map.p.rapidapi.com/weather?q=%s", safeQuery)

	req, _ := http.NewRequest("GET", url, nil)


	req.Header.Add("x-rapidapi-host", "community-open-weather-map.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "197b9992b3mshf84e47cf0693477p123b73jsnb27522c04ca3")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	fmt.Println(string(body))
	return body

}

func hello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Welcome to my awseome site!</h1>")
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	fmt.Print("Enter City: ")   //Print function is used to display output in same line
  var city string
  fmt.Scanln(&city)                  // Take input from user
	
	openweather := OpenWeather{}
	respBody := GetWeather(city)
	err := json.Unmarshal(respBody, &openweather)
	if (err != nil) {
		fmt.Println("error occured during unmarshalling: ", err)
	}
	fmt.Println("openweather after: ", openweather)
	// fmt.Fprint(w, openweather)

	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	// r.HandleFunc("/weather", weatherHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
