// models/option.go
package models

// Option represents a selectable option in the interactive menu
type Option struct {
	Name        string
	Description string
	Action      func() error
}

// OptionList is a collection of options
type OptionList struct {
	Options []Option
}

// NewOptionList creates a new option list
func NewOptionList() *OptionList {
	return &OptionList{
		Options: []Option{},
	}
}

// AddOption adds a new option to the list
func (ol *OptionList) AddOption(name, description string, action func() error) {
	ol.Options = append(ol.Options, Option{
		Name:        name,
		Description: description,
		Action:      action,
	})
}

// GetOption returns an option by index
func (ol *OptionList) GetOption(index int) *Option {
	if index < 0 || index >= len(ol.Options) {
		return nil
	}
	return &ol.Options[index]
}

// GetOptionNames returns all option names
func (ol *OptionList) GetOptionNames() []string {
	names := make([]string, len(ol.Options))
	for i, opt := range ol.Options {
		names[i] = opt.Name
	}
	return names
}
