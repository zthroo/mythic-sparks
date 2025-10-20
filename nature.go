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
	Weather:    "Weather",
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

// First of two sky descriptors
var skyToneSlice = []string{
	"Glittering",
	"Violet",
	"Sapphire",
	"Pale",
	"Fiery",
	"Ivory",
	"Slate",
	"Pink",
	"Golden",
	"Bloody",
	"Bright",
	"Inky",
}

// Second of two sky descriptors
var skyTextureSlice = []string{
	"Aurora",
	"Haze",
	"Marble",
	"Glow",
	"Billows",
	"Swirl",
	"Streaks",
	"Dapple",
	"Rays",
	"Pillars",
	"Shimmer",
	"Swells",
}

var skySparkTable = SparkTable{Nature, NatureType(1).String(), "Tone", "Texture", skyToneSlice, skyTextureSlice}

// First of two water descriptors
var waterToneSlice = []string{
	"Crystal",
	"Teal",
	"Pearlescent",
	"Mucky",
	"Cobalt",
	"Verdant",
	"Frosted",
	"Dark",
	"Verdigris",
	"Silver",
	"Emerald",
	"Jade",
}

// Second of two water descriptors
var waterTextureSlice = []string{
	"Silk",
	"Ripples",
	"Abyss",
	"Churn",
	"Froth",
	"Mirror",
	"Surge",
	"Glass",
	"Surf",
	"Rapids",
	"Spray",
	"Bubbles",
}

var waterSparkTable = SparkTable{Nature, NatureType(2).String(), "Tone", "Texture", waterToneSlice, waterTextureSlice}

// First of two weather descriptors
var weatherDescriptionSlice = []string{
	"Gentle",
	"Fleeting",
	"Persistent",
	"Bright",
	"Thin",
	"Cool",
	"Hot",
	"Solid",
	"Dull",
	"Faint",
	"Abundant",
	"Harsh",
}

// Second of two weather descriptors
var weatherElementSlice = []string{
	"Rain",
	"Gusts",
	"Cloud",
	"Sunlight",
	"Mist",
	"Humidity",
	"Thunder",
	"Dust",
	"Warmth",
	"Drizzle",
	"Breeze",
	"Fog",
}

var weatherSparkTable = SparkTable{Nature, NatureType(3).String(), "Description", "Element", weatherDescriptionSlice, weatherElementSlice}

// First of two flora descriptors
var floraNatureSlice = []string{
	"Aromatic",
	"Ashen",
	"Blooming",
	"Twisted",
	"Towering",
	"Fruitful",
	"Stinging",
	"Vibrant",
	"Brittle",
	"Thorny",
	"Sturdy",
	"Resinous",
}

// Second of two flora descriptors
var floraFormSlice = []string{
	"Grasses",
	"Heather",
	"Shrubs",
	"Brambles",
	"Canopy",
	"Ferns",
	"Trunks",
	"Vines",
	"Conifers",
	"Saplings",
	"Reeds",
	"Roots",
}

var floraSparkTable = SparkTable{Nature, NatureType(4).String(), "Nature", "Form", floraNatureSlice, floraFormSlice}

var natureTableMap = map[string]SparkTable{
	NatureType(0).String(): landSparkTable,
	NatureType(1).String(): skySparkTable,
	NatureType(2).String(): waterSparkTable,
	NatureType(3).String(): weatherSparkTable,
	NatureType(4).String(): floraSparkTable,
}
