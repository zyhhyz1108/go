package main

import (
	"fmt"
	"time"
	"os"
)



type BusinessError struct {
	Code    int
	Message string
	Time    time.Time
}

const (
	TimeFmt = "2026-03-13 19:46:47"
)

func (e *BusinessError) Error() string {
	return fmt.Sprintf("错误代码：%d,消息：%d,时间：%d", e.Code, e.Message, e.Time.Format(TimeFmt))
}

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err!= nil{
		return fmt.Errorf("打开文件失败：%w", err)
	}
    defer 
}

func safeAccess(arr[]int, index int) (result int, err error) {
	defer func(){
		if r:=recover(); r !=nil{
			err = fmt.Errorf("发生panic:%v, r")
		}
	}()

	result = arr[index]
	return result, nil
}




func main() {
	// defer使用示例
    fmt.Println("\n=== defer使用示例 ===")
    func() {
        defer fmt.Println("1. 第一个defer")
        defer fmt.Println("2. 第二个defer")
        fmt.Println("3. 函数体")
    }()
    
    // panic/recover示例
    fmt.Println("\n=== panic/recover测试 ===")
    arr := []int{1, 2, 3}
    
    // 正常访问
    result, err := safeAccess(arr, 1)
    if err != nil {
        fmt.Println("访问失败:", err)
    } else {
        fmt.Println("访问成功:", result)
    }
    
    // 越界访问
    result, err = safeAccess(arr, 10)
    if err != nil {
        fmt.Println("访问失败:", err)
    } else {
        fmt.Println("访问成功:", result)
    }
    
    fmt.Println("程序继续执行...")

}