package tsunami

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readCSV(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	var line uint64
	var headers []string
	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		// fmt.Printf("%+v\n", rec)
		if line == 0 {
			headers = record
		}
		line++
	}
	fmt.Printf("%+v\n", headers[0])
	fmt.Println("------")
	fmt.Printf("- Read %d|%s [Complate]\n", line, filename)
}

func main() {
	// open file
	readCSV("in.product.csv")
}