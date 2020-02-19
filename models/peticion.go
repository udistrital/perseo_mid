package models

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
)

// GetJSONJBPM ...
func GetJSONJBPM(urlp string, target interface{}) error {
	b := new(bytes.Buffer)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlp, b)
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

// SendJSONJBPM ...
func SendJSONJBPM(urlp string, trequest string, target interface{}, datajson interface{}) error {
	b := new(bytes.Buffer)
	if datajson != nil {
		// logs.Info(datajson)
		json.NewEncoder(b).Encode(datajson)
	}
	// jsonValue, _ := json.Marshal(datajson)
	//proxyUrl, err := url.Parse("http://10.20.4.15:3128")
	//http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}

	client := &http.Client{}
	req, err := http.NewRequest(trequest, urlp, b)
	// req, err := http.http.Post(urlp, "application/json", bytes.NewBuffer(jsonValue))
	req.Header.Set("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		beego.Error("Error reading response. ", err)
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
