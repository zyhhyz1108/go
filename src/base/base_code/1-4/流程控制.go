package main

import (
	"fmt"
	"time"
)

func main() {
	score := 98
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("需要努力")
	}
	//批量处理订单
	orders := []string{"订单1", "订单2"}
	for i, order := range orders {
		fmt.Printf("处理第%d个订单：%s\n", i+1, order)
	}
	//http状态码处理
	statusCode := 200
	switch statusCode {
	case 200:
		fmt.Println("请求成功")
	case 400:
		fmt.Println("页面未找到")
	default:
		fmt.Println("其他状态码")
	}
	//超时控制
	ch := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "任务完成"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("任务超时")
	}
}
