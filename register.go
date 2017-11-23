package sdlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Register the service, its version and its endpoints
func Register(service string, version string, paths []*Path) {

	var buf = new(bytes.Buffer)

	instance, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	url := buildRegisterURL(service, instance)

	payload := &Manifest{Paths: paths, Hostname: instance, IP: ip, Version: version}
	json.NewEncoder(buf).Encode(&payload)

	req, err := http.NewRequest("POST", url, buf)
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

func buildRegisterURL(service string, instance string) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://")
	buffer.WriteString(serviceDiscoveryURL)
	buffer.WriteString("/")
	buffer.WriteString(service)
	buffer.WriteString("/")
	buffer.WriteString(instance)
	return buffer.String()
}

//
