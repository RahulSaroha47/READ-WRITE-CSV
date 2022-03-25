package csvutils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Write(readFromFile, writeToFile string) {

	csvFile, err := os.Open(readFromFile)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Csv file opened succesfully")
	defer csvFile.Close()

	// reading the csv file
	csvdata, err1 := csv.NewReader(csvFile).ReadAll()

	if err1 != nil {
		fmt.Println(err)
	}

	// creating another csv file

	csvfile2, err2 := os.Create(writeToFile)

	if err2 != nil {
		fmt.Println(err2)
	}

	csvwriter := csv.NewWriter(csvfile2)

	// writing data to csv file
	for i := 0; i < 20; i++ {

		for _, line := range csvdata {

			csvwriter.Write(line)

		}

	}
	csvwriter.Flush()
	csvfile2.Close()
}
