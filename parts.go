package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Subject of the spark the user is requesting
type Subject int

const (
	Nature Subject = iota
	Civilization
	People
	Combat
	Person
	Name
	Characteristic
	Object
	Beast
	State
	Theme
	subjectCount
)

var subjectName = map[Subject]string{
	Nature:         "Nature",
	Civilization:   "Civilization",
	People:         "People",
	Combat:         "Combat",
	Person:         "Person",
	Name:           "Name",
	Characteristic: "Characteristic",
	Object:         "Object",
	Beast:          "Beast",
	State:          "State",
	Theme:          "Theme",
}

// have the subject show up as its name rather than the number value
func (subject Subject) String() string {
	return subjectName[subject]
}

type SparkTable struct {
	Subject
	subType     string
	descriptor1 string
	descriptor2 string
	options1    []string
	options2    []string
}

func GetAndPrintSparkResult(subject Subject, num int) {
	fmt.Println("Spark:")
	fmt.Println(subject)
	fmt.Println(NatureType(num)) //TODO this can't stay hardcoded to NatureType

	var table, err = getSparkTable(subject, NatureType(num).String())

	if err == nil {
		fmt.Println(table.descriptor1 + ": " + table.options1[rand.Intn(len(table.options1))])
		fmt.Println(table.descriptor2 + ": " + table.options2[rand.Intn(len(table.options2))])
	} else {
		fmt.Println(err)
	}
}

func getSparkTable(subject Subject, tableName string) (SparkTable, error) {
	switch {
	case subject == Subject(0): //Nature
		table, ok := natureTableMap[tableName]
		if ok {
			return table, nil
		} else {
			return natureTableMap[tableName], errors.New("No spark table for " + subject.String() + " and " + tableName)
		}
	default:
		return natureTableMap[tableName], errors.New("No spark table for " + subject.String() + " and " + tableName) //TODO determine better thing to return than a default table on default case here
	}
}
