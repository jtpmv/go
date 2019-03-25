package main

import(
	"fmt"
	"net/http"
)

func SimpleBasicAuth(username, password string) BasicAuthFunc {
	return BasicAuthFunc(func(user, pass string) bool {
		return username == user && password == pass
	})
}

