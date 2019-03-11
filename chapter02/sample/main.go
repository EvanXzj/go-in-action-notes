package main

import (
	_ "github.com/EvanXzj/go-in-action-notes/chapter02/sample/matchers"
	"github.com/EvanXzj/go-in-action-notes/chapter02/sample/search"
	"log"
	"os"
)

func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdin)
}

func main() {
	// 使用特定的项做搜索
	search.Run("president")
}
