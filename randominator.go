package main

import "math/rand"

var generatorTypes = []string{"categorical", "numerical", "date", "time", "timestamp"}

// Generators returns a slice of datatype generator for each colunm
func Generators(number int) []func(int) string {
	generators := make([]func(int) string, number)
	for i := 0; i < number; i++ {
		generators[i] = newRandomGenerator()
	}
	return generators
}

func newRandomGenerator() func(int) string {
	switch generatorTypes[rand.Intn(len(generatorTypes))] {
	case "categorical":
		if rand.Intn(2) == 0 {
			return RandStringRunes
		}
		return RandStringBytesMaskImprSrc
	case "numerical":
		return RandNumerical
	case "date":
		return RandDate()
	case "time":
		return RandTime()
	case "timestamp":
		return RandTimestamp()
	default:
		return RandStringBytesMaskImprSrc
	}
}
