package main

import (
	"fmt"
	"github.com/ouotuo/go-opencc"
	"flag"
)

const (
	input      = "中国鼠标软件打印机"
	config_s2t = "s2t.json"
	config_t2s = "t2s.json"
)

//go run test.go -c /usr/local/share/opencc/s2t.json
func main() {
	configFile:=flag.String("c","","convert config file")
	flag.Parse()

	if *configFile==""{
		fmt.Printf("config file empty\n")
		return
	}

	fmt.Printf("config file %s\n",*configFile)

	c := opencc.NewConverter(*configFile)
	defer c.Close()

	fmt.Printf("input:%s\n",input)
	output := c.Convert(input)
	fmt.Printf("output:%s\n",output)
}
