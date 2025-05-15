package gui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/symonk/dnsui/internal/dns"
)

type Model struct {
	checker *dns.DnsChecker
}

func New(checker *dns.DnsChecker) *Model {
	return &Model{checker: checker}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) View() string {
	return m.checker.GetIpOfHost("google.com")
}
