package main

import "fmt"

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	//购物车
	var cart []CartItem
	cart = append(cart, CartItem{"笔记本电脑", 5999.00, 1})
	cart = append(cart, CartItem{"鼠标", 1999.00, 1})
	cart = append(cart, CartItem{"键盘", 599.00, 1})
	total := 0.0
	for _, i := range cart {
		total += i.Price * float64(i.Quantity)
	}
	fmt.Printf("总价为%.2f", total)
}
