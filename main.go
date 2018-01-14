package orc

import (
	"math/rand"
	"time"
	"unicode"
)

var prefixes = []string{
	"pig",
	"shmal",
	"snog",
	"urk",
	"vid",
	"wu",
	"gla",
	"shuf",
	"mur",
	"yaz",
	"gha",
	"mor",
	"gul",
	"rog",
}
var suffixes = []string{
	"ak",
	"dug",
	"gagh",
	"gulg",
	"hat",
	"lar",
	"rug",
	"sha",
	"snog",
	"thu",
	"rog",
	"oth",
	"shuk",
}

func randomPrefix() string {
	return prefixes[rand.Intn(len(prefixes)-1)]
}

func randomSuffix() string {
	return suffixes[rand.Intn(len(suffixes)-1)]
}

func uppercaseFirstLetter(word string) string {
	a := []rune(word)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

func lowercaseFirstLetter(word string) string {
	a := []rune(word)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

// GenerateName is returns a random Orc name
func GenerateName() string {
	rand.Seed(time.Now().UnixNano())
	prefix := uppercaseFirstLetter(randomPrefix())
	suffix := lowercaseFirstLetter(randomSuffix())
	return prefix + suffix
}
