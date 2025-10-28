package cmd

import (
	"fmt"
	"os"

	"github.com/zthroo/mythic-sparks/parts"
)

func Generate(subject string, subtype string) string {
	parsedSubject, err := parts.ParseSubject(subject)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error occured while executing spark '%s'\n", err)
		os.Exit(1)
	}

	switch parsedSubject {
	case parts.Nature:
		parsedSubtype, err := parts.ParseNatureType(subtype)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while executing spark '%s'\n", err)
			os.Exit(1)
		}
		return parts.GetSparkResult(parsedSubject, int(parsedSubtype))
	case parts.Civilization:
		parsedSubtype, err := parts.ParseCivilizationType(subtype)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Oops. An error occured while executing spark '%s'\n", err)
			os.Exit(1)
		}
		return parts.GetSparkResult(parsedSubject, int(parsedSubtype))
	}
	return "Error"
}
