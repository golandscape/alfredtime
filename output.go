package main

import (
	"encoding/json"
	"fmt"
)

func output(title, arg string, subtext string) {
	marshal, _ := json.Marshal(Items{
		Items: []Item{{
			Title:    title,
			Arg:      arg,
			Subtitle: subtext,
		}},
	})
	fmt.Println(string(marshal))
}
