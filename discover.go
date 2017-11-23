package sdlib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Discover the service, its version and its endpoints
func Discover(service string, version string, endpoints *[]Endpoint) {

	var buf = new(bytes.Buffer)

	url := buildDiscoverURL(service)

	req, err := http.NewRequest("GET", url, buf)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

// helpers

func buildDiscoverURL(service string) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://")
	buffer.WriteString(serviceDiscoveryURL)
	buffer.WriteString("/")
	buffer.WriteString(service)
	return buffer.String()
}

//
