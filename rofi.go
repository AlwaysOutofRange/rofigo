package rofigo

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// RofiRunner is a utility for running Rofi-based menus with custom options.
type RofiRunner struct {
	Menu *Menu
}

// NewRofiRunner creates a new RofiRunner instance with the specified menu.
func NewRofiRunner(menu *Menu) *RofiRunner {
	return &RofiRunner{
		Menu: menu,
	}
}

// Run executes the Rofi menu using the specified options and runs the associated action when an item is selected.
func (r *RofiRunner) Run() {
	var outputBuffer bytes.Buffer

	path, err := exec.LookPath("rofi")
	if err != nil {
		log.Fatalf("[ERROR] An error ocurred: %s", err)
	}

	cmd := exec.Command(
		path,
		"-dmenu",
		"-i",
		"-p", r.Menu.Prompt,
	)
	cmd.Stdout = &outputBuffer

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("[ERROR] An error ocurred during StdinPipe creation: %s", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("[ERROR] An error ocurred during command starting: %s", err)
	}

	for _, item := range r.Menu.Items {
		if item.Icon != "" {
			fmt.Fprintf(stdin, "%s\000icon\x1f%s\n", item.Label, item.Icon)
			continue
		}

		fmt.Fprintf(stdin, "%s\n", item.Label)
	}
	stdin.Close()

	if err := cmd.Wait(); err != nil {
		log.Fatalf("[ERROR] An error occured during command execution: %s", err)
	}

	selected := outputBuffer.String()

	for _, item := range r.Menu.Items {
		// Trim leading and trailing spaces to ensure an exact match.
		if strings.TrimSpace(selected) == strings.TrimSpace(item.Label) {
			err := item.Action()

			if err != nil {
				log.Fatalf("[ERROR] An error ocurred during Action execution: %s", err)
			}

			break
		}
	}
}
