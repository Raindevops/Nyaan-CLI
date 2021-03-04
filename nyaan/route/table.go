package route

import (
	"flag"
	"fmt"
)

// Args : List of available arguments from the command line
var Args = [3]string{"create", "update", "delete"}

// Table : route table for cli
func Table() {

	text := flag.String("text", "", "Text to parse. (Require)")
	metric := flag.String("metric", "chars", "Metric {chars|words|line}")
	unique := flag.Bool("unique", false, "Measure unique value of a metric")

	flag.Parse()

	// if len(*text) == 0 {
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	fmt.Printf("text: %s, metric: %s, unique: %t", *text, *metric, *unique)
}
