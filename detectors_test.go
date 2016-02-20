package LangDetect

import (
	"github.com/MichaelLeachim/NumBayes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	// "os"
	"path/filepath"
	"testing"
)

func train_lang_detector() {
	d := bayes.BayesMemory{}
	for _, item := range LANGMAP {
		f, err := ioutil.ReadFile(filepath.Join("./norma_stop_lists/", item.ISO))
		if err == nil {
			d.Train(LANGSPLITTER.FindAllString(string(f), -1), []string{item.Name})
		} else {
			log.Println(err)
		}
	}
	data, err := d.Marshal()
	if err != nil {
		log.Println(err)
	} else {
		ioutil.WriteFile("./langs.gob", data, 0777)
	}
}

func TestLangDetect(t *testing.T) {
	// os.RemoveAll("./bases/langs.gob")
	// train_lang_detector()
	var lang_samples = map[string]string{
		"English":   `This is a list of file formats used by computers, organized by type. Filename extensions are usually noted in parentheses if they differ from the format n`,
		"Spanish":   `El colapso de la Unión Europea desde el interior es iniciada por el gobierno ruso, - así dijo el actual presidente del Parlamento Europeo, Martin Schulz, expresó que apoyar al Presidente de Ucrania sobre la conducta de varios años de guerra de la información en el territorio de los Estados europeos.`,
		"Russian":   `Что ещё почитать? Twitter Политоты — российские и не только оппозиционные новости, оперативно и полноценно`,
		"Japanese":  `他に何を読むには？ TwitterのPolitoty - ロシアの迅速かつ完全にだけでなく、野党のニュース`,
		"Norwegian": `Hvad andet at læse? Twitter Politoty - russisk og ikke kun oppositionen nyheder hurtigt og fuldt ud`,
		"Bulgarian": `Какво друго да се чете? Twitter Politoty - руски и не само опозицията новини своевременно и напълно`,
		"Italian":   `Che altro da leggere? Twitter Politoty - russo e non solo notizie opposizione tempestivamente e completamente`,
		"Hebrew":    `טוויטר Politoty - רוסית ולא חדשות האופוזיציה רק מיד ובאופן מלא`,
	}

	detector, err := LangDetect()
	if err != nil {
		t.Error(err)
	}
	for k, v := range lang_samples {
		assert.Equal(t, detector(v).Name, k)
	}
	// PASS
}
