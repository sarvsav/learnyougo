package render

import (
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"io/ioutil"
	"os"
)

const (
	padding  = 2
	maxWidth = 80
)

type model struct {
	viewport viewport.Model
}

func (_ model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	default:
		return m, nil
	}
}

func newExample() (*model, error) {
	vp := viewport.New(80, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	in, _ := ioutil.ReadFile("./exercises/1/solution/solution.md")
	data1 := string(in)
	out, _ := glamour.Render(data1, "dark")
	vp.SetContent(out)

	return &model{
		viewport: vp,
	}, nil
}

func Render() bool {
	m, _ := newExample()

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
	return true
}

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

func (e model) View() string {
	return e.viewport.View()
}
