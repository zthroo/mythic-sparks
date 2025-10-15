package main

import (
	"fmt"
	"math/rand"
)

func main() {
	getAndPrintSparkResult(Nature, 0)
}

func getAndPrintSparkResult(subject Subject, num int) {
	fmt.Println("Spark:")
	fmt.Println(subject)
	fmt.Println(NatureType(num))
	fmt.Println(landSparkTable.descriptor1, ": ", landSparkTable.options1[rand.Intn(len(landSparkTable.options1))]) //TODO: spark table not hardcoded
	fmt.Println(landSparkTable.descriptor2, ": ", landSparkTable.options2[rand.Intn(len(landSparkTable.options2))])
}
