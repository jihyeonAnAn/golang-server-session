package lab

import "fmt"

type Snack struct {
	Name  string
	Price int
}

type Drink struct {
	Name  string
	Price int
}

func saleAndGetPrice(item interface{}) int {
	switch v := item.(type) {
	case *Snack:
		return v.Price - (v.Price * 10 / 100)
	case *Drink:
		return v.Price - (v.Price * 20 / 100)
	default:
		return 0
	}
}

func Lab_2() {
	chips := Snack{"Pringles", 4000}
	cracker := Snack{"Ace", 2500}

	soda := Drink{"Sprite", 1800}
	coffee := Drink{"TOP", 2700}

	var total int = 0
	total += saleAndGetPrice(&chips)
	total += saleAndGetPrice(&cracker)
	total += saleAndGetPrice(&soda)
	total += saleAndGetPrice(&coffee)

	fmt.Println(total)
}
