package main

import (
	"fmt"

	"github.com/aisk/wizard"
	"github.com/codegangsta/cli"
	"github.com/leancloud/lean-cli/lean/api"
)

func inputAccountInfo() (string, string) {
	var email = new(string)
	var password = new(string)
	wizard.Ask([]wizard.Question{
		{
			Content: "请输入您的邮箱：",
			Input: &wizard.Input{
				Result: email,
				Hidden: false,
			},
		},
		{
			Content: "请输入您的密码：",
			Input: &wizard.Input{
				Result: password,
				Hidden: true,
			},
		},
	})
	return *email, *password
}

func loginAction(c *cli.Context) error {
	var email, password string

	if c.NArg() == 0 {
		email, password = inputAccountInfo()
	} else if c.NArg() == 2 {
		email = c.Args().Get(0)
		password = c.Args().Get(1)
	} else {
		cli.ShowCommandHelp(c, "login")
		return cli.NewExitError("", 1)
	}

	info, err := api.Login(email, password)
	if err != nil {
		return newCliError(err)
	}

	err = api.LoginUSRegion()
	if err != nil {
		return newCliError(err)
	}

	fmt.Println("登录成功：")
	fmt.Printf("用户名: %s\r\n", info.UserName)
	fmt.Printf("邮箱: %s\r\n", info.Email)
	return nil
}
