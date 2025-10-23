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
	"Filthy",
	"Abandoned",
	"Joyous",
	"Sophisticated",
	"Industrious",
	"Humble",
	"Majestic",
	"Hallowed",
	"Rustic",
	"Solumn",
	"Bustling",
	"Immaculate",
}

// Second of two bailey descriptors
var baileyFeatureSlice = []string{
	"Marketplace",
	"Forge",
	"Library",
	"Fountain",
	"Temple",
	"Forum",
	"Tomb",
	"Garden",
	"Hall",
	"Workshops",
	"Arena",
	"Garrison",
}

var baileySparkTable = SparkTable{Civilization, CivilizationType(1).String(), "Style", "Feature", baileyStyleSlice, baileyFeatureSlice}

// First of two keep descriptors
var keepCenterpieceSlice = []string{
	"Hearth",
	"Throne",
	"Musicians",
	"Pool",
	"Advisers",
	"Servants",
	"Shrine",
	"Table",
	"Reliquary",
	"Cauldron",
	"Chandelier",
	"Guards",
}

// Second of two keep descriptors
var keepDecorationSlice = []string{
	"Antlers",
	"Silver",
	"Heraldry",
	"Bones",
	"Flowers",
	"Scripture",
	"Jewels",
	"Wreaths",
	"Candles",
	"Fur",
	"Tapestries",
	"Shields",
}

var keepSparkTable = SparkTable{Civilization, CivilizationType(2).String(), "Centerpiece", "Decoration", keepCenterpieceSlice, keepDecorationSlice}

var civilizationTableMap = map[string]SparkTable{
	CivilizationType(0).String(): holdingSparkTable,
	CivilizationType(1).String(): baileySparkTable,
	CivilizationType(2).String(): keepSparkTable,
}
