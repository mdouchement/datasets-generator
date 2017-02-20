package main

import (
	"math/rand"

	uuid "github.com/satori/go.uuid"
)

type baseRows [][]string

var (
	staticHeader = []string{"KEY", "class", "Xäé", "Yŷè", "cust_id", "sub_id"}
	baseRowPool  = baseRows{
		[]string{"uuid", "Trüe", "éà", "ëx"},
		[]string{"uuid", "Fälse", "ïö", "ôo"},
	}
)

// Header returns a CSV header
func Header() []string {
	header := make([]string, len(staticHeader))
	copy(header, staticHeader)
	for i := 0; i < nbOfColumns-len(staticHeader); i++ {
		header = append(header, RandStringBytesMaskImprSrc(varLength))
	}
	return header
}

// Row returns a CSV line
func Row(generators []func(int) string) []string {
	base := baseRowPool[rand.Intn(len(baseRowPool))] // Choose a random base from the pool
	row := make([]string, len(base))
	copy(row, base)
	row[0] = uuid.NewV4().String() // KEY

	for _, generator := range generators {
		row = append(row, generator(varLength)) // Generate random fields
	}
	return row
}
