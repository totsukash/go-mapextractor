package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ShotaTotsuka/go-mapextractor/extractor"
)

func main() {
	sampleJSON, err := os.ReadFile("example/sample.json")
	if err != nil {
		panic(err)
	}

	var m map[string]any
	if err = json.Unmarshal(sampleJSON, &m); err != nil {
		panic(err)
	}

	taro := extractor.Get[string](m, "user", "name", "first", "kanji")
	age := extractor.Get[float64](m, "user", "age")

	fmt.Println(taro)     // 太郎
	fmt.Println(int(age)) // 20
}
