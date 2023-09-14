package main

import (
	"log"
	"os"
	"text/template"
)

type hotels struct {
	Name []string
}

type california struct {
	Hotelzzz hotels
	Address  string
	City     string
	Zip      int
	Region   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tmp.gohtml"))
}

func main() {

	cali := []california{
		california{
			Hotelzzz: hotels{[]string{"Alabama Hotel", "Becroft Lux Hotel", "Chidera Spa Hotel", "Dubem Fancy Hotel"}},
			Address:  "No 4, vet road",
			City:     "Los Angeles",
			Zip:      40892,
			Region:   "Southern",
		},
		california{
			Hotelzzz: hotels{[]string{"Nzaza Hotel", "Joha Lux Hotel", "Ototo Spa Hotel", "Sunmomi Fancy Hotel"}},
			Address:  "No 9, Odin road",
			City:     "Bel Air",
			Zip:      97802,
			Region:   "Northern",
		},
		california{
			Hotelzzz: hotels{[]string{"Dull Hotel", "Terminator Lux Hotel", "Organise Spa Hotel", "PBUG Fancy Hotel"}},
			Address:  "No 9, Laos road",
			City:     "Atlanta",
			Zip:      50923,
			Region:   "Central",
		},
	}

	err := tpl.Execute(os.Stdout, cali)
	if err != nil {
		log.Fatalln(err)
	}
}
