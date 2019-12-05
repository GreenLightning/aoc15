package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var input interface{}

	bytes, err := ioutil.ReadFile("input.txt")
	check(err)
	err = json.Unmarshal(bytes, &input)
	check(err)

	{
		fmt.Println("--- Part One ---")
		fmt.Println(count(input, true))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(count(input, false))
	}
}

func count(value interface{}, includeRed bool) (total float64) {
	switch value := value.(type) {
	case float64:
		return value

	case string:
		return 0

	case []interface{}:
		for _, child := range value {
			total += count(child, includeRed)
		}
		return

	case map[string]interface{}:
		if !includeRed {
			for _, child := range value {
				if text, _ := child.(string); text == "red" {
					return 0
				}
			}
		}
		for _, child := range value {
			total += count(child, includeRed)
		}
		return

	default:
		panic("unhandled type")
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
