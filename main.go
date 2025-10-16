package main

import (
	"fmt"
	"math/rand"
)

func main() {
	getAndPrintSparkResult(Nature, 0)
	getAndPrintSparkResult(Nature, 1)
}

func getAndPrintSparkResult(subject Subject, num int) {
	fmt.Println("Spark:")
	fmt.Println(subject)
	fmt.Println(NatureType(num))

	var table = getSparkTable(subject, NatureType(num).String())

	fmt.Println(table.descriptor1 + ": " + table.options1[rand.Intn(len(table.options1))])
	fmt.Println(table.descriptor2 + ": " + table.options2[rand.Intn(len(table.options2))])
}

func getSparkTable(subject Subject, table string) SparkTable {
	switch {
	case subject == Subject(0): //Nature
		return natureTableMap[table]
	default:
		return natureTableMap[table]
	}
}
