package bom

import (
	"fmt"
	pprint "github.com/NubeDev/bom-api/print"
	"testing"
)

var client = New(&Client{})

func TestClient_Search(t *testing.T) {
	search, geo, err := client.SearchByTown("Parkville", "Vic")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(geo)
	pprint.PrintJOSN(search)

	search, geo, err = client.SearchByZip("2508")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(geo)
	pprint.PrintJOSN(search)

}

func TestClient_SearchByZip(t *testing.T) {
	search, geo, err := client.SearchByZip("2508")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(geo)
	pprint.PrintJOSN(search)

}
