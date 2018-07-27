package main

import (
	"fmt"

	"github.com/xrash/smetrics"
)

func main() {
	fmt.Println(smetrics.Jaro("componente curricular", "disciplina"))
}
