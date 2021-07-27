package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/oschwald/geoip2-golang"
)

func main() {
	file, err := os.Open("ip.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	c := make(map[string]int)
	db, err := geoip2.Open("country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for scanner.Scan() {
		ip := scanner.Text()
		name := getCountryName(db, ip)
		c[name]++
	}

	fmt.Println(c)
}

func getCountryName(db *geoip2.Reader, ipAddr string) string {
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipAddr)
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Chinese country name: %v\n", record.Country.Names["zh-CN"])
	//fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	//fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
	return record.Country.Names["zh-CN"]
}
