package insulter

import (
	"math/rand"
	"time"
)

// https://github.com/aimxhaisse/fuu/blob/master/cfuu/dictionnary.json
type insultGender struct {
	Prefix, Name, Suffix []string
}

// Insults struct
type Insults struct {
	Prefix     []string
	Man, Woman insultGender
}

// CreateInsultDict generate the insult dict
func CreateInsultDict() Insults {
	rand.Seed(time.Now().UnixNano())
	maleInsults := insultGender{
		Prefix: []string{
			"gros", "con de", "moche", "fils de", "ptit",
			"bite de", "crouton de", "foutre de", "poil de",
			"marchand de", "suceur de", "mangeur de",
			"buveur de", "dealer de", "fabriquant de",
			"rongeur de", "voleur de",
		},
		Name: []string{
			"raton-laveur", "grumly", "lépreux",
			"hamster", "branloman", "chuck norris",
			"morbak", "jésus", "jedi", "haricot", "poireau",
			"triton", "larve", "chnoque",
		},
		Suffix: []string{
			"vert", "orange", "communiste", "capitaliste",
			"immonde", "gaulliste", "jaune", "bleu", "marron",
			"sale", "obèse", "infecte", "poilu", "puant",
			"végétarien", "pédéraste", "végétatif",
			"néandertalien", "de nazareth",
		},
	}
	femaleInsults := insultGender{
		Prefix: []string{
			"grosse", "conne", "moche", "immonde", "ptite",
			"bite de", "crouton de", "foutre de", "poil de",
			"marchande de", "suceuse de", "mangeuse de",
			"buveuse de", "dealeuse de", "rongeuse de",
			"voleuse de",
		},
		Name: []string{
			"conne", "grosse autiste", "vache", "louloute", "tache",
			"punaise", "mamie", "pomme-de-terre", "moule",
			"raclure", "vieille ordure", "larve", "chnoque",
		},
		Suffix: []string{
			"verte", "orange", "communiste", "poilue", "puante",
			"vegetarienne", "moche", "neandertalienne",
		},
	}
	return Insults{
		Prefix: []string{
			"espèce de", "", "putain de",
		},
		Man:   maleInsults,
		Woman: femaleInsults,
	}
}

// GenerateInsult returns an insult
func (insults Insults) GenerateInsult() string {
	globalPrefix := getRandomFromSliceString(insults.Prefix)

	var insultsPerGender insultGender
	if gender := getRandomFromSliceString([]string{"man", "woman"}); gender == "man" {
		insultsPerGender = insults.Man
	} else {
		insultsPerGender = insults.Woman
	}
	prefix := getRandomFromSliceString(insultsPerGender.Prefix)
	name := getRandomFromSliceString(insultsPerGender.Name)
	suffix := getRandomFromSliceString(insultsPerGender.Suffix)

	return globalPrefix + " " + prefix + " " + name + " " + suffix
}

func getRandomFromSliceString(s []string) string {
	return s[rand.Intn(len(s))]
}
