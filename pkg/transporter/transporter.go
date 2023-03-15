package transporter

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func SendHttpRequest(username, password, url string) (*http.Response, string) {
	// For control over proxies, TLS configuration, keep-alives,
	// compression, and other settings, create a Transport:
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	// set http client
	client := &http.Client{
		Transport: tr,
		Timeout:   30 * time.Second,
	}
	// Create http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// set http request headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(bodyText)

	// return http response and response body as text
	return resp, s

}
