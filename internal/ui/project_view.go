package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

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
			if p.cursorIndex < len(p.project.Products)-1 {
				p.cursorIndex++
			}
		case "k", "up":
			if p.cursorIndex > 0 {
				p.cursorIndex--
			}
		case " ":
			if len(p.project.Products) > 0 {
				p.project.Products[p.cursorIndex].Completed = !p.project.Products[p.cursorIndex].Completed
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

	return p, nil
}

func (p *ProjectView) View() string {
	return p.viewport.View()
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
		sb.WriteString("\n")
	}

	if len(p.project.Products) > 0 {
		sb.WriteString(p.renderDetailSection())
	}

	return sb.String()
}

func (p *ProjectView) renderCategory(category string, products []model.Product) string {
	var sb strings.Builder

	categoryStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("203"))

	sb.WriteString(categoryStyle.Render(category))
	sb.WriteString("\n")

	for i, product := range products {
		cursor := "  "
		if i+p.getProductStartIndex(category) == p.cursorIndex {
			cursor = "> "
		}

		checkbox := "☐"
		if product.Completed {
			checkbox = "☑"
		}

		line := fmt.Sprintf("%s[%s] %d. %s", cursor, checkbox, product.Index, product.Name)
		sb.WriteString(line + "\n")
	}

	return sb.String()
}

func (p *ProjectView) getProductStartIndex(category string) int {
	index := 0
	for _, cat := range p.project.Categories {
		if cat == category {
			return index
		}
		for _, product := range p.project.Products {
			if product.Category == cat {
				index++
			}
		}
	}
	return index
}

func (p *ProjectView) renderDetailSection() string {
	if p.cursorIndex < 0 || p.cursorIndex >= len(p.project.Products) {
		return ""
	}

	product := p.project.Products[p.cursorIndex]

	divider := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Render(strings.Repeat("─", p.width-4))

	detailStyle := lipgloss.NewStyle().Padding(0, 1)

	details := fmt.Sprintf(`%s
Ausgewähltes Produkt:
Name: %s
Kategorie: %s
Geschätzte Kosten: %.2f €
Tatsächliche Kosten: %.2f €
Versandkosten: %.2f €
Shop-Link: %s
Notizen: %s
Status: %s`,
		divider,
		product.Name,
		product.Category,
		product.EstimatedCost,
		product.ActualCost,
		product.ShippingCost,
		product.ShopLink,
		product.Notes,
		p.statusText(product.Completed),
	)

	return detailStyle.Render(details)
}

func (p *ProjectView) statusText(completed bool) string {
	if completed {
		return "☑ erledigt"
	}
	return "☐ offen"
}

func (p *ProjectView) GetShortcuts() string {
	return "[↑]↑ [↓]↓ [←][→]Scroll [Space]Erledigt [q]Quit"
}

func (p *ProjectView) SetProject(project *model.Project) {
	p.project = project
}

func (p *ProjectView) GetProject() *model.Project {
	return p.project
}
