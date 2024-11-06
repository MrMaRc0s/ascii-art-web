package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(banner string) ([]string, error) {
	bannerFile, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		return nil, fmt.Errorf("there was a problem reading the banner file template")
	}
	line := strings.Replace(string(bannerFile), "\r\n", "\n", -1)
	return strings.Split(line, "\n"), nil
}
