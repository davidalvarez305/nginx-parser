package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
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

	if len(os.Args) < 2 {
		fmt.Print("Missing dir path arg\n")
		os.Exit(1)
	}

	DIR_PATH := os.Args[1]

	items, _ := ioutil.ReadDir(DIR_PATH)

	for _, file := range items {

		bytes, err := os.ReadFile(DIR_PATH + "/" + file.Name())

		if err != nil {
			fmt.Println(err)
		}

		if strings.Contains(file.Name(), "access") {
			arr := strings.Split(string(bytes), "\n")

			method := regexp.MustCompile("[A-Z]{3}")
			url := regexp.MustCompile(`(\/)([[a-z\-]{5,})`)
			bot := regexp.MustCompile("([a-zA-Z0-9]+)(Bot|bot)")
			status := regexp.MustCompile("( [0-9]{3} )")

			for i := 0; i < len(arr); i++ {
				m := method.FindString(arr[i])
				u := url.FindString(arr[i])
				b := bot.FindString(arr[i])
				s := strings.TrimSpace(status.FindString(arr[i]))
				fin := Request{
					Method: m,
					Url:    u,
					Bot:    b,
					Status: s,
				}
				var status int

				if len(s) > 0 {
					sc, err := strconv.Atoi(s)

					if err != nil {
						fmt.Println("Error converting status to int.")
					}

					status = sc
				}

				if len(b) > 0 && status >= 400 {
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
		}
	}
	fmt.Printf("%+v\n", botCounter)
}
