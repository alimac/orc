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

var greetings = []string{
	"bubu",
	"dabu",
	"for the glory of the warchief",
	"for the horde",
	"lok'tar",
	"my life for the horde",
	"okidoki",
	"something need doing?",
	"swobu",
	"time for killing",
	"time to die",
	"work, work",
	"yes, chieftain?",
	"yes, warchief",
	"zub zub",
	"zug zug",
}

var weapons = []string{
	"arrow",
	"blade",
	"bomb",
	"chain",
	"grenade",
	"hammer",
	"kettle",
	"knife",
	"paperweight",
	"plier",
	"rake",
	"scoop",
	"shovel",
	"sickle",
	"thrower",
	"trowel",
	"tweezer",
	"wrench",
}

var adjectives = []string{
	"dark",
	"death",
	"dismal",
	"doom",
	"dreary",
	"dull",
	"fire",
	"gloom",
	"killing",
	"murderous",
	"noxious",
	"ruthless",
	"soul",
	"war",
}

func getRandom(slice []string) string {
	return slice[rand.Intn(len(slice)-1)]
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
	prefix := uppercaseFirstLetter(getRandom(prefixes))
	suffix := lowercaseFirstLetter(getRandom(suffixes))
	return prefix + suffix
}

// GenerateGreeting is returns a random Orc greeting
func GenerateGreeting() string {
	rand.Seed(time.Now().UnixNano())
	return uppercaseFirstLetter(getRandom(greetings))
}

// GenerateWeapon is returns a random Orc weapon
func GenerateWeapon() string {
	rand.Seed(time.Now().UnixNano())
	adjective := uppercaseFirstLetter(getRandom(adjectives))
	weapon := uppercaseFirstLetter(getRandom(weapons))
	return adjective + weapon
}
