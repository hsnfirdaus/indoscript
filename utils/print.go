package utils

import "encoding/json"

func PrintJson(content interface{}) {
	text, _ := json.MarshalIndent(content, "", "\t")
	println(string(text))
}
