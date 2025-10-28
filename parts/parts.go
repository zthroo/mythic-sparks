package parts

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

func ParseSubject(input string) (Subject, error) {
	//TODO strip whitespace, first character uppercase, only alpha
	switch input {
	case Nature.String():
		return Nature, nil
	case Civilization.String():
		return Civilization, nil
	case People.String():
		return People, nil
	case Combat.String():
		return Combat, nil
	case Person.String():
		return Person, nil
	case Name.String():
		return Name, nil
	case Characteristic.String():
		return Characteristic, nil
	case Object.String():
		return Object, nil
	case Beast.String():
		return Beast, nil
	case State.String():
		return State, nil
	case Theme.String():
		return Theme, nil
	}

	return Nature, fmt.Errorf("%q is not a valid Subject type: %w", input, errors.New("invalid subject"))
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

	table, err := getSparkTable(subject, num)

	if err == nil {
		fmt.Println(table.subType)
		fmt.Println(table.descriptor1 + ": " + table.options1[rand.Intn(len(table.options1))])
		fmt.Println(table.descriptor2 + ": " + table.options2[rand.Intn(len(table.options2))])
	} else {
		fmt.Println(err)
	}
}

func getSparkTable(subject Subject, num int) (SparkTable, error) {
	switch {
	case subject == Subject(0): //Nature
		tableName := NatureType(num).String()
		table, ok := natureTableMap[tableName]
		if ok {
			return table, nil
		} else {
			return natureTableMap[tableName], errors.New("No spark table for Subject: " + subject.String() + " and subtype: " + tableName)
		}
	case subject == Subject(1): //Civilization
		tableName := CivilizationType(num).String()
		table, ok := civilizationTableMap[tableName]
		if ok {
			return table, nil
		} else {
			return civilizationTableMap[tableName], errors.New("No spark table for Subject: " + subject.String() + " and subtype: " + tableName)
		}
	default:
		tableName := NatureType(num).String()
		return natureTableMap[tableName], errors.New("No spark table for Subject: " + subject.String() + " and subtype: " + tableName) //TODO determine better thing to return than a default table on default case here
	}
}
