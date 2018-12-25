package main

import (
	"fmt"

	"github.com/dixneuf19/insult-jmk/insult_generator"
)

func main() {
	insults := insultGenerator.CreateInsultDict()
	fmt.Println(insults.GenerateInsult())
}
