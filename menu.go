package rofigo

// MenuItem represents an item within a menu. Each MenuItem has a label, an associated action, and an optional icon.
type MenuItem struct {
	Label  string
	Action func() error
	Icon   string
}

// Menu represents a collection of menu items along with an optional prompt.
type Menu struct {
	Prompt string
	Items  []*MenuItem
}

// NewMenuItem creates a new MenuItem with the specified label, action, and icon.
func NewMenuItem(label string, action func() error, icon string) *MenuItem {
	return &MenuItem{
		Label:  label,
		Action: action,
		Icon:   icon,
	}
}

// NewMenu creates a new Menu with an empty prompt and an empty list of items.
func NewMenu() *Menu {
	return &Menu{
		Prompt: "",
		Items:  make([]*MenuItem, 0),
	}
}

// WithPrompt sets the prompt text for the menu and returns the modified Menu.
func (m *Menu) WithPrompt(prompt string) *Menu {
	m.Prompt = prompt

	return m
}

// WithItem adds a single MenuItem to the Menu and returns the modified Menu. You can use this method to
// customize the menu by adding individual items.
func (m *Menu) WithItem(item *MenuItem) *Menu {
	m.Items = append(m.Items, item)

	return m
}

// method to add multiple items to the menu in a single call.
func (m *Menu) WithItems(items ...*MenuItem) *Menu {
	for _, item := range items {
		m.Items = append(m.Items, item)
	}

	return m
}
