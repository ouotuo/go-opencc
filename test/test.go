package main

import "fmt"
import "github.com/ouotuo/go-opencc"
import "time"

const (
	input      = "中国鼠标软件打印机"
	config_s2t = "/usr/local/share/opencc/s2t.json"

	config_t2s = "/usr/local/share/opencc/t2s.json"
)

func main() {
	fmt.Println("Test Converter class:")
	c := opencc.NewConverter(config_s2t)
	defer c.Close()
	output := c.Convert(input)
	fmt.Println(output)

	fmt.Println("Test Convert function:")
	s := opencc.Convert(input, config_s2t)
	fmt.Println(s)
	fmt.Println(opencc.Convert(s, config_t2s))

	fmt.Println("Test ConvertAsync function:")
	retChan := make(chan string)
	opencc.ConvertAsync(input, config_s2t, func(output string) {
		retChan <- output
	})

	fmt.Println(<-retChan)

	//测试速度
	var in = "在windows平台下mingw 编译使用第三方的 C语言库，我需要给mingw设置 环境变量 指示mingw 的 gcc命令编译时 寻找头文件和 库文件的目录"
	bt := time.Now()
	num := 10000
	var out string
	outNum := 0
	for i := 0; i < num; i++ {
		out = c.Convert(in)
		outNum += len(out)
	}
	et := time.Now()
	fmt.Printf("outNum=%d,convert %d times,costTime=%d ns,aver %d ns\n", outNum, num, et.UnixNano()-bt.UnixNano(), (et.UnixNano()-bt.UnixNano())/int64(num))
	fmt.Println(out)
}
