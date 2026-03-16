package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ValidateUser(username, password string) (bool, error) {
	if len(username) < 3 {
		return false, fmt.Errorf("用户名长度不能小于3")
	}
	if len(password) < 6 {
		return false, fmt.Errorf("密码长度不能小于6")
	}
	return true, nil
}

func QueryUser(userID int) (string, int, error) {
	if userID <= 0 {
		return "", 0, errors.New("用户ID不能小于等于0")
	}
	users := map[int]struct {
		Name string
		Age  int
	}{
		1: {"张一", 20},
		2: {"张er", 21},
		3: {"张san", 22},
	}
	if user, exist := users[userID]; exist {
		return user.Name, user.Age, nil
	}
	return "", 0, errors.New("用户不存在")
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 定义函数类型
type StringProcessor func(string) string

// 接收函数作为参数
func processString(s string, processor StringProcessor) string {
	return processor(s)
}

// 用户注册验证函数
func main() {
	_, err := ValidateUser("admin", "123456")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("用户注册成功")
	}

	username, age, err := QueryUser(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("用户信息", username, age)
	}

	// 闭包使用
	counter := createCounter()
	fmt.Println("计数:", counter()) // 1
	fmt.Println("计数:", counter()) // 2

	// 匿名函数：字符串处理
	strs := []string{"apple", "Banana", "cherry"}
	fmt.Println("原始:", strs)

	// 使用匿名函数排序（忽略大小写）
	sort.Slice(strs, func(i, j int) bool {
		return strings.ToLower(strs[i]) < strings.ToLower(strs[j])
	})
	fmt.Println("排序后:", strs)

	//函数作为参数传递
	toUpper := func(s string) string {
		return strings.ToUpper(s)
	}
	result := processString("hello go", toUpper)
	fmt.Println(result) // HELLO GO

	//直接传递匿名函数
	result2 := processString("  hello     ", func(s string) string {
		return strings.TrimSpace(s)
	})
	fmt.Printf("'%s'\n", result2) //'hello'

}
