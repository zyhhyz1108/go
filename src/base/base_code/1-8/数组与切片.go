package main

import "fmt"

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func processBatch(batch []int) {
	fmt.Printf("处理批次，大小：%d\n", len(batch))
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
	//顶分配容量优化性能
	fmt.Printf("总价为%.2f", total)
	ids := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		ids = append(ids, i)
	}
	//切片复制避免内存泄漏
	importantData := make([]byte, 0, 10*1024*1024)
	returnDate := make([]byte, len(importantData))
	copy(returnDate, importantData)
	//批量处理
	data := make([]int, 1000)
	batchSize := 100
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		batch := data[i:end]
		processBatch(batch)
		//多维数组
		matrix := make([][]int, 3)
		for i := range matrix {
			matrix[i] = make([]int, 3)
		}
	}
}
