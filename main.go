package LangDetect
import (
	"github.com/MichaelLeachim/NumBayes"
	"io/ioutil"
	"log"
	"regexp"
)
var LANGSPLITTER = regexp.MustCompile("[^\\s\\n]+")

func LangDetect(basepath string) (func(data string) LangDetectorLangItem, error) {
	data, err := ioutil.ReadFile(basepath)
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
