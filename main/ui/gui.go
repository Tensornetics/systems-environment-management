package ui

import (
	"fmt"
	"log"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

// RunGUI runs the graphical user interface for the "systems-environment-management" program
func RunGUI() {
	// Initialize GTK
	gtk.Init(nil)

	// Create the main window
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Systems Environment Management")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create the main vertical box
	vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	win.Add(vbox)

	// Create the "Start" button
	startButton, err := gtk.ButtonNewWithLabel("Start")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	startButton.Connect("clicked", func() {
		log.Println("Starting system...")
		// Start the system
		// ...
	})
	vbox.PackStart(startButton, false, false, 0)

	// Create the "Stop" button
	stopButton, err := gtk.ButtonNewWithLabel("Stop")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	stopButton.Connect("clicked", func() {
		log.Println("Stopping system...")
		// Stop the system
		// ...
	})
	vbox.PackStart(stopButton, false, false, 0)

	// Create the "Status" button
	statusButton, err := gtk.ButtonNewWithLabel("Status")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	statusButton.Connect("clicked", func() {
		log.Println("Getting system status...")
		// Get the status of the system
		// ...
	})
	vbox.PackStart(statusButton, false, false, 0)

	// Show the window and start the GTK main loop
	win.ShowAll()
	gtk.Main()
}
