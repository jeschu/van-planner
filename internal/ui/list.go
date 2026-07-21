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
	projectName   string
	contentHeight int
}

func newListModel(data model.Data, projectName string) listModel {
	items := createGroupedListItems(data)

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = ""
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	l.SetFilteringEnabled(false)

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
		projectName:  projectName,
	}
}

type groupedListItem struct {
	category          string
	product           model.Product
	isHeader          bool
	isSum             bool
	categorySum       float64
	categoryActualSum float64
	totalSum          float64
	totalActualSum    float64
}

func (i groupedListItem) Title() string {
	if i.isHeader {
		return i.category
	}
	if i.isSum {
		return "Summe"
	}
	completed := " "
	if i.product.Completed {
		completed = "✓"
	}
	return fmt.Sprintf("[%s] %s", completed, i.product.Name)
}

func (i groupedListItem) Description() string {
	if i.isHeader {
		return ""
	}
	if i.isSum {
		return fmt.Sprintf("%.2f € | %.2f €", i.categorySum, i.categoryActualSum)
	}
	return fmt.Sprintf("%.2f € | %.2f €", i.product.EstimatedCost, i.product.ActualCost)
}

func (i groupedListItem) FilterValue() string {
	if i.isHeader || i.isSum {
		return ""
	}
	return i.product.Name + " " + i.product.Category
}

