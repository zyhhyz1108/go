package main

import "fmt"

// 定义User结构体
type User struct {
	ID      int
	Name    string
	Balance float64
}

// 【关键修改】值传递函数：返回修改后的User副本
func updateBalanceValue(user User, amount float64) User {
	user.Balance += amount // 修改的是副本
	fmt.Println("updateBalanceValue 更新后的余额:", user.Balance)
	return user // 返回修改后的副本
}

func main() {
	// 初始化用户
	user := User{ID: 1, Name: "张三", Balance: 100.0}

	// 【核心步骤】用值传递函数的返回值覆盖原变量
	user = updateBalanceValue(user, 50.0)
	// 此时原变量已经被替换成“修改后的副本”，余额变为150
	fmt.Println("updateBalanceValue ", user) // 输出：{1 张三 150}

	users := []User{
		{ID: 1, Name: "用户1", Balance: 100},
		{ID: 2, Name: "用户2", Balance: 200},
		{ID: 3, Name: "用户3", Balance: 300},
	}

	// 批量充值
	for _, u := range users {
		u.Balance += 50 // 直接通过指针修改
	}

	for _, u := range users {
		fmt.Printf("%s的新余额: %.2f\n", u.Name, u.Balance)
	}
	//二级指针

}
