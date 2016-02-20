package LangDetect

import (
	"github.com/GeertJohan/go.rice"
	"github.com/MichaelLeachim/NumBayes"
	"log"
	"regexp"
)

var LANGSPLITTER = regexp.MustCompile("[^\\s\\n]+")

func LangDetect() (func(data string) LangDetectorLangItem, error) {
	templateBox, err := rice.FindBox("bases")
	if err != nil {
		return nil, err
	}

	data, err := templateBox.Bytes("langs.gob")
	if err != nil {
		return nil, err
	}

	solver := bayes.BayesMemory{}
	solver.UnMarshal(data)
	return func(data string) LangDetectorLangItem {
		result := solver.Classify(LANGSPLITTER.FindAllString(data, -1))
		if len(result) > 0 {
			return LANGMAP[result[0].Category]
		}
		log.Fatal("No mapping opened")
		return LangDetectorLangItem{}
	}, nil
}
