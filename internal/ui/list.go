package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
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
	return i.product.Name + " " + i.product.Category
}

type listModel struct {
	list          list.Model
	data          model.Data
	categoryIndex int
	filteredData  model.Data
	searchInput   textinput.Model
	showSearch    bool
}

func newListModel(data model.Data) listModel {
	filtered := filterByCategory(data, data.Categories[0])
	items := createListItems(filtered)

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = ""
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	searchInput := textinput.New()
	searchInput.Placeholder = "Suche..."
	searchInput.CharLimit = 50
	searchInput.Width = 30

	return listModel{
		list:          l,
		data:          data,
		filteredData:  filtered,
		categoryIndex: 0,
		searchInput:   searchInput,
		showSearch:    false,
	}
}

func createListItems(data model.Data) []list.Item {
	items := make([]list.Item, len(data.Products))
	for i, p := range data.Products {
		items[i] = listItem{product: p}
	}
	return items
}

func filterByCategory(data model.Data, category string) model.Data {
	filtered := model.Data{
		Categories: data.Categories,
		Products:   []model.Product{},
	}
	for _, p := range data.Products {
		if p.Category == category {
			filtered.Products = append(filtered.Products, p)
		}
	}
	return filtered
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
				m.filteredData = filterByCategory(m.data, m.data.Categories[m.categoryIndex])
				m.list.SetItems(createListItems(m.filteredData))
				return m, tea.Batch(cmds...)
			case "enter":
				m.showSearch = false
				m.filteredData = filterBySearch(m.filteredData, m.searchInput.Value())
				m.list.SetItems(createListItems(m.filteredData))
				return m, tea.Batch(cmds...)
			}
		}
		return m, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			if len(m.filteredData.Products) > 0 {
				idx := m.list.Index()
				if idx < len(m.filteredData.Products) {
					for i := range m.data.Products {
						if m.data.Products[i].ID == m.filteredData.Products[idx].ID {
							m.data.Products[i].Completed = !m.data.Products[i].Completed
							break
						}
					}
					m.filteredData = filterByCategory(m.data, m.data.Categories[m.categoryIndex])
					m.list.SetItems(createListItems(m.filteredData))
				}
			}
			return m, nil
		case "j", "down":
			m.list.CursorDown()
			return m, nil
		case "k", "up":
			m.list.CursorUp()
			return m, nil
		case "right":
			m.categoryIndex++
			if m.categoryIndex >= len(m.data.Categories) {
				m.categoryIndex = 0
			}
			m.filteredData = filterByCategory(m.data, m.data.Categories[m.categoryIndex])
			m.list.SetItems(createListItems(m.filteredData))
			return m, nil
		case "left":
			m.categoryIndex--
			if m.categoryIndex < 0 {
				m.categoryIndex = len(m.data.Categories) - 1
			}
			m.filteredData = filterByCategory(m.data, m.data.Categories[m.categoryIndex])
			m.list.SetItems(createListItems(m.filteredData))
			return m, nil
		case "/":
			m.showSearch = true
			m.searchInput.Focus()
			return m, textinput.Blink
		case "1", "2", "3", "4", "5":
			idx := int(msg.String()[0] - '1')
			if idx < len(m.data.Categories) {
				m.categoryIndex = idx
				m.filteredData = filterByCategory(m.data, m.data.Categories[m.categoryIndex])
				m.list.SetItems(createListItems(m.filteredData))
			}
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

	b.WriteString(m.categoryTabs())
	b.WriteString("\n")

	if m.showSearch {
		b.WriteString(m.searchInput.View() + "\n\n")
	}

	if len(m.filteredData.Products) == 0 {
		if len(m.data.Products) == 0 {
			b.WriteString("Keine Produkte. Drücke 'n' um ein neues Produkt zu erstellen.\n\n")
		} else {
			b.WriteString("Keine Produkte in dieser Kategorie.\n\n")
		}
	}

	return b.String() + m.list.View()
}

func (m listModel) categoryTabs() string {
	var tabs []string
	for i, cat := range m.data.Categories {
		count := 0
		for _, p := range m.data.Products {
			if p.Category == cat {
				count++
			}
		}

		if i == m.categoryIndex {
			tabs = append(tabs, CategoryStyle.Render(fmt.Sprintf(" [%d] %s (%d) ", i+1, cat, count)))
		} else {
			tabs = append(tabs, fmt.Sprintf(" [%d] %s (%d) ", i+1, cat, count))
		}
	}
	return strings.Join(tabs, " ")
}

func (m listModel) SetSize(width, height int) {
	m.list.SetSize(width, height-12)
	m.searchInput.Width = width - 4
}

func (m listModel) GetData() model.Data {
	return m.data
}

func (m listModel) GetCurrentCategory() string {
	if len(m.data.Categories) > 0 {
		return m.data.Categories[m.categoryIndex]
	}
	return ""
}
