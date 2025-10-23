package main

type CivilizationType int

const (
	Holding CivilizationType = iota
	Bailey
	Keep
	Food
	Goods
	Luxuries
	Drama
	Woe
	News
	civilizationTypeCount
)

var civilizationTypeName = map[CivilizationType]string{
	Holding:  "Holding",
	Bailey:   "Bailey",
	Keep:     "Keep",
	Food:     "Food",
	Goods:    "Goods",
	Luxuries: "Luxuries",
	Drama:    "Drama",
	Woe:      "Woe",
	News:     "News",
}

func (civilizationType CivilizationType) String() string {
	return civilizationTypeName[civilizationType]
}

// First of two holding descriptors
var holdingStyleSlice = []string{
	"Dark",
	"Ruined",
	"Hostile",
	"Ancient",
	"Ornate",
	"Wild",
	"Pristine",
	"Fortified",
	"Unfinished",
	"Welcoming",
	"Proud",
	"Bright",
}

// Second of two holding descriptors
var holdingFeatureSlice = []string{
	"Turrets",
	"Tower",
	"Wall",
	"Battlements",
	"Citadel",
	"Gate",
	"Spire",
	"Dome",
	"Beacons",
	"Bridge",
	"Pillars",
	"Moat",
}

var holdingSparkTable = SparkTable{Civilization, CivilizationType(0).String(), "Style", "Feature", holdingStyleSlice, holdingFeatureSlice}

// First of two bailey descriptors
var baileyStyleSlice = []string{
	"Dark",
	"Ruined",
	"Hostile",
	"Ancient",
	"Ornate",
	"Wild",
	"Pristine",
	"Fortified",
	"Unfinished",
	"Welcoming",
	"Proud",
	"Bright",
}

// Second of two bailey descriptors
var baileyFeatureSlice = []string{
	"Turrets",
	"Tower",
	"Wall",
	"Battlements",
	"Citadel",
	"Gate",
	"Spire",
	"Dome",
	"Beacons",
	"Bridge",
	"Pillars",
	"Moat",
}

var baileySparkTable = SparkTable{Civilization, CivilizationType(1).String(), "Style", "Feature", baileyStyleSlice, baileyFeatureSlice}

var civilizationTableMap = map[string]SparkTable{
	CivilizationType(0).String(): holdingSparkTable,
	CivilizationType(1).String(): baileySparkTable,
}
