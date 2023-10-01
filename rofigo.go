// Package rofigo provides a simple integration with the Rofi menu system for creating interactive command-line menus.

// Rofi is a versatile application launcher and window switcher for X11. The "rofigo" package allows you to easily
// create and display interactive menus in your command-line applications using Rofi. With "rofigo," you can define
// custom menus with items, labels, icons, and actions, providing a user-friendly way to navigate and interact with
// your application.
//
// # Key Components
//
// This package consists of the following key components:
//
//  1. MenuItem: Represents an item within a menu. Each MenuItem has a label, an associated action, and an optional
//     icon. Menu items can be customized to perform specific actions when selected.
//
//  2. Menu: Represents a collection of menu items along with an optional prompt. Menus are used to group related
//     items together and provide a user-friendly interface for selecting actions.
//
//  3. RofiRunner: A utility for running Rofi-based menus with custom options. You can create an instance of RofiRunner
//     and execute Rofi menus with specific settings, such as prompts and menu items.
//
// # Usage Example
//
// Here's a simple example of how to use "rofigo" to create a menu and run it with Rofi:
//
//	// Create a menu
//	myMenu := rofigo.NewMenu().WithPrompt("Select an option:")
//
//	// Add menu items
//	item1 := rofigo.NewMenuItem("Option 1", func() error {
//	    // Define the action for Option 1
//	    // ...
//	    return nil
//	}, "")
//	item2 := rofigo.NewMenuItem("Option 2", func() error {
//	    // Define the action for Option 2
//	    // ...
//	    return nil
//	}, "")
//
//	myMenu.WithItems(item1, item2)
//
//	// Create a RofiRunner and run the menu
//	runner := rofigo.NewRofiRunner(myMenu)
//	runner.Run()
//
// This code creates a menu with two options and executes it using Rofi.
//
// # Package Dependencies
//
// The "rofigo" package relies on external tools like Rofi, so make sure you have Rofi installed on your system
// to use this package.
package rofigo
