package inflect

import (
	"strings"
	"regexp"
	"strconv"
)

var (
	spliter = regexp.MustCompile(`[\ \-_\.]+`)
	upper_case = regexp.MustCompile(`([A-Z])`)
	underscore_prefix = regexp.MustCompile(`^_`)
	id = regexp.MustCompile(`_id$`)
)


func Pluralize(str, locale string) string {
	
	return Inflections(locale).Pluralize(str)
}


func Singularize(str, locale string) string {
	
	return Inflections(locale).Singularize(str)
}


func Camelize(str string) string{
	strs := spliter.Split(strings.ToLower(str), -1)
	
	for i, str := range strs {
		strs[i] = strings.ToUpper(string(str[0])) + str[1:len(str)]
	}
	
	return strings.Join(strs, "");
}


func Underscore(str string) string{
	
	return underscore_prefix.ReplaceAllString(strings.ToLower(upper_case.ReplaceAllString(str, "_${1}")), "")
}


func Dasherize(str string) string{
	
	return strings.Replace(str, "_", "-", -1)
}


func Humanize(str string) string{
	
	return strings.Join(spliter.Split(id.ReplaceAllString(strings.ToUpper(string(str[0])) + str[1:len(str)], ""), -1), " ")
}


func Titleize(str string) string{
	strs := spliter.Split(strings.ToLower(str), -1)
	
	for i, str := range strs {
		strs[i] = strings.ToUpper(string(str[0])) + str[1:len(str)]
	}
	
	return strings.Join(strs, " ");
}


func Tableize(str, locale string) string{
	
	return Pluralize(Underscore(str), locale);
}


func Classify(str, locale string) string{
	
	return Singularize(Camelize(str), locale);
}


func ForeignKey(str, locale string) string{
	
	return Singularize(Underscore(str), locale) + "_id";
}


func Ordinal(num int) string{
	if num < 0 {
		num = -num;
	}
	num = (num % 100) % 10
	
	switch(num){
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


func Ordinalize(num int) string{
	
	return strconv.Itoa(num) + Ordinal(num)
}