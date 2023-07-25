package bom

import (
	"fmt"
	pprint "github.com/NubeDev/bom-api/print"
	"testing"
)

func TestClient_Observations(t *testing.T) {
	search, err := client.ObservationByTown("helensburgh", "nsw")
	if err != nil {
		fmt.Print(err)
	}
	pprint.PrintJOSN(search)
}
