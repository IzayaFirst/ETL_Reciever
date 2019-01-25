package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Address struct {
	City     string `json:"city"`
	State    string `json:"state"`
	Postcode string `json:"postcode"`
	Address1 string `json:"address_1"`
	Address2 string `json:"address_2"`
	Address3 string `json:"address_3"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
	ParcelID string `json:"parcel_id"`
}

func main() {

	address := readData()
	for _, element := range address {
		fmt.Println(element)
	}

}

func readData() []Address {
	csvFile, _ := os.Open("receiver_detail.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var address []Address

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		address = append(address, Address{
			City:     line[3],
			State:    line[4],
			Postcode: line[5],
			Address1: line[0],
			Address2: line[1],
			Address3: line[2],
			Lat:      line[6],
			Lng:      line[7],
			ParcelID: line[8],
		})
	}
	return address
}
