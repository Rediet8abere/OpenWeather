package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
  "github.com/gorilla/mux"
	"encoding/json"
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



// func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// openweather := OpenWeather{}
	// fmt.Println("openweather: ", openweather)
	// jsn, err := ioutil.ReadAll(r.Body)
	// fmt.Println("getting jsn: ", jsn)
	//
	// if err != nil {
	// 	log.Fatal("Error while reading r.Body: ", err)
	// }
	// err = json.Unmarshal(jsn, &openweather)
	// if err != nil {
	// 	log.Fatal("Decoding error: ", err)
	// }
	// log.Printf("Recived: %v\n", openweather)

// }
func getOpenWeather(body []byte) (*OpenWeather, error) {
	var s = new(OpenWeather)
	err := json.Unmarshal(body, &s)
	if (err != nil) {
		fmt.Println("error occured during unmarshalling: ", err)
	}
	return s, err
}

func hello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Welcome to my awseome site!</h1>")
}

func main() {
	GetRequest()

	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	// r.HandleFunc("/weather", weatherHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func GetRequest() {
	url := "https://community-open-weather-map.p.rapidapi.com/weather?q=London"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "community-open-weather-map.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "197b9992b3mshf84e47cf0693477p123b73jsnb27522c04ca3")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	fmt.Println(string(body))
	s, err:= getOpenWeather([]byte(body))
	if (err!=nil) {
		fmt.Println("error!!")
	}
	fmt.Println(s)
}
