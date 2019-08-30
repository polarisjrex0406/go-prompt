package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/c-bata/go-prompt"
)

func main() {

	p := prompt.New(executer, completer, options()...)
	go p.Run()

	startSysCallChannel()
}

func executer(input string) {

	fmt.Println("You selected " + input)

	// here you can do whatever you want :)

	if input == "exit" {
		os.Exit(0)
	}
}

func completer(d prompt.Document) []prompt.Suggest {

	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
		{Text: "quotes", Description: "Store the text commented to articles"},
		{Text: "workshops", Description: "Store the text commented to articles"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func options() []prompt.Option {

	return []prompt.Option{
		prompt.OptionTitle("go-prompt-completer"),
		prompt.OptionMaxSuggestion(3),
		prompt.OptionPrefix("$ "),
		prompt.OptionPrefixTextColor(prompt.Turquoise),
		// prompt.OptionPrefixBackgroundColor(prompt.DarkGray),
		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionSuggestionBGColor(prompt.DarkGray),
		prompt.OptionPreviewSuggestionTextColor(prompt.Green),
		// prompt.OptionPreviewSuggestionBGColor(prompt.DarkGray),
		prompt.OptionInputTextColor(prompt.DarkGreen),
		// prompt.OptionInputBGColor(prompt.DarkGray),
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionDescriptionBGColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSelectedDescriptionTextColor(prompt.Black),
		prompt.OptionSelectedDescriptionBGColor(prompt.LightGray),
	}
}

func startSysCallChannel() {

	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT)
	<-syscallCh
}
