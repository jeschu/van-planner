package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jens/van-planner/internal/model"
)

type listModel struct {
	list          list.Model
	data          model.Data
	categoryIndex int
	filteredData  model.Data
	searchInput   textinput.Model
	showSearch    bool
	categoryWidth int
}

func newListModel(data model.Data) listModel {
	items := createGroupedListItems(data)

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = ""
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)

	searchInput := textinput.New()
	searchInput.Placeholder = "Suche..."
	searchInput.CharLimit = 50
	searchInput.Width = 30

	return listModel{
		list:         l,
		data:         data,
		filteredData: data,
		searchInput:  searchInput,
		showSearch:   false,
	}
}

type groupedListItem struct {
	category string
	product  model.Product
	isHeader bool
}

func (i groupedListItem) Title() string {
	if i.isHeader {
		return i.category
	}
	completed := " "
	if i.product.Completed {
		completed = "✓"
	}
	return fmt.Sprintf("  [%s] %s", completed, i.product.Name)
}

func (i groupedListItem) Description() string {
	if i.isHeader {
		return ""
	}
	if i.product.Price > 0 {
		return fmt.Sprintf("%.2f €", i.product.Price)
	}
	return ""
}

func (i groupedListItem) FilterValue() string {
	if i.isHeader {
		return ""
	}
	return i.product.Name + " " + i.product.Category
}

func createGroupedListItems(data model.Data) []list.Item {
	var items []list.Item
	for _, cat := range data.Categories {
		items = append(items, groupedListItem{category: cat, isHeader: true})
		for _, p := range data.Products {
			if p.Category == cat {
				items = append(items, groupedListItem{category: cat, product: p, isHeader: false})
			}
		}
	}
	return items
}

func filterBySearch(data model.Data, query string) model.Data {
	if query == "" {
		return data
	}
	filtered := model.Data{
		Categories: data.Categories,
		Products:   []model.Product{},
	}
	query = strings.ToLower(query)
	for _, p := range data.Products {
		if strings.Contains(strings.ToLower(p.Name), query) ||
			strings.Contains(strings.ToLower(p.Notes), query) ||
			strings.Contains(strings.ToLower(p.Category), query) {
			filtered.Products = append(filtered.Products, p)
		}
	}
	return filtered
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {
	var cmds []tea.Cmd

	if m.showSearch {
		var cmd tea.Cmd
		m.searchInput, cmd = m.searchInput.Update(msg)
		cmds = append(cmds, cmd)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				m.showSearch = false
				m.searchInput.SetValue("")
				m.filteredData = m.data
				m.list.SetItems(createGroupedListItems(m.filteredData))
				return m, tea.Batch(cmds...)
			case "enter":
				m.showSearch = false
				m.filteredData = filterBySearch(m.data, m.searchInput.Value())
				m.list.SetItems(createGroupedListItems(m.filteredData))
				return m, tea.Batch(cmds...)
			}
		}
		return m, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			idx := m.list.Index()
			items := m.list.Items()
			if idx < len(items) {
				if item, ok := items[idx].(groupedListItem); ok && !item.isHeader {
					for i := range m.data.Products {
						if m.data.Products[i].ID == item.product.ID {
							m.data.Products[i].Completed = !m.data.Products[i].Completed
							break
						}
					}
					m.filteredData = m.data
					m.list.SetItems(createGroupedListItems(m.filteredData))
				}
			}
			return m, nil
		case "/":
			m.showSearch = true
			m.searchInput.Focus()
			return m, textinput.Blink
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

	if m.showSearch {
		b.WriteString(m.searchInput.View() + "\n\n")
	}

	if len(m.data.Products) == 0 {
		b.WriteString("Keine Produkte. Drücke 'n' um ein neues Produkt zu erstellen.\n\n")
	}

	return b.String() + m.list.View()
}

func (m listModel) SetSize(width, height int) {
	m.list.SetSize(width, height-12)
	m.searchInput.Width = width - 4
}

func (m listModel) GetData() model.Data {
	return m.data
}

func (m listModel) GetCurrentCategory() string {
	idx := m.list.Index()
	items := m.list.Items()
	if idx < len(items) {
		if item, ok := items[idx].(groupedListItem); ok && !item.isHeader {
			return item.product.Category
		}
	}
	if len(m.data.Categories) > 0 {
		return m.data.Categories[0]
	}
	return ""
}
