package io

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"stocks/types"
)

func Prompt(label string) (string, error) {

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		//Validate:  validate,
	}

	input, err := prompt.Run()

	if err != nil {
		return "", err
	}

	if input == "exit" || input == "" {
		fmt.Printf("Bye!\n")
		os.Exit(0)
	}

	return input, err
}

func Select(stocks types.Results, err error) (int, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ .Source }}?",
		Active:   "\U0001F4B0 {{ .Source.Entity | cyan }} ({{ .Source.Ticker | red }})",
		Inactive: "  {{ .Source.Entity | cyan }} ({{ .Source.Ticker | red }})",
		Selected: "\U0001F4B0 {{ .Source.Entity | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select a stock",
		Items:     stocks.Hits.Hits,
		Templates: templates,
		Size:      4,
	}

	selected, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0, err
	}

	return selected, nil
}
