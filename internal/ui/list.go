package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jens/van-planner/internal/model"
)

type listItem struct {
	product model.Product
}

func (i listItem) Title() string {
	completed := " "
	if i.product.Completed {
		completed = "✓"
	}
	return fmt.Sprintf("[%s] %s", completed, i.product.Name)
}

func (i listItem) Description() string {
	if i.product.Price > 0 {
		return fmt.Sprintf("%.2f €", i.product.Price)
	}
	return ""
}

func (i listItem) FilterValue() string {
	return i.product.Name
}

type listModel struct {
	list          list.Model
	data          model.Data
	categoryIndex int
}

func newListModel(data model.Data) listModel {
	items := make([]list.Item, len(data.Products))
	for i, p := range data.Products {
		items[i] = listItem{product: p}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Van Planner"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return listModel{
		list: l,
		data: data,
	}
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			if len(m.data.Products) > 0 {
				idx := m.list.Index()
				m.data.Products[idx].Completed = !m.data.Products[idx].Completed
			}
			return m, nil
		case "j", "down":
			m.list.CursorDown()
			return m, nil
		case "k", "up":
			m.list.CursorUp()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m listModel) View() string {
	var b strings.Builder

	b.WriteString(TitleStyle.Render("Van Planner TUI") + "\n")
	b.WriteString(SubtitleStyle.Render("Planer für deinen Campervan-Ausbau") + "\n\n")

	completed := 0
	for _, p := range m.data.Products {
		if p.Completed {
			completed++
		}
	}

	b.WriteString(fmt.Sprintf("Fortschritt: %d/%d erledigt\n", completed, len(m.data.Products)))
	b.WriteString("\n")
	b.WriteString(HelpStyle.Render("j/k: Navigation | Space: Toggle | q: Quit"))
	b.WriteString("\n")

	return b.String() + m.list.View()
}

func (m listModel) SetSize(width, height int) {
	m.list.SetSize(width, height-10)
}

func (m listModel) GetData() model.Data {
	products := make([]model.Product, len(m.data.Products))
	for i, item := range m.list.Items() {
		if li, ok := item.(listItem); ok {
			products[i] = li.product
		}
	}
	m.data.Products = products
	return m.data
}
