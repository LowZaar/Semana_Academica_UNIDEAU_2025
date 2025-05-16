package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	items  []string
	selecionados map[int]struct{}
}

func initialModel() model {
	return model{
		cursor: 0,
		items: []string{"Comprar Café", "Comprar Miojo", "Não sair sem pagar"},
		selecionados: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Lista de compras")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selecionados[m.cursor]
			if ok {
				delete(m.selecionados, m.cursor)
			} else {
				m.selecionados[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "O que vamos fazer no mercado?\n\n"

	for index, escolha := range m.items {
		cursor := " "
		if m.cursor == index {
			cursor = ">"
		}

		checked := " "
		
		if _, ok := m.selecionados[index]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, escolha)
	}

	s += "\nAperte Q para sair.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro ao inicializar: %v\n", err)
		os.Exit(1)
	}
}