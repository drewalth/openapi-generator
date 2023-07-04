package main

import (
	"fmt"
	"os"
	"strings"
)

// replaceTildeWithHome takes a string as input and if the first character of
// the string is a "~", it replaces the "~" with the user's HOME directory path.
func replaceTildeWithHome(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting the user's home directory:", err)
			return path
		}
		return home + path[1:]
	}
	return path
}
