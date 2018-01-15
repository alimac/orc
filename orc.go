// Package orc provides helper methods to generate orc attributes.
//
package orc

import (
	"math/rand"
	"strings"
	"time"
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

// Forge returns the requested orc attribute (name, greeting, or weapon).
func Forge(attribute string) string {
	rand.Seed(time.Now().Unix())

	switch attribute {
	case "name":
		prefix := strings.Title(prefixes[rand.Intn(len(prefixes))])
		suffix := suffixes[rand.Intn(len(suffixes))]
		return prefix + suffix
	case "greeting":
		return strings.Title(greetings[rand.Intn(len(greetings))])
	case "weapon":
		adjective := strings.Title(adjectives[rand.Intn(len(adjectives))])
		weapon := strings.Title(weapons[rand.Intn(len(weapons))])
		return adjective + weapon
	default:
		return ""
	}
}
