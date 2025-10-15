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

// Have the nature type show up as its name rather than the number value
func (natureType NatureType) String() string {
	return natureTypeName[natureType]
}

// First of two land descriptors
type LandCharacter int

const (
	Barren LandCharacter = iota
	Dry
	Grey
	Sparse
	Sharp
	Teeming
	Still
	Soft
	Overgrown
	Vivid
	Sodden
	Lush
	landCharacterCount
)

var landCharacterName = map[LandCharacter]string{
	Barren:    "Barren",
	Dry:       "Dry",
	Grey:      "Grey",
	Sparse:    "Sparse",
	Sharp:     "Sharp",
	Teeming:   "Teeming",
	Still:     "Still",
	Soft:      "Soft",
	Overgrown: "Overgrown",
	Vivid:     "Vivid",
	Sodden:    "Sodden",
	Lush:      "Lush",
}

func (landCharacter LandCharacter) String() string {
	return landCharacterName[landCharacter]
}

// Second of two land descriptors
type Landscape int

const (
	Marsh Landscape = iota
	Heath
	Crags
	Peaks
	Forest
	Valley
	Hills
	Meadow
	Bog
	Lakes
	Glades
	Plain
	landscapeCount
)

var landscapeName = map[Landscape]string{
	Marsh:  "Marsh",
	Heath:  "Heath",
	Crags:  "Crags",
	Peaks:  "Peaks",
	Forest: "Forest",
	Valley: "Valley",
	Hills:  "Hills",
	Meadow: "Meadow",
	Bog:    "Bog",
	Lakes:  "Lakes",
	Glades: "Glades",
	Plain:  "Plain",
}

func (Landscape Landscape) String() string {
	return landscapeName[Landscape]
}
