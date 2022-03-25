package main

import (
	"io/ioutil"
	"net/http"
)

func getIP() string {
	res, err := http.Get("https://api.ipify.org/")
	if err != nil {
		return err.Error()
	}

	defer res.Body.Close()
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	return string(ip)
}
