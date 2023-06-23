package render

import (
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"io/ioutil"
	"os"
	"time"
)

// Standard terminal size
const (
	rows    = 24
	columns = 80
)

type tickMsg time.Time

type model struct {
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}

	case tickMsg:
		return m, tea.Quit

	default:
		return m, nil
	}
}

func renderFile(file string) (*model, error) {
	vp := viewport.New(columns, rows)

	in, _ := ioutil.ReadFile(file)
	data := string(in)
	out, _ := glamour.Render(data, "dark")
	vp.SetContent(out)

	return &model{
		viewport: vp,
	}, nil
}

func Render(file string) bool {
	m, _ := renderFile(file)

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
	return true
}

func (m model) View() string {
	return m.viewport.View()
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
