package inflect

import "regexp"

type Rule struct{
	Regexp   *regexp.Regexp
    Replacer string     
}

func NewRule(matcher, replacer string) (this *Rule){
	this = new(Rule)
	this.Regexp = regexp.MustCompile(matcher)
	this.Replacer = replacer
	
	return
}

type Inflector struct{
	Locale string
	Plurals []*Rule
	Singulars []*Rule
	Uncountables []string
}


func NewInflector(locale string) (this *Inflector){
	this = new(Inflector)
	
	this.Locale = locale;
	this.Plurals = make([]*Rule, 0)
	this.Singulars = make([]*Rule, 0)
	this.Uncountables = make([]string, 0)
	
	return
}


func (this *Inflector) Singularize(str string) string{
	
	return applyRule(str, this.Singulars, this)
}


func (this *Inflector) Pluralize(str string) string{
	
	return applyRule(str, this.Plurals, this)
}


func (this *Inflector) Singular(matcher, replacer string) *Inflector{
	
	this.Singulars = append(this.Singulars, NewRule(matcher, replacer))
	return this
}


func (this *Inflector) Plural(matcher, replacer string) *Inflector{
	
	this.Plurals = append(this.Plurals, NewRule(matcher, replacer))
	return this
}


func (this *Inflector) Irregular(singular, plural string) *Inflector{
	
	this.Plural("\\b" + singular + "\\b", plural);
    this.Singular("\\b" + plural + "\\b", singular);
	return this
}


func (this *Inflector) Uncountable(uncountable string) *Inflector{
	
	this.Uncountables = append(this.Uncountables, uncountable)
	return this
}


func applyRule(str string, rules []*Rule, this *Inflector) string{
	for _, word := range this.Uncountables {
		if word == str {
			return str
		}
	}
	var rule *Rule
	
	for i := len(rules)-1; i >= 0; i-- {
		rule = rules[i]
		
		if rule.Regexp.MatchString(str) {
			return rule.Regexp.ReplaceAllString(str, rule.Replacer)
		}
	}
	
	return str
}

