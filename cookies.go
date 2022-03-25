package main

import (
	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/allbrowsers"
)

func getCookies() map[string]map[string]string {
	cookies := kooky.ReadCookies()
	data := make(map[string]map[string]string)

	for _, cookie := range cookies {
		if _, ok := data[cookie.Domain]; !ok {
			data[cookie.Domain] = make(map[string]string)
		}
		data[cookie.Domain][cookie.Name] = cookie.Value
	}

	return data
}
