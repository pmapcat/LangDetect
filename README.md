# Language detector written in Go. It detects one of the following languages

It returns it's name in English, and it's ISO code.

```go
var LANGMAP = map[string]LangDetectorLangItem{
	"Chinese":    LangDetectorLangItem{ISO: "zh", Name: "Chinese"},
	"Turkish":    LangDetectorLangItem{ISO: "tr", Name: "Turkish"},
	"Swedish":    LangDetectorLangItem{ISO: "sv", Name: "Swedish"},
	"Slovak":     LangDetectorLangItem{ISO: "sk", Name: "Slovak"},
	"Russian":    LangDetectorLangItem{ISO: "ru", Name: "Russian"},
	"Romanian":   LangDetectorLangItem{ISO: "ro", Name: "Romanian"},
	"Portuguese": LangDetectorLangItem{ISO: "pt", Name: "Portuguese"},
	"Polish":     LangDetectorLangItem{ISO: "pl", Name: "Polish"},
	"Norwegian":  LangDetectorLangItem{ISO: "no", Name: "Norwegian"},
	"Dutch":      LangDetectorLangItem{ISO: "nl", Name: "Dutch"},
	"Latvian":    LangDetectorLangItem{ISO: "lv", Name: "Latvian"},
	"Japanese":   LangDetectorLangItem{ISO: "ja", Name: "Japanese"},
	"Italian":    LangDetectorLangItem{ISO: "it", Name: "Italian"},
	"Indonesian": LangDetectorLangItem{ISO: "id", Name: "Indonesian"},
	"Hungarian":  LangDetectorLangItem{ISO: "hu", Name: "Hungarian"},
	"Hindi":      LangDetectorLangItem{ISO: "hi", Name: "Hindi"},
	"French":     LangDetectorLangItem{ISO: "fr", Name: "French"},
	"Finnish":    LangDetectorLangItem{ISO: "fi", Name: "Finnish"},
	"Persian":    LangDetectorLangItem{ISO: "fa", Name: "Persian"},
	"Spanish":    LangDetectorLangItem{ISO: "es", Name: "Spanish"},
	"English":    LangDetectorLangItem{ISO: "en", Name: "English"},
	"Greek":      LangDetectorLangItem{ISO: "el", Name: "Greek"},
	"German":     LangDetectorLangItem{ISO: "de", Name: "German"},
	"Danish":     LangDetectorLangItem{ISO: "da", Name: "Danish"},
	"Czech":      LangDetectorLangItem{ISO: "cz", Name: "Czech"},
	"Catalan":    LangDetectorLangItem{ISO: "ca", Name: "Catalan"},
	"Breton":     LangDetectorLangItem{ISO: "br", Name: "Breton"},
	"Bulgarian":  LangDetectorLangItem{ISO: "bg", Name: "Bulgarian"},
	"Hebrew":     LangDetectorLangItem{ISO: "he", Name: "Hebrew"},
	"Arabic":     LangDetectorLangItem{ISO: "ar", Name: "Arabic"}}
```

## It is used like this

```go
  var lang_samples = map[string]string{
	"English":   This is a list of file formats used by computers, organized by type. Filename extensions are usually noted in parentheses if they differ from the format n",
	"Spanish":   "El colapso de la Unión Europea desde el interior es iniciada por el gobierno ruso, - así dijo el actual presidente del Parlamento Europeo, Martin Schulz, expresó que apoyar al Presidente de Ucrania sobre la conducta de varios años de guerra de la información en el territorio de los Estados europeos.",
	"Russian":   "Что ещё почитать? Twitter Политоты — российские и не только оппозиционные новости, оперативно и полноценно",
	"Japanese":  "他に何を読むには？ TwitterのPolitoty - ロシアの迅速かつ完全にだけでなく、野党のニュース",
	"Norwegian": "Hvad andet at læse? Twitter Politoty - russisk og ikke kun oppositionen nyheder hurtigt og fuldt ud",
	"Bulgarian": "Какво друго да се чете? Twitter Politoty - руски и не само опозицията новини своевременно и напълно",
	"Italian":   "Che altro da leggere? Twitter Politoty - russo e non solo notizie opposizione tempestivamente e completamente",
	"Hebrew":    "טוויטר Politoty - רוסית ולא חדשות האופוזיציה רק מיד ובאופן מלא",
  }
  
 	detector, err := LangDetect("./bases/langs.gob")
	if err != nil {
		t.Error(err)
	}
	for k, v := range LANG_SAMPLES {
		assert.Equal(t, detector(v).Name, k)
	}
  // PASS
```
  
## LICENSE
MIT License

