package main

import "fmt"

func main() {

	iphone := Iphone{Price: 999}

	iphone13 := &Iphone13{
		phone: &iphone,
	}

	iphone14 := &Iphone14{
		phone: iphone13,
	}

	fmt.Println(iphone13.getPrice())
	fmt.Println(iphone14.getPrice())
}

type Phone interface {
	getPrice() int
}

///

type Iphone struct {
	Price int
}

func (p *Iphone) getPrice() int {
	return p.Price
}

///

type Iphone13 struct {
	phone Phone
}

func (c *Iphone13) getPrice() int {
	price := c.phone.getPrice()
	return price + 100
}

///

type Iphone14 struct {
	phone Phone
}

func (c *Iphone14) getPrice() int {
	price := c.phone.getPrice()
	return price + 200
}
