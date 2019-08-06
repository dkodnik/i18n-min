package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	_ "golang.org/x/text/message/catalog"
)

var matcher language.Matcher

func init() {
	matcher = language.NewMatcher(message.DefaultCatalog.Languages())
}

func lng(i18n, phrs string) string {
	lang := []string{}
	if i18n != "" {
		lang = append(lang, i18n)
	}

	if len(lang) < 1 {
		lang = append(lang, language.English.String())
	}
	tag, _, _ := matcher.Match(language.MustParse(lang[0]))

	p := message.NewPrinter(tag)
	retPhr := p.Sprintf(phrs)
	return retPhr
}
