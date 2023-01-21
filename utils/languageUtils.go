package utils

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func TextLanguage(txt string, lang string) string {
	bundle := i18n.NewBundle(language.Turkish)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("./location/en.json")
	bundle.MustLoadMessageFile("./location/tr.json")
	loc := i18n.NewLocalizer(bundle, lang)
	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: txt,
	})
	return translation
}
