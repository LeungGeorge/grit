# 用法

配置文件 `config.json` 内容：  
```html
{
    "model":"aaaaa"
}
```

demo：  
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/LeungGeorge/grit/lib/hotreload"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
)

type Config struct {
	Model string `json:"model"`
}

var cnf Config

func loadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	err = jsoniter.Unmarshal(file, &cnf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("config:%+v\n", cnf)

	return nil
}

func main() {
	r := gin.Default()

	hotreload.Register("conf/config.json", loadConfig)
	hotreload.Watcher()

	r.Run()
}
```

输出（可以看到 `demo` 读取到 `conf/config.json` 的内容）：  
```html
config:{Model:122}
```