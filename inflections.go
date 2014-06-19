package inflect

var (
	locales       = map[string]*Inflector{}
	DefaultLocale = "en"
)

func Inflections(locale string) *Inflector {
	inflector := locales[locale]

	if inflector == nil {
		inflector = NewInflector(DefaultLocale)
		locales[locale] = inflector
	}

	return inflector
}

func init() {
	var en = Inflections("en")

	en.Plural(`$`, "s")
	en.Plural(`([xz]|[cs]h)$`, "${1}es")
	en.Plural(`([^aeiouy]o)$`, "${1}es")
	en.Plural(`([^aeiouy])y$`, "${1}ies")

	en.Singular(`s$`, "")
	en.Singular(`(ss)$`, "$1")
	en.Singular(`([sxz]|[cs]h)es$`, "${1}")
	en.Singular(`([^aeiouy]o)es$`, "${1}")
	en.Singular(`([^aeiouy])ies$`, "${1}y")

	en.Irregular("child", "children")
	en.Irregular("person", "people")
	en.Irregular("self", "selves")

	en.Uncountable("series")
}
