/*
Copyright © 2022 qniwwwersss@gmail.com
*/
package whereAreYou

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	timezone string `json:"timezone"`
	postel   string `json:"postel"`
}

func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response data")
	}

	resInBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return resInBytes
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	resByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(resByte, &data)

	if err != nil {
		log.Println("Unable to unmarshall the response")
	}

	fmt.Println("Data found")
	fmt.Printf("IP: %s\nCITY :%s\nREGION :%s\nCOUNTRY:%s\nLOC :%s\nTIMEZONE: %s\nPOSTEL:%s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.timezone, data.postel)
}

func main() {
	ip := "125.19.91.235"
	showData(ip)
}
