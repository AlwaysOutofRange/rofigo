# rofigo - Rofi Integration for Command-Line Menus

**rofigo** is a Go package that simplifies the integration of Rofi, a versatile application launcher and window switcher for X11, into your command-line applications. It allows you to create interactive menus with customizable options, labels, and actions, making it easy for users to navigate and interact with your CLI application.

## Overview

- **MenuItem**: Represents an item within a menu. Each MenuItem has a label, an associated action, and an optional icon. You can customize menu items to perform specific actions when selected.

- **Menu**: Represents a collection of menu items along with an optional prompt. Menus help group related items together and provide a user-friendly interface for selecting actions.

- **RofiRunner**: A utility for running Rofi-based menus with custom options. You can create an instance of RofiRunner and execute Rofi menus with specific settings, such as prompts and menu items.

## Example

Here's a simple example of how to use **rofigo** to create and run a menu with Rofi:

```go
// Create a menu
myMenu := rofigo.NewMenu().WithPrompt("Select an option:")

// Add menu items
item1 := rofigo.NewMenuItem("Option 1", func() error {
    // Define the action for Option 1
    // ...
    return nil
}, "")
item2 := rofigo.NewMenuItem("Option 2", func() error {
    // Define the action for Option 2
    // ...
    return nil
}, "")

myMenu.WithItems(item1, item2)

// Create a RofiRunner and run the menu
runner := rofigo.NewRofiRunner(myMenu)
runner.Run()
```

This code creates a menu with two options and executes it using Rofi.
Usage
```
Step 1: Install Rofi on your system if it's not already installed.
Step 2: Install the rofigo pacakge:
"go get github.com/AlwaysOutofRange/rofigo"

Step 3: Import the rofigo package into your Go project:
"import "github.com/AlwaysOutofRange/rofigo""
```

Follow the example above to create and run interactive menus in your command-line application.

## Package Dependencies
rofigo relies on the external tool Rofi, so ensure you have Rofi installed on your system to use this package effectively.
License

## License
This project is licensed under the MIT License - see the LICENSE file for details.

