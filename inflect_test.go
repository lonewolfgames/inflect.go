package inflect

import (
	"testing"
)

func check(t *testing.T, actual, expected string) {
	if expected != actual {
		t.Error("Failed Test got " + actual + " expected " + expected)
	}
}

func testPluralize(t *testing.T, test, expected string) {
	actual := Pluralize(test)
	check(t, actual, expected)
}
func TestPluralize(t *testing.T) {
	testPluralize(t, "box", "boxes")
	testPluralize(t, "cart", "carts")
	testPluralize(t, "man", "men")
	testPluralize(t, "sky", "skies")
	testPluralize(t, "police", "police")
}

func testPluralizeLocale(t *testing.T, test, expected string) {
	actual := PluralizeLocale(test, "en")
	check(t, actual, expected)
}
func TestPluralizeLocale(t *testing.T) {
	testPluralizeLocale(t, "box", "boxes")
	testPluralizeLocale(t, "cart", "carts")
	testPluralizeLocale(t, "man", "men")
	testPluralizeLocale(t, "sky", "skies")
	testPluralizeLocale(t, "police", "police")
}

func testSingularize(t *testing.T, test, expected string) {
	actual := Singularize(test)
	check(t, actual, expected)
}
func TestSingularize(t *testing.T) {
	testSingularize(t, "boxes", "box")
	testSingularize(t, "carts", "cart")
	testSingularize(t, "men", "man")
	testSingularize(t, "skies", "sky")
	testSingularize(t, "police", "police")
}

func testSingularizeLocale(t *testing.T, test, expected string) {
	actual := SingularizeLocale(test, "en")
	check(t, actual, expected)
}
func TestSingularizeLocale(t *testing.T) {
	testSingularizeLocale(t, "boxes", "box")
	testSingularizeLocale(t, "carts", "cart")
	testSingularizeLocale(t, "men", "man")
	testSingularizeLocale(t, "skies", "sky")
	testSingularizeLocale(t, "police", "police")
}

func TestCamelize(t *testing.T) {
	check(t, Camelize("test_of test", false), "TestOfTest")
	check(t, Camelize("test_of test", true), "testOfTest")
}

func TestUnderscore(t *testing.T) {
	check(t, Underscore("TestOfTest"), "test_of_test")
}

func TestDasherize(t *testing.T) {
	check(t, Dasherize("test_of_test"), "test-of-test")
}

func TestHumanize(t *testing.T) {
	check(t, Humanize("test_of test"), "Test of test")
}

func TestTitleize(t *testing.T) {
	check(t, Titleize("test_of test"), "Test Of Test")
}

func TestTableize(t *testing.T) {
	check(t, Tableize("TestTest"), "test_tests")
}
func TestTableizeLocale(t *testing.T) {
	check(t, TableizeLocale("TestTest", "en"), "test_tests")
}

func TestClassify(t *testing.T) {
	check(t, Classify("test_tests"), "TestTest")
}
func TestClassifyLocale(t *testing.T) {
	check(t, ClassifyLocale("test_tests", "en"), "TestTest")
}

func TestForeignKey(t *testing.T) {
	check(t, ForeignKey("TestTests"), "test_test_id")
}
func TestForeignKeyLocale(t *testing.T) {
	check(t, ForeignKeyLocale("TestTests", "en"), "test_test_id")
}

func TestOrdinal(t *testing.T) {
	check(t, Ordinal(1), "st")
	check(t, Ordinal(2), "nd")
	check(t, Ordinal(3), "rd")
	check(t, Ordinal(4), "th")
}

func TestOrdinalize(t *testing.T) {
	check(t, Ordinalize(1), "1st")
	check(t, Ordinalize(2), "2nd")
	check(t, Ordinalize(3), "3rd")
	check(t, Ordinalize(4), "4th")
}
