package LangDetect
type LangDetectorLangItem struct {
	ISO  string
	Name string
}

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


