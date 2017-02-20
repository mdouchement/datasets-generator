package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	k           = 2
	nbOfColumns = 10
	nbOfLines   = 10000000
	varLength   = 8
	separator   = "|"
	eol         = "\n"
)

func main() {
	if nbOfColumns < len(staticHeader) {
		panic(fmt.Errorf("nbOfColumns must be >= to %d", len(staticHeader)))
	}

	ofile, _ := os.Create("/tmp/fake.tsv")
	defer ofile.Close()
	w := bufio.NewWriter(ofile)

	columns := make(chan []string, k)
	go writer(w, columns)

	columns <- Header()

	generators := Generators(nbOfColumns - len(baseRowPool[0]))
	for i := 1; i < nbOfLines; i++ {
		columns <- Row(generators)
	}
	w.WriteString(eol)
	w.Flush()
}

func writer(w io.Writer, columns <-chan []string) {
	for cs := range columns {
		io.WriteString(w, strings.Join(cs, separator))
		io.WriteString(w, eol)
	}
}
