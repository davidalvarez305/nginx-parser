package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Request struct {
	Method string `json:"method"`
	Url    string `json:"url"`
	Bot    string `json:"bot"`
	Status string `json:"status"`
}

type Counter struct {
	Google  int `json:"google"`
	Bing    int `json:"bing"`
	Ahrefs  int `json:"ahrefs"`
	SemRush int `json:"semrush"`
}

func main() {
	var req []Request
	var botCounter Counter
	bytes, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}

	arr := strings.Split(string(bytes), "\n")

	method := regexp.MustCompile("[A-Z]{3}")
	url := regexp.MustCompile(`(\/)([[a-z\-]{5,})`)
	bot := regexp.MustCompile("([a-zA-Z0-9]+)(Bot|bot)")
	status := regexp.MustCompile("( [0-9]{3} )")

	for i := 0; i < len(arr); i++ {
		m := method.FindString(arr[i])
		u := url.FindString(arr[i])
		b := bot.FindString(arr[i])
		s := status.FindString(arr[i])
		fin := Request{
			Method: m,
			Url:    u,
			Bot:    b,
			Status: strings.TrimSpace(s),
		}
		if len(b) > 0 {
			req = append(req, fin)
			if b == "Googlebot" {
				botCounter.Google += 1
			}
			if b == "AhrefsBot" {
				botCounter.Ahrefs += 1
			}
			if b == "bingbot" {
				botCounter.Bing += 1
			}
			if b == "SemrushBot" {
				botCounter.SemRush += 1
			}
		}
	}
	fmt.Printf("%+v\n", req)
	fmt.Printf("%+v\n", botCounter)
}
