package main

import (
	"fmt"
	"time"

	"github.com/dixneuf19/insult-jmk/insulter"
)

func main() {
	insults := insulter.CreateInsultDict()
	for {
		fmt.Println(insults.GenerateInsult())
		time.Sleep(10 * time.Second)
	}
}