func createGroupedListItems(data model.Data) []list.Item {
	var items []list.Item
	var grandTotal float64
	var grandTotalActual float64

	for _, cat := range data.Categories {
		items = append(items, groupedListItem{category: cat, isHeader: true})

		catSum := 0.0
		catActualSum := 0.0

		for _, p := range data.Products {
			if p.Category == cat {
				items = append(items, groupedListItem{
					category: cat,
					product:  p,
					isHeader: false,
				})
				catSum += p.EstimatedCost
				catActualSum += p.ActualCost
			}
		}

		items = append(items, groupedListItem{
			category:          cat,
			isHeader:          false,
			isSum:             true,
			categorySum:       catSum,
			categoryActualSum: catActualSum,
		})

		grandTotal += catSum
		grandTotalActual += catActualSum
	}

	items = append(items, groupedListItem{
		isHeader:       false,
		isSum:          true,
		totalSum:       grandTotal,
		totalActualSum: grandTotalActual,
	})

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
				if item, ok := items[idx].(groupedListItem); ok && !item.isHeader && !item.isSum {
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

	b.WriteString(TitleStyle.Render("Van Planner") + SubtitleStyle.Render(" – Planer für deinen Campervan-Ausbau") + "\n")
	b.WriteString(ProjectStyle.Render("Projekt: "+m.projectName) + "\n\n")

	if m.showSearch {
		b.WriteString(m.searchInput.View() + "\n\n")
	}

	if len(m.data.Products) == 0 {
		b.WriteString("Keine Produkte. Drücke 'n' um ein neues Produkt zu erstellen.\n\n")
	}

	b.WriteString(renderTable(m, m.list.Height()))

	return b.String()
}

func renderTable(m listModel, height int) string {
	var b strings.Builder

	items := m.list.Items()
	if len(items) == 0 {
		return ""
	}

	tableWidth := m.list.Width()
	if tableWidth < 60 {
		tableWidth = 60
	}

	linkWidth := tableWidth / 3
	if linkWidth > 40 {
		linkWidth = 40
	}
	if linkWidth < 20 {
		linkWidth = 20
	}

	remainingWidth := tableWidth - linkWidth - 10
	nameWidth := remainingWidth / 2
	estimatedWidth := remainingWidth / 4
	actualWidth := remainingWidth / 4

	if nameWidth < 20 {
		nameWidth = 20
	}
	if estimatedWidth < 10 {
		estimatedWidth = 10
	}
	if actualWidth < 10 {
		actualWidth = 10
	}

	b.WriteString(renderHeader(nameWidth, estimatedWidth, actualWidth, linkWidth))

	selectedIndex := m.list.Index()
	itemIndex := 0
	contentItems := []string{}

	for _, item := range items {
		if gi, ok := item.(groupedListItem); ok {
			if gi.isHeader {
				contentItems = append(contentItems, renderCategoryHeader(gi.category, nameWidth, estimatedWidth, actualWidth, linkWidth))
			} else if gi.isSum && gi.totalSum > 0 {
				contentItems = append(contentItems, renderTotalSum(gi.totalSum, gi.totalActualSum, nameWidth, estimatedWidth, actualWidth, linkWidth))
			} else if gi.isSum {
				contentItems = append(contentItems, renderCategorySum(gi.categorySum, gi.categoryActualSum, nameWidth, estimatedWidth, actualWidth, linkWidth))
			} else {
				if itemIndex == selectedIndex {
					contentItems = append(contentItems, renderProductRowSelected(gi.product, nameWidth, estimatedWidth, actualWidth, linkWidth))
				} else {
					contentItems = append(contentItems, renderProductRow(gi.product, nameWidth, estimatedWidth, actualWidth, linkWidth))
				}
				itemIndex++
			}
		}
	}

	for i := 0; i < len(contentItems); i++ {
		b.WriteString(contentItems[i])
	}

	return b.String()
}

func renderHeader(nameWidth, estimatedWidth, actualWidth, linkWidth int) string {
	header := fmt.Sprintf("%-*s %*s %*s %s\n",
		nameWidth, "Produkt",
		estimatedWidth, "Kosten geschätzt",
		actualWidth, "Kosten tatsächlich",
		"Link")
	return CategoryStyle.Render(header)
}

func renderCategoryHeader(category string, nameWidth, estimatedWidth, actualWidth, linkWidth int) string {
	totalWidth := nameWidth + estimatedWidth + actualWidth + linkWidth + 6
	separator := strings.Repeat("─", totalWidth)
	return "\n" + CategoryStyle.Render(category) + "\n" + separator + "\n"
}

func renderProductRow(p model.Product, nameWidth, estimatedWidth, actualWidth, linkWidth int) string {
	name := p.Name
	if len(name) > nameWidth {
		name = name[:nameWidth-3] + "..."
	}

	estimated := fmt.Sprintf("%*.2f €", estimatedWidth, p.EstimatedCost)
	actual := fmt.Sprintf("%*.2f €", actualWidth, p.ActualCost)

	link := p.ShopLink
	if len(link) > linkWidth {
		link = link[:linkWidth-3] + "..."
	}

	return fmt.Sprintf("  %-*s %s %s %s\n", nameWidth, name, estimated, actual, link)
}

func renderProductRowSelected(p model.Product, nameWidth, estimatedWidth, actualWidth, linkWidth int) string {
	name := p.Name
	if len(name) > nameWidth {
		name = name[:nameWidth-3] + "..."
	}

	estimated := fmt.Sprintf("%*.2f €", estimatedWidth, p.EstimatedCost)
	actual := fmt.Sprintf("%*.2f €", actualWidth, p.ActualCost)

	link := p.ShopLink
	if len(link) > linkWidth {
		link = link[:linkWidth-3] + "..."
	}

	return SelectedItemStyle.Render(fmt.Sprintf("▶ %-*s %s %s %s\n", nameWidth, name, estimated, actual, link))
}

func renderCategorySum(catSum, catActualSum float64, nameWidth, estimatedWidth, actualWidth, linkWidth int) string {
	sumStr := fmt.Sprintf("%-*s %*.2f € %*.2f €\n",
		nameWidth+2, "Summe",
		estimatedWidth, catSum,
		actualWidth, catActualSum)
	return sumStr
}

func renderTotalSum(totalSum, totalActualSum float64, nameWidth, estimatedWidth, actualWidth, linkWidth int) string {
	sumStr := fmt.Sprintf("\n%-*s %*.2f € %*.2f €\n",
		nameWidth+2, "Gesamtsumme",
		estimatedWidth, totalSum,
		actualWidth, totalActualSum)
	return CategoryStyle.Render(sumStr)
}

func (m listModel) SetSize(width, height int) {
	headerHeight := 4
	footerHeight := 4
	contentHeight := height - headerHeight - footerHeight
	if contentHeight < 1 {
		contentHeight = 1
	}
	m.list.SetSize(width, contentHeight)
	m.searchInput.Width = width - 4
}

func (m listModel) GetData() model.Data {
	return m.data
}

func (m listModel) GetCurrentCategory() string {
	idx := m.list.Index()
	items := m.list.Items()
	if idx < len(items) {
		if item, ok := items[idx].(groupedListItem); ok && !item.isHeader && !item.isSum {
			return item.product.Category
		}
	}
	if len(m.data.Categories) > 0 {
		return m.data.Categories[0]
	}
	return ""
}
