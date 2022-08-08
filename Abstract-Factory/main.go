package main

import (
	products "abstractFactory/data"
	"log"
)

func main() {

	forMen, err := products.ForGender("man")
	if err != nil {
		log.Fatal(err)
	}
	forMenJeans := forMen.MakeJeans()
	forMenJeans.SetColor("DarkBlue")
	forMenJeans.SetSize(30)
	log.Printf("A pair of jeans for men, size: %d , color:%s", forMenJeans.GetSize(), forMenJeans.GetColor())

	forWomen, err := products.ForGender("woman")
	if err != nil {
		log.Fatal(err)
	}
	forWomenShirt := forWomen.MakeTShirt()
	forWomenShirt.SetSize(24)
	forWomenShirt.SetColor("Pink")
	log.Printf("A Tshirt for women, size: %d , color:%s", forWomenShirt.GetSize(), forWomenShirt.GetColor())
}
