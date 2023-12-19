package cli

import (
	"bytes"
	"fmt"
	"github.com/aperturedev/aperture/internal/app"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run aperture app",
	// Short: "Initialize new aperture app",
	Run: func(cmd *cobra.Command, args []string) {
		// Determine language
		// Run specific language app parser
		// Create aperture.yaml
		// Run specific projection parser
		// (this also happens on each change / each build
		// - build the project
		// - analyze it
		// - output yaml config
		// Generate docker compose in .aperture from base and projections

		a, err := app.Load()
		if err != nil {
			// TODO - Check if aperture yaml exists
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
		}

		fmt.Println()

		fmt.Println(fmt.Sprintf("> Running %s", a.ID))

		fmt.Println("> Detected .NET Solution")

		// TODO - if no netConfig section present then show this prompt, otherwise skip
		if a.NetConfig != nil {
			fmt.Println(fmt.Sprintf("> Using %s project", a.NetConfig.ProjectName))
		} else {
			err = chooseProject(a)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, err.Error())
			}
		}

		// build
		// (this also happens on each change / each build
		// - build the project
		// - analyze it
		// - output yaml config

		fmt.Println(fmt.Sprintf("> Building %s", a.NetConfig.ProjectName))

		err = build(a)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
		}

		fmt.Println(fmt.Sprintf("> Analyzing %s", a.NetConfig.ProjectName))

		err = analyze(a)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
		}

		fmt.Println("> Done!")

		fmt.Println()

		fmt.Println("> Aperture running on http://aperture.localhost")
		fmt.Println("  Watching for changes...")

		var ln string
		fmt.Scanln(&ln)
	},
}

func chooseProject(a *app.App) error {
	fmt.Println()

	m := initialModel()

	lines, err := listProjects()
	if err != nil {
		return err
	}

	m.choices = lines

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	fmt.Println()

	// TODO - Write to netConfig section
	a.SetNetConfig(app.NetConfig{
		ProjectName: m.selected,
	})

	return a.Save()
}

func listProjects() ([]string, error) {
	cmd := exec.Command("dotnet", "sln", "list")

	buf := bytes.NewBufferString("")

	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(buf.String(), "\n")

	var projects []string

	for _, l := range lines[2 : len(lines)-1] {
		s := strings.Split(l, "/")
		projects = append(projects, s[0])
	}

	return projects, nil
}

func build(a *app.App) error {
	cmd := exec.Command("dotnet", "build", a.NetConfig.ProjectName, "-o", ".aperture/build")

	// cmd.Stdout = os.Stderr

	return cmd.Run()
}

func analyze(a *app.App) error {
	cmd := exec.Command("/Users/anes/projects/Aperture/Aperture.Cli/bin/Debug/net8.0/Aperture.Cli")

	buf := bytes.NewBuffer(nil)

	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf(".aperture/%s.cfg", a.NetConfig.ProjectName), buf.Bytes(), 0755)
}

type model struct {
	choices  []string // items on the to-do list
	cursor   int      // which to-do list item our cursor is pointing at
	selected string
}

func initialModel() *model {
	return &model{
		// Our to-do list is a grocery list
		choices: []string{"Aperture.Core", "Aperture.Tests", "Aperture.Projections"},
	}
}

func (m *model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			m.selected = m.choices[m.cursor]

			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m *model) View() string {
	// The header
	s := "> Which project contains your projections?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// The footer
	s += "\nPress RETURN to select\n"

	// Send the UI for rendering
	return s
}
