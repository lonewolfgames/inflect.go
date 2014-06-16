Inflect.go
=====

A port of the Rails / ActiveSupport inflector to Go

##Example
```
package main

import (
	"github.com/nathanfaucett/inflect"
)

func main() {
	// also PluralizeLocale(string, locale)
	inflect.Pluralize("string") // return "strings"
}
```