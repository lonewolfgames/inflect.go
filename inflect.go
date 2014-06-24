package inflect

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	spliter    = regexp.MustCompile(`[ \_\-\.]+`)
	underscore = regexp.MustCompile(`([a-z\d])([A-Z])`)
	id         = regexp.MustCompile(`_id$`)
)

// Pluralizes string using inflect.DefaultLocale as the locale
func Pluralize(str string) string {

	return Inflections(DefaultLocale).Pluralize(str)
}

// Pluralizes string using passed locale
func PluralizeLocale(str, locale string) string {

	return Inflections(locale).Pluralize(str)
}

// Singularizes string using inflect.DefaultLocale as the locale
func Singularize(str string) string {

	return Inflections(DefaultLocale).Singularize(str)
}

// Singularizes string using passed locale
func SingularizeLocale(str, locale string) string {

	return Inflections(locale).Singularize(str)
}

// Camelizes string "camelizes string" -> "CamelizesString"
func Camelize(str string, lowerFirstLetter bool) string {
	strs := spliter.Split(strings.ToLower(str), -1)

	for i, str := range strs {
		strs[i] = strings.ToUpper(string(str[0])) + str[1:len(str)]
	}
	str = strings.Join(strs, "")

	if lowerFirstLetter == false {
		return str
	}

	return strings.ToLower(string(str[0])) + str[1:len(str)]
}

// Underscores string "UnderscoreString" -> "underscore_string"
func Underscore(str string) string {

	return strings.ToLower(underscore.ReplaceAllString(str, "${1}_${2}"))
}

// Dasherizes string "dasherize_string" -> "dasherize-string"
func Dasherize(str string) string {

	return strings.Replace(str, "_", "-", -1)
}

// Humanizes string "dasherize_string" -> "dasherize-string"
func Humanize(str string) string {

	return strings.Join(spliter.Split(id.ReplaceAllString(strings.ToUpper(string(str[0]))+str[1:len(str)], ""), -1), " ")
}

// Titleizes string "my good_title" -> "My Good Title"
func Titleize(str string) string {
	strs := spliter.Split(strings.ToLower(str), -1)

	for i, str := range strs {
		strs[i] = strings.ToUpper(string(str[0])) + str[1:len(str)]
	}

	return strings.Join(strs, " ")
}

// Tableizes string "Item" -> "items"
func Tableize(str string) string {

	return Pluralize(Underscore(str))
}

// Tableizes string with locale  "Item", "en" -> "items"
func TableizeLocale(str, locale string) string {

	return PluralizeLocale(Underscore(str), locale)
}

// Classify string "items" -> "Item"
func Classify(str string) string {

	return Singularize(Camelize(str, false))
}

// Classify string with locale  "items", "en" -> "Item"
func ClassifyLocale(str, locale string) string {

	return SingularizeLocale(Camelize(str, false), locale)
}

// returns Singularized underscored string with _id "Items" -> "item_id"
func ForeignKey(str string) string {

	return Singularize(Underscore(str)) + "_id"
}

// returns Singularized underscored string with _id using locale "Items", "en" -> "item_id"
func ForeignKeyLocale(str, locale string) string {

	return SingularizeLocale(Underscore(str), locale) + "_id"
}

// returns Ordinal of number 123 -> "rd"
func Ordinal(num int) string {
	if num < 0 {
		num = -num
	}
	num = (num % 100) % 10

	switch num {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

// returns Ordinalize number 123 -> "123th"
func Ordinalize(num int) string {

	return strconv.Itoa(num) + Ordinal(num)
}
