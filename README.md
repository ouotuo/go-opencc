go-opencc
=========

opencc wrapper for Golang

fork from 'github.com/stevenyao/go-opencc'

修复一些bug

### Install

Run 'go get github.com/ouotuo/go-opencc'

由于cgo使用pkg-config配置来查找动态库，安装完opencc后，确认安装正确。
```apple js
pkg-config --libs --cflags opencc
```

### Example

```go
configFile:="/usr/local/share/opencc/s2t.json"
c := opencc.NewConverter(configFile)
defer c.Close()

input:="我是中国人"
fmt.Printf("input:%s\n",input)
output := c.Convert(input)
fmt.Printf("output:%s\n",output)

```
