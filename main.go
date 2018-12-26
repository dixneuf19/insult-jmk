package main

import (
	"fmt"

	"github.com/dixneuf19/insult-jmk/insulter"
)

func main() {
	insults := insulter.CreateInsultDict()
	fmt.Println(insults.GenerateInsult())
}
