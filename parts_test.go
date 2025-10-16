package main

import (
	"reflect"
	"testing"
)

func TestGetSparkTableWithExistingTable(t *testing.T) {
	want := SparkTable{Nature, NatureType(0).String(), "Character", "Landscape", landCharacterSlice, landscapeSlice}
	table, err := getSparkTable(Nature, "Land")
	if err != nil || !reflect.DeepEqual(want, table) {
		t.Errorf(`getSparkTable(Nature, "Land") = %q, %v, want match for %#q, nil`, table, err, want)
	}
}

func TestGetSparkTableWithNonxistantTable(t *testing.T) {

}

func TestGetSparkTableWithNonxistantSubject(t *testing.T) {

}
