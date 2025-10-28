package parts

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetSparkTableWithExistingTable(t *testing.T) {
	want := SparkTable{Nature, NatureType(0).String(), "Character", "Landscape", landCharacterSlice, landscapeSlice}
	table, err := getSparkTable(Nature, 0)
	if err != nil || !reflect.DeepEqual(want, table) {
		t.Errorf(`getSparkTable(Nature, "Land") = %q, %v, want match for %#q, nil`, table, err, want)
	}
}

func TestGetSparkTableWithNonexistantTable(t *testing.T) {
	want := errors.New("No spark table for Subject: " + Nature.String() + " and subtype: " + "")
	table, err := getSparkTable(Nature, 27)
	if err == nil {
		t.Errorf(`getSparkTable(Nature, "TEST") error = %q, want nil. table = %v`, err, table)
	}
	if err.Error() != want.Error() {
		t.Errorf(`getSparkTable(Nature, "TEST") error = %q, want match for %v`, err, want)
	}
}

func TestGetSparkTableWithNotImplementedSubject(t *testing.T) {
	want := errors.New("No spark table for Subject: " + subjectCount.String() + " and subtype: " + "")
	table, err := getSparkTable(subjectCount, 27)
	if err == nil {
		t.Errorf(`getSparkTable(subjectCount, "TEST") error = %q, want nil. table = %v`, err, table)
	}
	if err.Error() != want.Error() {
		t.Errorf(`getSparkTable(subjectCount, "TEST") error = %q, want match for %v`, err, want)
	}
}
