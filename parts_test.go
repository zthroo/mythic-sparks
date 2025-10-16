package main

import (
	"errors"
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
	want := errors.New("No spark table for " + Nature.String() + " and " + "TEST")
	table, err := getSparkTable(Nature, "TEST")
	if err == nil {
		t.Errorf(`getSparkTable(Nature, "TEST") error = %q, want nil. table = %v`, err, table)
	}
	if err.Error() != want.Error() {
		t.Errorf(`getSparkTable(Nature, "TEST") error = %q, want match for %v`, err, want)
	}
}
