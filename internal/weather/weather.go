package weather

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const baseUrl = "https://api.openweathermap.org/data/2.5/weather"

type Condition struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

type Weather struct {
	Name      string      `json:"name"`
	Condition []Condition `json:"weather"`
	Main      Main        `json:"main"`
	Wind      Wind        `json:"wind"`
}

func Forecast(city string, token string) (*Weather, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("q", city)
	q.Add("units", "metric")
	q.Add("lang", "ru")
	q.Add("appid", token)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var weather Weather

	if err = json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
