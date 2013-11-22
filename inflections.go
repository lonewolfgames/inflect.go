package inflect

var locales = map[string]*Inflector{}


func Inflections(locale string) *Inflector{
	inflector := locales[locale]
	
	if inflector == nil {
		inflector = NewInflector("en")
		locales[locale] = inflector
	}
	
	return inflector;
}


func init() {
	var en = Inflections("en")

	en.Plural(`$`, "s")
	en.Plural(`s$`, "s")
	en.Plural(`^(ax|test)is$`, "${1}es")
	en.Plural(`(octop|vir)us$`, "${1}i")
	en.Plural(`(octop|vir)i$`, "${1}i")
	en.Plural(`(alias|status)$`, "${1}es")
	en.Plural(`(bu)s$`, "${1}ses")
	en.Plural(`(buffal|tomat)o$`, "${1}oes")
	en.Plural(`([ti])um$`, "${1}a")
	en.Plural(`([ti])a$`, "${1}a")
	en.Plural(`sis$`, "ses")
	en.Plural(`(?:([^f])fe|([lr])f)$`, "${1}${2}ves")
	en.Plural(`(hive)$`, "${1}s")
	en.Plural(`([^aeiouy]|qu)y$`, "${1}ies")
	en.Plural(`(x|ch|ss|sh)$`, "${1}es")
	en.Plural(`(matr|vert|ind)(?:ix|ex)$`, "${1}ices")
	en.Plural(`^(m|l)ouse$`, "${1}ice")
	en.Plural(`^(m|l)ice$`, "${1}ice")
	en.Plural(`^(ox)$`, "${1}en")
	en.Plural(`^(oxen)$`, "${1}")
	en.Plural(`(quiz)$`, "${1}zes")
	
	en.Singular(`s$`, "")
	en.Singular(`(ss)$`, "${1}")
	en.Singular(`(n)ews$`, "${1}ews")
	en.Singular(`([ti])a$`, "${1}um")
	en.Singular(`((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)(sis|ses)$`, "${1}sis")
	en.Singular(`(^analy)(sis|ses)$`, "${1}sis")
	en.Singular(`([^f])ves$`, "${1}fe")
	en.Singular(`(hive)s$`, "${1}")
	en.Singular(`(tive)s$`, "${1}")
	en.Singular(`([lr])ves$`, "${1}f")
	en.Singular(`([^aeiouy]|qu)ies$`, "${1}y")
	en.Singular(`(s)eries$`, "${1}eries")
	en.Singular(`(m)ovies$`, "${1}ovie")
	en.Singular(`(x|ch|ss|sh)es$`, "${1}")
	en.Singular(`^(m|l)ice$`, "${1}ouse")
	en.Singular(`(bus)(es)?$`, "${1}")
	en.Singular(`(o)es$`, "${1}")
	en.Singular(`(shoe)s$`, "${1}")
	en.Singular(`(cris|test)(is|es)$`, "${1}is")
	en.Singular(`^(a)x[ie]s$`, "${1}xis")
	en.Singular(`(octop|vir)(us|i)$`, "${1}us")
	en.Singular(`(alias|status)(es)?$`, "${1}")
	en.Singular(`^(ox)en`, "${1}")
	en.Singular(`(vert|ind)ices$`, "${1}ex")
	en.Singular(`(matr)ices$`, "${1}ix")
	en.Singular(`(quiz)zes$`, "${1}")
	en.Singular(`(database)s$`, "${1}")
	
	en.Irregular("person", "people")
	en.Irregular("man", "men")
	en.Irregular("woman", "women")
	en.Irregular("child", "children")
	en.Irregular("sex", "sexes")
	en.Irregular("move", "moves")
	en.Irregular("zombie", "zombies")
	en.Irregular("cactus", "cacti")
	
	en.Uncountable("equipment")
	en.Uncountable("information")
	en.Uncountable("rice")
	en.Uncountable("money")
	en.Uncountable("species")
	en.Uncountable("series")
	en.Uncountable("fish")
	en.Uncountable("sheep")
	en.Uncountable("jeans")
	en.Uncountable("police")
}