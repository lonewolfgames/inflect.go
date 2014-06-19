package inflect

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	spliter    = regexp.MustCompile(`[ \_\-\.]+`)
	upperCase  = regexp.MustCompile(`([A-Z])`)
	underscore = regexp.MustCompile(`([a-z\d])([A-Z])`)
	id         = regexp.MustCompile(`_id$`)
)

func Pluralize(str string) string {

	return Inflections(DefaultLocale).Pluralize(str)
}

func PluralizeLocale(str, locale string) string {

	return Inflections(locale).Pluralize(str)
}

func Singularize(str string) string {

	return Inflections(DefaultLocale).Singularize(str)
}

func SingularizeLocale(str, locale string) string {

	return Inflections(locale).Singularize(str)
}

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

func Underscore(str string) string {

	return strings.ToLower(underscore.ReplaceAllString(str, "${1}_${2}"))
}

func Dasherize(str string) string {

	return strings.Replace(str, "_", "-", -1)
}

func Humanize(str string) string {

	return strings.Join(spliter.Split(id.ReplaceAllString(strings.ToUpper(string(str[0]))+str[1:len(str)], ""), -1), " ")
}

func Titleize(str string) string {
	strs := spliter.Split(strings.ToLower(str), -1)

	for i, str := range strs {
		strs[i] = strings.ToUpper(string(str[0])) + str[1:len(str)]
	}

	return strings.Join(strs, " ")
}

func Tableize(str string) string {

	return Pluralize(Underscore(str))
}
func TableizeLocale(str, locale string) string {

	return PluralizeLocale(Underscore(str), locale)
}

func Classify(str string) string {

	return Singularize(Camelize(str, false))
}
func ClassifyLocale(str, locale string) string {

	return SingularizeLocale(Camelize(str, false), locale)
}

func ForeignKey(str string) string {

	return Singularize(Underscore(str)) + "_id"
}
func ForeignKeyLocale(str, locale string) string {

	return SingularizeLocale(Underscore(str), locale) + "_id"
}

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

func Ordinalize(num int) string {

	return strconv.Itoa(num) + Ordinal(num)
}
