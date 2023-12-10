package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nikhil133/duplo/src/models"
	"github.com/nikhil133/duplo/src/repositoryfactory"
	"github.com/nikhil133/duplo/src/repositoryinterface"
	"github.com/nikhil133/duplo/utils/constant"
	utils "github.com/nikhil133/duplo/utils/request"
	"github.com/sirupsen/logrus"
)

type WeatherService struct {
	Repo repositoryinterface.WeatherRepository
}

type meteoData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly    hourly  `json:"hourly"`
}
type hourly struct {
	Time        []string `json:"time"`
	WeatherCode []int    `json:"weather_code"`
}

func (ws WeatherService) GetForcast(req utils.Request) (*Forecast, error) {
	coordintate := models.Coordinate{}
	forecast := Forecast{}
	data, err := meteoRequest(req.Latitude, req.Longitude)
	if data == nil {
		return nil, err
	}
	for i, d := range data.Hourly.Time {
		w := Weather{}
		w.Time = d
		w.Weather = weatherConv(data.Hourly.WeatherCode[i])
		forecast.Forecast = append(forecast.Forecast, w)
	}
	coordintate.Latitude = req.Latitude
	coordintate.Longitude = req.Longitude
	ws.Repo = repositoryfactory.NewWeatherRepository()
	err = ws.Repo.Insert(&coordintate)
	forecast.Status.Code = http.StatusOK
	forecast.Status.Error = false
	forecast.Status.Message = "success"
	return &forecast, err
}

func (w WeatherService) GetCoordinates() ([]Coordinate, error) {
	coordinates := []Coordinate{}
	w.Repo = repositoryfactory.NewWeatherRepository()
	err := w.Repo.Get(&coordinates)
	return coordinates, err
}

func (w WeatherService) DeleteCoordinates(id uint) error {
	w.Repo = repositoryfactory.NewWeatherRepository()
	return w.Repo.Delete(id)
}

func meteoRequest(lat, long float64) (*meteoData, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, constant.METOAPI, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add(constant.QUERYNAME1, strconv.FormatFloat(long, 'f', -1, 64))
	q.Add(constant.QUERYNAME2, strconv.FormatFloat(lat, 'f', -1, 64))
	q.Add("hourly", "weather_code")
	req.URL.RawQuery = q.Encode()
	logrus.Info("Request ", req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	var data meteoData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &data, nil
}

func weatherConv(code int) string {
	m := map[int]string{
		0:  "Clear sky",
		1:  "Mainly clear",
		2:  "Partly cloudy",
		3:  "Overcast",
		45: "Fog",
		48: "Depositing rime fog",
		96: "Thunderstorm with slight hail",
		99: "Thunderstorm with heavy hail",
		95: "Thunderstorm",
		85: "Slight Snow showers",
		86: "Heavy Snow showers",
		80: "Slight Rain showers",
		81: "Moderate Rain showers",
		82: "Violent Rain showers",
		77: "Snow grains",
		71: "Slight Snow fall",
		73: "Moderate Snow fall",
		75: "Heavy intensity Snow fall",
		66: "Light Freezing Rain",
		67: "Heavy intensity Freezing Rain",
		61: "Slight Rain",
		63: "Moderate Rain",
		65: "Heavy intensity Rain",
		56: "Light Freezing Drizzle",
		57: "Dense intensity Freezing Drizzle",
		51: "Light Drizzle",
		53: "Moderate Drizzle",
		55: "Dense intensity Drizzle",
	}
	return m[code]
}
