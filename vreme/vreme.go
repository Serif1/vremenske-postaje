
// Package vreme omogoča shranjevanje temperaturnih meritev
// vremenskih postaj in izračun statistike.
package vreme

import (
	"fmt"
	"sort"
)

// Postaja predstavlja vremensko postajo in njene meritve.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda novo vremensko postajo v slovar postaj.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}

// DodajMeritev doda temperaturo postaji, če postaja obstaja
// in je temperatura v dovoljenem območju.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	p, obstaja := postaje[oznaka]

	if !obstaja {
		fmt.Println("Postaja", oznaka, "ne obstaja.")
		return
	}

	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Neveljavna temperatura:", temperatura)
		return
	}

	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p
	fmt.Println("Dodana meritev za", oznaka, ":", temperatura)
}

// statistika izračuna najnižjo, najvišjo in povprečno temperaturo.
// Če ni dovolj meritev, vrne false.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	p, obstaja := postaje[oznaka]

	if !obstaja || len(p.meritve) < stMeritev || len(p.meritve) == 0 {
		return 0, 0, 0, false
	}

	najnizja = p.meritve[0]
	najvisja = p.meritve[0]
	vsota := 0.0

	for _, temp := range p.meritve {
		if temp < najnizja {
			najnizja = temp
		}
		if temp > najvisja {
			najvisja = temp
		}
		vsota += temp
	}

	return najnizja, najvisja, vsota / float64(len(p.meritve)), true
}

// oznakePostaj vrne oznake postaj, urejene po abecedi.
func oznakePostaj(postaje map[string]Postaja) []string {
	oznake := make([]string, 0, len(postaje))

	for oznaka := range postaje {
		oznake = append(oznake, oznaka)
	}

	sort.Strings(oznake)
	return oznake
}

// IzpisMeritev izpiše vse postaje in njihove meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("\nMERITVE:")

	for _, oznaka := range oznakePostaj(postaje) {
		p := postaje[oznaka]
		fmt.Printf("%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

// IzpisPorocila izpiše statistiko in opis temperatur za vsako postajo.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	fmt.Println("\nPOROCILO:")

	for _, oznaka := range oznakePostaj(postaje) {
		p := postaje[oznaka]
		min, max, povp, ok := statistika(postaje, oznaka, stMeritev)

		if !ok {
			fmt.Printf("%s: premalo meritev\n", p.ime)
			continue
		}

		opis := ""

		switch {
		case povp >= 25:
			opis = "Vroče"
		case povp >= 15:
			opis = "Toplo"
		case povp >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n",
			p.ime, min, max, povp, opis)
	}
}
