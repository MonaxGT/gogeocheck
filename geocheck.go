package main

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"strings"
	"fmt"
	"os"
	"bufio"
	"flag"
)

func process_shot (database string, ip_address string, cityon bool ) {
	db, err := geoip2.Open(database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ip := net.ParseIP(strings.TrimRight(ip_address,"\n"))
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	if cityon == false {
		fmt.Printf("%v,%v\n", ip, record.Country.Names["en"])
	} else {
		fmt.Printf("%v,%v,%v\n", ip, record.Country.Names["en"], record.City.Names["en"])
	}

}

func process (database string, filehost string, cityon bool ) {
	file, _ := os.Open(filehost)
	f := bufio.NewReader(file)
	db, err := geoip2.Open(database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for {
		read_line, _ := f.ReadString('\n')
		if read_line != "" {
			ip := net.ParseIP(strings.TrimRight(read_line,"\n"))
			record, err := db.City(ip)
			if err != nil {
				log.Fatal(err)
			}
			if cityon == false {
				fmt.Printf("%v,%v\n", ip, record.Country.Names["en"])
			} else {
				fmt.Printf("%v,%v,%v\n", ip, record.Country.Names["en"], record.City.Names["en"])
			}
		} else {
			break
		}

	}

}
func main() {
	dbname := "GeoLite2-City.mmdb"
	filename :="ip_part"
	databasePtr := flag.String("db", dbname,"Path to database" )
	filePtr := flag.String("f", filename ,"Path to file" )
	cityPtr := flag.Bool("c", false,"Include information about City" )
	shootPtr := flag.String("a", "false","IP to resolve" )
	helpPtr := flag.Bool("h", false,"help" )
	flag.Parse()

	if (*helpPtr == false) && (*shootPtr != "false") {
		process_shot(*databasePtr,*shootPtr,*cityPtr)
	} else if *helpPtr == false {
		process(*databasePtr,*filePtr,*cityPtr)
	} else {
		fmt.Printf("-db	Path to database\n" + "-f	Path to file\n" + "-c	Include information about City\n" + "-a	ipaddress IP to resolve\n" + "-h	help\n")
		fmt.Printf("You can don't use flags, if import db to directory with program and use name GeoLite2-City.mmdb, use  list of ip addresses with name ip_part\n")
	}


}

