package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
        "github.com/gobs/readline"
	"log"
)

type Options struct {
	ScreenName string `short:"a" long:"account" description:"login as an account of selected screen_name"`
}

func main() {
	account := loadAccount()

	startUserStream(account)
	invokeInteractiveShell(account)
}

func loadAccount() *Account {
	options := new(Options)
	if _, err := flags.Parse(options); err != nil {
		log.Fatal(err)
	}

	if len(options.ScreenName) > 0 {
		return AccountByScreenName(options.ScreenName)
	} else {
		return DefaultAccount()
	}
}

func invokeInteractiveShell(account *Account) {

	for {
		currentLine := readline.ReadLine(prompt(account))
		if currentLine == nil || *currentLine == ":exit" {
			return
		}

		executeCommand(account, *currentLine)
		readline.AddHistory(*currentLine)
	}
}

func prompt(account *Account) *string {
	prompt := fmt.Sprintf("[%s] ", coloredScreenName(account.ScreenName))
	return &prompt
}
