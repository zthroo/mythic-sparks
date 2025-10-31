package parts

import (
	"errors"
	"fmt"
)

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

func ParseCivilizationType(input string) (CivilizationType, error) {
	//TODO strip whitespace, first character uppercase, only alpha
	switch input {
	case Holding.String():
		return Holding, nil
	case Bailey.String():
		return Bailey, nil
	case Keep.String():
		return Keep, nil
	case Food.String():
		return Food, nil
	case Goods.String():
		return Goods, nil
	case Luxuries.String():
		return Luxuries, nil
	case Drama.String():
		return Drama, nil
	case Woe.String():
		return Woe, nil
	case News.String():
		return News, nil
	}

	return Holding, fmt.Errorf("%q is not a valid Civilization type: %w", input, errors.New("invalid Civilization type"))
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

// First of two food descriptors
var foodQualitySlice = []string{
	"Spiced",
	"Herbal",
	"Crunchy",
	"Sour",
	"Dry",
	"Fermented",
	"Salted",
	"Wet",
	"Fatty",
	"Chewy",
	"Sweet",
	"Mild",
}

// Second of two food descriptors
var foodTypeSlice = []string{
	"Fish",
	"Fruit",
	"Stew",
	"Mushrooms",
	"Pie",
	"Cheese",
	"Nuts",
	"Cake",
	"Porridge",
	"Bread",
	"Vegetables",
	"Meat",
}

var foodSparkTable = SparkTable{Civilization, CivilizationType(3).String(), "Quality", "Type", foodQualitySlice, foodTypeSlice}

// First of two goods descriptors
var goodsThemeSlice = []string{
	"Military",
	"Abundant",
	"Traditional",
	"Specialist",
	"Industrious",
	"Innovative",
	"Secretive",
	"Simple",
	"Strong",
	"Decorated",
	"Fine",
	"Lucky",
}

// Second of two goods descriptors
var goodsTypeSlice = []string{
	"Textile",
	"Livestock",
	"Grain",
	"Mead",
	"Tools",
	"Stone",
	"Wood",
	"Pottery",
	"Metal",
	"Leather",
	"Honey",
	"Herb",
}

var goodsSparkTable = SparkTable{Civilization, CivilizationType(4).String(), "Theme", "Type", goodsThemeSlice, goodsTypeSlice}

var civilizationTableMap = map[string]SparkTable{
	CivilizationType(0).String(): holdingSparkTable,
	CivilizationType(1).String(): baileySparkTable,
	CivilizationType(2).String(): keepSparkTable,
	CivilizationType(3).String(): foodSparkTable,
	CivilizationType(4).String(): goodsSparkTable,
}
