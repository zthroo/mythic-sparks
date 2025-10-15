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

type NatureType int

const (
	Land NatureType = iota
	Sky
	Water
	Weather
	Flora
	Fauna
	Feature
	Wonder
	Otherworld
	natureTypeCount
)

var natureTypeName = map[NatureType]string{
	Land:       "Land",
	Sky:        "Sky",
	Water:      "Water",
	Weather:    "Combat",
	Flora:      "Flora",
	Fauna:      "Fauna",
	Feature:    "Feature",
	Wonder:     "Wonder",
	Otherworld: "Otherworld",
}

// have the nature type show up as its name rather than the number value
func (natureType NatureType) String() string {
	return natureTypeName[natureType]
}
