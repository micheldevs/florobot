package clients

import (
	"encoding/json"

	"github.com/micheldevs/florobot/utils"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var localizer *i18n.Localizer
var bundle *i18n.Bundle

var currentLanguage string = utils.Config("TG_BOT_LANGUAGE")

func Init() {
	bundle = i18n.NewBundle(language.Spanish)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("assets/locale/en/en.json")
	bundle.LoadMessageFile("assets/locale/es/es.json")

	localizer = i18n.NewLocalizer(bundle, GetCurrLang())
}

func GetCurrLang() string {
	if currentLanguage == "" {
		currentLanguage = language.Spanish.String()
	}

	return currentLanguage
}

func Trans(key string) string {
	localizeConfig := i18n.LocalizeConfig{
		MessageID: key,
	}

	localization, _ := localizer.Localize(&localizeConfig)

	return localization
}

func TransWithValues(key string, values map[string]string) string {
	localizeConfig := i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: values,
	}

	localization, _ := localizer.Localize(&localizeConfig)

	return localization
}
