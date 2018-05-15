package main

import (
	"fmt"
	"strconv"
)

// LogResponse log response
func LogResponse(method string, path string, statusCode int, proxy bool) {
	if proxy {
		fmt.Printf("%s  \x1b[0;33m↳\x1b[0m  %s %s\n", method, path, colorStatusCode(statusCode))
	} else {
		fmt.Printf("%s %s %s\n", method, path, colorStatusCode(statusCode))
	}
}

func colorStatusCode(statusCode int) string {
	coloredStatus := ""
	switch strconv.Itoa(statusCode)[:1] {
	case "2":
		coloredStatus = fmt.Sprintf("\x1b[0;32m%d\x1b[0m", statusCode)
	case "4":
		coloredStatus = fmt.Sprintf("\x1b[0;33m%d\x1b[0m", statusCode)
	case "3":
		coloredStatus = fmt.Sprintf("\x1b[0;36m%d\x1b[0m", statusCode)
	case "5":
		coloredStatus = fmt.Sprintf("\x1b[0;31m%d\x1b[0m", statusCode)
	}
	return coloredStatus
}
