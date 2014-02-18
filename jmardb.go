package jmardb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Forecast struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Data struct {
	JISX0402  string     `json:"jisx0402"`
	Longitude float64    `json:"longitude"`
	Latitude  float64    `json:"latitude"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Forecast  []Forecast `json:"forecast"`
}

type client http.Client

func NewClient() *client {
	return &client{}
}

func (c *client) Area(latitude, longitude float64) ([]Data, error) {
	param := url.Values{}
	param.Set("latitude", fmt.Sprintf("%f", latitude))
	param.Set("longitude", fmt.Sprintf("%f", longitude))
	res, err := http.Get("http://api.aitc.jp/jmardb/area?" + param.Encode())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var datas struct {
		Data []Data `json:"data"`
	}
	err = json.NewDecoder(res.Body).Decode(&datas)
	if err != nil {
		return nil, err
	}
	return datas.Data, nil
}
