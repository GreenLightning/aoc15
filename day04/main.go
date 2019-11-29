package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	key := readFile("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(mineAdventCoin(key, "00000"))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(mineAdventCoin(key, "000000"))
	}
}

func mineAdventCoin(key string, prefix string) int {
	for i := 0; ; i++ {
		data := fmt.Sprintf("%s%d", key, i)
		bytes := md5.Sum([]byte(data))
		hash := hex.EncodeToString(bytes[:])
		if strings.HasPrefix(hash, prefix) {
			return i
		}
	}
}

func readFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
