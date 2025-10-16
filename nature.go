package main

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

func (natureType NatureType) String() string {
	return natureTypeName[natureType]
}

// First of two land descriptors
var landCharacterSlice = []string{
	"Barren",
	"Dry",
	"Grey",
	"Sparse",
	"Sharp",
	"Teeming",
	"Still",
	"Soft",
	"Overgrown",
	"Vivid",
	"Sodden",
	"Lush",
}

// Second of two land descriptors
var landscapeSlice = []string{
	"Marsh",
	"Heath",
	"Crags",
	"Peaks",
	"Forest",
	"Valley",
	"Hills",
	"Meadow",
	"Bog",
	"Lakes",
	"Glades",
	"Plain",
}

var landSparkTable = SparkTable{Nature, NatureType(0).String(), "Character", "Landscape", landCharacterSlice, landscapeSlice}

var natureTableMap = map[string]SparkTable{
	NatureType(0).String(): landSparkTable,
}
