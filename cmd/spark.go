package cmd

import (
	"github.com/zthroo/mythic-sparks/parts"
)

func Generate(subject string, subtype string) (string, error) {
	parsedSubject, err := parts.ParseSubject(subject)
	if err != nil {
		return "", err
	}

	switch parsedSubject {
	case parts.Nature:
		parsedSubtype, err := parts.ParseNatureType(subtype)
		if err != nil {
			return "", err
		}
		spark, err := parts.GetSparkResult(parsedSubject, int(parsedSubtype))
		if err != nil {
			return "", err
		}
		return spark, err
	case parts.Civilization:
		parsedSubtype, err := parts.ParseCivilizationType(subtype)
		if err != nil {
			return "", err
		}
		spark, err := parts.GetSparkResult(parsedSubject, int(parsedSubtype))
		if err != nil {
			return "", err
		}
		return spark, err
	}
	return "", err
}
