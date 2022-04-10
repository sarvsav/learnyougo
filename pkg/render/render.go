package render

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	padding  = 2
	maxWidth = 80
)

type tickMsg time.Time

type model struct {
	progress progress.Model
}

func (_ model) Init() tea.Cmd {
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		cmd := m.progress.IncrPercent(0.25)
		return m, tea.Batch(tickCmd(), cmd)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func Render() bool {
	m := model{
		progress: progress.New(progress.WithDefaultGradient()),
	}

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}

	in, _ := ioutil.ReadFile("./exercises/1/hint.en.md")
	data1 := string(in)
	out, _ := glamour.Render(data1, "dark")
	fmt.Print(out)
	return true
}

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

func (e model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + e.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}
