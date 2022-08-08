package main

import (
	sweets "builderPattern/data"
	"log"
)

func main() {
	baker := &sweets.Baker{}
	applePieToBake, err := baker.GetToBake("APPLEPIE")
	if err != nil {
		log.Panicln(err)
	}
	pancakeToBake, err := baker.GetToBake("APPLEPIE")
	if err != nil {
		log.Panicln(err)
	}
	applePie := applePieToBake.AddEggs(1).AddMilk(1).AddSugar(1).Bake()
	pancake := pancakeToBake.AddEggs(1).AddMilk(3).AddSugar(1).Bake()

	log.Println(applePie.GetRecipe())
	log.Println(pancake.GetRecipe())
}
