package models

// Option struct representing an item in the list
type Option struct {
	Text   string
	Action func()
}
