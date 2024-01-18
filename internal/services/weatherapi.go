package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type WeatherApiResponse struct {
	Location WeatherApiResponseLocation `json:"location"`
	Current  WeatherApiResponseCurrent  `json:"current"`
}

type WeatherApiResponseLocation struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type WeatherApiResponseCurrent struct {
	TemperatureCelsius float64 `json:"temp_c"`
}

type WeatherApiService interface {
	QueryWeather(ctx context.Context, location string) (*WeatherApiResponse, error)
}

type WeatherApiServiceImpl struct {
	client *http.Client
}

func NewWeatherApiService() WeatherApiService {
	return &WeatherApiServiceImpl{
		client: &http.Client{},
	}
}

func (s *WeatherApiServiceImpl) QueryWeather(ctx context.Context, location string) (*WeatherApiResponse, error) {

	queryParams := url.Values{}
	queryParams.Add("key", os.Getenv("WEATHER_API_KEY"))
	queryParams.Add("q", location)
	queryParams.Add("aqi", "no")

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?%s", queryParams.Encode())

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)

	response, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("WeatherAPI API error")
	}

	weatherApiResponse := WeatherApiResponse{}
	err = json.Unmarshal([]byte(body), &weatherApiResponse)
	if err != nil {
		return nil, errors.New("invalid WeatherAPI API response")
	}

	return &weatherApiResponse, nil
}
