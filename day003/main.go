package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var language string

	app := &cli.App{
		Name:    "godaily day003",      // cli name
		Version: "v13",                 // cli version
		Usage:   "godaily day003 test", // usage
		Flags: []cli.Flag{ // 接受的 flag
			&cli.StringFlag{ // string
				Name:        "lang",        // flag 名称
				Aliases:     []string{"l"}, // 别名
				Value:       "english",     // 默认值
				Usage:       "language for the greeting",
				Destination: &language, // 指定地址，如果没有可以通过 *cli.Context 的 GetString 获取
				Required:    true,      // flag 必须设置
			},
		},
		Action: func(c *cli.Context) error {
			name := "who"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if language == "chinese" {
				fmt.Println("你好啊", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
