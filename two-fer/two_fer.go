package twofer

import "fmt"

func ShareWith(name string) string {

	var twofer string

	switch len(name) {
	case 0:
		twofer = "One for you, one for me."

	default:
		twofer = fmt.Sprintf("One for %s, one for me.", name)

	}
	return twofer
}
