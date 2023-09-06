package main

import (
	"fmt"

	builder "github.com/sabrs0/L2_WB/patterns/builder/builderStuff"
)

func main() {
	lampBuilder := builder.NewLampBuilder()
	director := builder.NewProductDirector()
	director.SetBuilder(lampBuilder)
	defaultLamp := director.Construct()
	fmt.Println(defaultLamp)

	tableBuilder := builder.NewTableBuilder()
	tableBuilder.SetDescription("1x2x1.5")
	director.SetBuilder(tableBuilder)
	customTable := director.Construct()
	fmt.Println(customTable)
}
