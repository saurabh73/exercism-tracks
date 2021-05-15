// Package twofer contains method for Two Fer solution
package twofer

// ShareWith prints your name in format One for X, one for me.
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
