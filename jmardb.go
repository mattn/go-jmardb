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

func (c *client) Search(status string, infotype string, title string, headtitle string, eventid string, areaname string, areacode string, datetime string, reportdatetime string, targetdatetime string, validdatetime string, limit int, offset int, order string) ([]Data, error) {
	param := url.Values{}
	param.Set("status", status)
	param.Set("infotype", infotype)
	param.Set("title", title)
	param.Set("headtitle", headtitle)
	param.Set("eventid", eventid)
	param.Set("areaname", areaname)
	param.Set("areacode", areacode)
	param.Set("datetime", datetime)
	param.Set("reportdatetime", reportdatetime)
	param.Set("targetdatetime", targetdatetime)
	param.Set("validdatetime", validdatetime)
	param.Set("limit", fmt.Sprint(limit))
	param.Set("offset", fmt.Sprint(offset))
	param.Set("order", order)
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

func (c *client) Area(latitude, longitude float64, limit int) ([]Data, error) {
	param := url.Values{}
	param.Set("latitude", fmt.Sprint(latitude))
	param.Set("longitude", fmt.Sprint(longitude))
	if limit != -1 {
		param.Set("limit", fmt.Sprintf("%d", limit))
	}
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
