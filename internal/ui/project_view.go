package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/jens/van-planner/internal/model"
)

type ProjectView struct {
	project          *model.Project
	cursorIndex      int
	viewport         viewport.Model
	width            int
	height           int
	horizontalOffset int
}

func NewProjectView(project *model.Project) *ProjectView {
	vp := viewport.New(80, 20)
	return &ProjectView{
		project:  project,
		viewport: vp,
	}
}

func (p *ProjectView) Init() tea.Cmd {
	return nil
}

func (p *ProjectView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if len(p.project.Products) > 0 {
				p.cursorIndex = (p.cursorIndex + 1) % len(p.project.Products)
			}
		case "k", "up":
			if len(p.project.Products) > 0 {
				p.cursorIndex = (p.cursorIndex - 1 + len(p.project.Products)) % len(p.project.Products)
			}
		case " ":
			if len(p.project.Products) > 0 {
				p.toggleProductCompleted(p.cursorIndex)
			}
		case "left":
			if p.horizontalOffset > 0 {
				p.horizontalOffset--
			}
		case "right":
			p.horizontalOffset++
		}

	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
		p.viewport.Width = msg.Width
		p.viewport.Height = msg.Height - 10
	}

	p.viewport.SetContent(p.renderContent())
	p.scrollViewport()

	return p, nil
}

func (p *ProjectView) View() string {
	return "\n" + p.viewport.View()
}

func (p *ProjectView) renderContent() string {
	var sb strings.Builder

	categoryProducts := make(map[string][]model.Product)
	for _, product := range p.project.Products {
		categoryProducts[product.Category] = append(categoryProducts[product.Category], product)
	}

	for _, category := range p.project.Categories {
		products, exists := categoryProducts[category]
		if !exists || len(products) == 0 {
			continue
		}

		sb.WriteString(p.renderCategory(category, products))
	}

	sb.WriteString(p.renderTotalSum())

	return sb.String()
}

func (p *ProjectView) renderCategory(category string, products []model.Product) string {
	var sb strings.Builder

	sb.WriteString(categoryStyle.Render(category))
	sb.WriteString("\n")

	sb.WriteString(p.renderTableHeader())

	categorySumEstimated := 0.0
	categorySumActual := 0.0

	for i, product := range products {
		cursor := "  "
		if i+p.getProductStartIndex(category) == p.cursorIndex {
			cursor = cursorStyle.Render("> ")
		}

		estimatedCost := fmt.Sprintf("%.2f €", product.EstimatedCost)
		actualCost := fmt.Sprintf("%.2f €", product.ActualCost)
		link := product.ShopLink
		if len(link) > 40 {
			link = link[:37] + "..."
		}

		line := fmt.Sprintf("%s%-40s %12s %12s %s",
			cursor,
			product.Name,
			estimatedCost,
			actualCost,
			linkStyle.Render(link))
		sb.WriteString(line + "\n")

		categorySumEstimated += product.EstimatedCost
		categorySumActual += product.ActualCost
	}

	sb.WriteString(p.renderSumRow(categorySumEstimated, categorySumActual))
	sb.WriteString("\n")

	return sb.String()
}

func (p *ProjectView) renderTableHeader() string {
	var sb strings.Builder
	header := fmt.Sprintf("  %-40s %12s %12s %s",
		tableHeaderStyle.Render("Name"),
		tableHeaderStyle.Render("Kosten geschätzt"),
		tableHeaderStyle.Render("Kosten tatsächlich"),
		tableHeaderStyle.Render("Link"))
	sb.WriteString(header + "\n")
	return sb.String()
}

func (p *ProjectView) renderSumRow(estimated, actual float64) string {
	return fmt.Sprintf("  %-40s %12s %12s",
		tableSumStyle.Render("Summe"),
		tableSumStyle.Render(fmt.Sprintf("%.2f €", estimated)),
		tableSumStyle.Render(fmt.Sprintf("%.2f €", actual)))
}

func (p *ProjectView) renderTotalSum() string {
	var sb strings.Builder
	sb.WriteString("\n")

	totalEstimated := 0.0
	totalActual := 0.0
	for _, product := range p.project.Products {
		totalEstimated += product.EstimatedCost
		totalActual += product.ActualCost
	}

	sb.WriteString(totalSumStyle.Render("Gesamtsumme"))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("  %-40s %12s %12s",
		"",
		totalSumStyle.Render(fmt.Sprintf("%.2f €", totalEstimated)),
		totalSumStyle.Render(fmt.Sprintf("%.2f €", totalActual))))
	sb.WriteString("\n")

	return sb.String()
}

func (p *ProjectView) getProductStartIndex(category string) int {
	index := 0
	for _, product := range p.project.Products {
		if product.Category == category {
			return index
		}
		index++
	}
	return index
}

func (p *ProjectView) GetShortcuts() string {
	return "[↑][↓][←][→]Scroll [Space]Erledigt [q]Quit"
}

func (p *ProjectView) scrollViewport() {
	targetLine := p.getCursorLineNumber()

	if targetLine < p.viewport.YOffset {
		if targetLine > 0 {
			p.viewport.YOffset = targetLine - 1
		} else {
			p.viewport.YOffset = 0
		}
	} else if targetLine >= p.viewport.YOffset+p.viewport.Height {
		p.viewport.YOffset = targetLine - p.viewport.Height + 1
	}
}

func (p *ProjectView) getCursorLineNumber() int {
	line := 0
	currentCategory := ""

	for i, product := range p.project.Products {
		if product.Category != currentCategory {
			if currentCategory != "" {
				line++
			}
			line += 2
			currentCategory = product.Category
		}
		if i == p.cursorIndex {
			return line
		}
		line++
	}
	return 0
}

func (p *ProjectView) SetProject(project *model.Project) {
	p.project = project
}

func (p *ProjectView) GetProject() *model.Project {
	return p.project
}

func (p *ProjectView) toggleProductCompleted(index int) {
	if index < 0 || index >= len(p.project.Products) {
		return
	}

	newProducts := make([]model.Product, len(p.project.Products))
	for i, product := range p.project.Products {
		newProducts[i] = product
	}

	newProducts[index].Completed = !newProducts[index].Completed
	p.project.Products = newProducts
}
