package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type report struct {
	ASN     uint
	IP      string
	City    string
	Country string
	Org     string
}

func asn(ip net.IP, dbNet string) (uint, string) {
	db, err := geoip2.Open(dbNet)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	record, err := db.ASN(ip)
	if err != nil {
		log.Fatal(err)
	}
	asNet := record.AutonomousSystemNumber
	asOrg := record.AutonomousSystemOrganization
	return asNet, asOrg
}

func geo(ip net.IP, dbGEO string) (string, string) {
	db, err := geoip2.Open(dbGEO)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	country := record.Country.Names["en"]
	city := record.City.Names["en"]
	return country, city
}

func output(report *report) {
	result := []string{report.IP,
		report.Country,
		report.City,
		strconv.FormatUint(uint64(report.ASN), 10),
		report.Org,
	}
	fmt.Println(strings.Join(result, ","))
}

func process(dbNet string, dbGeo string, readLine string) {
	var report report
	report.IP = readLine
	ip := net.ParseIP(readLine)
	report.Country, report.City = geo(ip, dbGeo)
	if dbNet != "" {
		report.ASN, report.Org = asn(ip, dbNet)
	}
	output(&report)

}

func handler(dbNet string, dbGeo string, filehost string, short string) {
	if short == "" {
		file, _ := os.Open(filehost)
		f := bufio.NewReader(file)
		for {
			readLine, _ := f.ReadString('\n')
			if readLine != "" {
				readLine = strings.TrimRight(readLine, "\n")
				process(dbNet, dbGeo, readLine)
			} else {
				break
			}

		}
	} else if short != "" {
		process(dbNet, dbGeo, short)
	} else {
		log.Fatal()
	}

}
func main() {
	filename := "ip_part"
	dbGeoPtr := flag.String("dbGeo", "GeoLite2-City.mmdb", "Path to database GEO")
	dbASNPtr := flag.String("dbNet", "GeoLite2-ASN.mmdb", "Path to database ASN")
	filePtr := flag.String("f", filename, "Path to file")
	shortPtr := flag.String("a", "", "IP check")
	flag.Parse()
	handler(*dbASNPtr, *dbGeoPtr, *filePtr, *shortPtr)
}
