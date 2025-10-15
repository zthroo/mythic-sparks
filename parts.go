package main

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
