package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	flag.Parse()
	domain := flag.Arg(0)

	// Check A records
	aRecords, err := net.LookupHost(domain)
	if err != nil {
		fmt.Println("Error fetching A records:", err)
	} else {
		fmt.Println("A records:", aRecords)
	}

	// Check MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("Error fetching MX records:", err)
	} else {
		for _, mx := range mxRecords {
			fmt.Printf("MX record: %s %d\n", mx.Host, mx.Pref)
		}
	}

	// Check TXT records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Println("Error fetching TXT records:", err)
	} else {
		fmt.Println("TXT records:", txtRecords)
	}

	// Check NS records
	nsRecords, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println("Error fetching NS records:", err)
	} else {
		for _, ns := range nsRecords {
			fmt.Printf("NS record: %s\n", ns.Host)
		}
	}

	// Check CNAME record
	cnameRecord, err := net.LookupCNAME(domain)
	if err != nil {
		fmt.Println("Error fetching CNAME record:", err)
	} else {
		fmt.Println("CNAME record:", cnameRecord)
	}
}
