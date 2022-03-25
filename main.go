package main

import (
	"csvutils"
)

func main() {

	var readFromFile string = "employee.csv"
	var writeToFile string = "writeEmp.csv"

	csvutils.Write(readFromFile,writeToFile)

}
