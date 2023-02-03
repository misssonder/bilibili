package main

import (
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
)

var ins = spinner.New(spinner.CharSets[35], 100*time.Millisecond)

func selectList(title string, items []string) (int, error) {
	question := &survey.Select{
		Message: title,
		Options: items,
		Filter: func(filter string, value string, index int) bool {
			return inText(strings.ToLower(filter), strings.ToLower(value))
		},
	}
	var idx int
	err := survey.AskOne(question, &idx)
	if err != nil {
		return 0, err
	}
	return idx, nil
}

func multipleSelectList(title string, items []string) ([]int, error) {
	question := &survey.MultiSelect{
		Message: title,
		Options: items,
		Filter: func(filter string, value string, index int) bool {
			return inText(strings.ToLower(filter), strings.ToLower(value))
		},
	}
	var idxs = make([]int, 0, len(items))
	err := survey.AskOne(question, &idxs)
	if err != nil {
		return nil, err
	}
	return idxs, nil
}

// inText a in `aaa
// abc in `aaa`bbb`ccc
// 你好世界 in `你`好啊 ，哈哈，这个`世`界
func inText(key, text string) bool {
	// textMin := 0
	// textMax := len(text) - 1
	matchCount := 0
	keyIndex := 0
	textIndex := 0
	for keyIndex < len(key) && keyIndex < len(text) && textIndex < len(text) {
		// if keyIndex > textMax {
		// 	break
		// }

		keyRune := key[keyIndex]
		textRune := text[textIndex]
		// fmt.Println(keyIndex,textIndex,)
		if keyRune == textRune {
			matchCount++
			keyIndex++
		}
		textIndex++
	}

	return len(key) == matchCount
}
