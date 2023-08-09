package display

import "fmt"

// ShowSimpleList show dates with simple style
func ShowSimpleList(list []string) {
	if len(list) == 0 {
		fmt.Printf("item no detected\n")
	}
	for i, item := range list {
		fmt.Printf("%02d : %s\n", i+1, item)
	}
}
