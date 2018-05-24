package main

import "fmt"

type Item struct {
	id       string
	price    float64
	quantity int
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item
	catalogId int
}

type LuxuryItem struct {
	Item
	markup float64
}

func (item *LuxuryItem) Cost() float64 {
	return item.Item.price * float64(item.quantity) * item.markup
}


func (item *LuxuryItem) Cost()float64{
	return item.price*float64(item.Item.quantity) * item.markup
}


func (item *LuxuryItem) Cost() float64{
	return  item.Item.Cost() * item.markup
}


func main() {
	special := SpecialItem{Item{"green", 3, 5}, 207}
	fmt.Println(special.id, special.price, special.quantity, special.catalogId)
	fmt.Println(special.Cost())
}
