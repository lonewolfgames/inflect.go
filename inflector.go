package inflect

import "regexp"

type rule struct {
	Regexp   *regexp.Regexp
	Replacer string
}

func newRule(matcher, replacer string) (this *rule) {
	this = new(rule)
	this.Regexp = regexp.MustCompile("(?i)" + matcher)
	this.Replacer = replacer

	return
}

type Inflector struct {
	locale       string
	plurals      []*rule
	singulars    []*rule
	uncountables []string
}

func NewInflector(locale string) *Inflector {
	this := new(Inflector)

	this.locale = locale

	return this
}

func (this *Inflector) Clear() *Inflector {

	this.plurals = nil
	this.singulars = nil
	this.uncountables = nil

	return this
}

func (this *Inflector) Singularize(str string) string {

	return applyRules(str, this.singulars, this)
}

func (this *Inflector) Pluralize(str string) string {

	return applyRules(str, this.plurals, this)
}

func (this *Inflector) Singular(matcher, replacer string) *Inflector {

	this.singulars = append(this.singulars, newRule(matcher, replacer))
	return this
}

func (this *Inflector) Plural(matcher, replacer string) *Inflector {

	this.plurals = append(this.plurals, newRule(matcher, replacer))
	return this
}

func (this *Inflector) Irregular(singular, plural string) *Inflector {

	this.Plural("\\b"+singular+"\\b", plural)
	this.Singular("\\b"+plural+"\\b", singular)
	return this
}

func (this *Inflector) Uncountable(uncountable string) *Inflector {

	this.uncountables = append(this.uncountables, uncountable)
	return this
}

func applyRules(str string, rules []*rule, this *Inflector) string {
	for _, word := range this.uncountables {
		if word == str {
			return str
		}
	}
	var ruleToApply *rule

	for i := len(rules) - 1; i >= 0; i-- {
		ruleToApply = rules[i]

		if ruleToApply.Regexp.MatchString(str) {
			return ruleToApply.Regexp.ReplaceAllString(str, ruleToApply.Replacer)
		}
	}

	return str
}
