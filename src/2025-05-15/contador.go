package main

import (
	"fmt"
	"os"
	
	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	contagem int
	fechando bool
}

func modeloInicial() model {
	return model{
		contagem: 0,
		fechando: false,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Contador - Bubble Tea")
}


type(
	increment struct{}
)


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "q", "ctrl+c":
				m.fechando = true
				return m, tea.Quit
			case " ":
				return m, func() tea.Msg { return increment{} }
			}
		case increment:
			m.contagem++
	}
	return m, nil
}


func (m model) View() string {
	if m.fechando {
		return fmt.Sprintf("Contagem Final: %d!\n", m.contagem)
	}
	
	s := fmt.Sprintf("\n 🔢 Contador: %d\n\n", m.contagem)
	s += "Espaço: Aumentar valor\n"
	s += "Q: Fechar\n\n"
	
	return s
}


func main() {
	// Crie e rode o programa
	p := tea.NewProgram(modeloInicial())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro ao inicializar: %v\n", err)
		os.Exit(1)
	}
}