package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/jens/van-planner/internal/model"
)

type ProductEditView struct {
	product       model.Product
	categories    []string
	categoryIndex int
	formInputs    []textinput.Model
	focusedIndex  int
	width         int
	height        int
	cursorYOffset int
	notesInput    textinput.Model
	showCategory  bool
}

func NewProductEditView(product model.Product, categories []string) *ProductEditView {
	inputs := make([]textinput.Model, 6)

	for i := range inputs {
		inputs[i] = textinput.New()
		inputs[i].PromptStyle = lipgloss.NewStyle().Foreground(sunOrange)
		inputs[i].TextStyle = lipgloss.NewStyle().Foreground(warmWhite)
	}

	inputs[0].Placeholder = "Produktname"
	inputs[0].SetValue(product.Name)
	inputs[0].Width = 50

	inputs[1].SetValue(fmt.Sprintf("%.2f", product.EstimatedCost))
	inputs[1].Width = 15

	inputs[2].SetValue(fmt.Sprintf("%.2f", product.ActualCost))
	inputs[2].Width = 15

	inputs[3].SetValue(fmt.Sprintf("%.2f", product.ShippingCost))
	inputs[3].Width = 15

	countStr := ""
	if product.Count != nil {
		countStr = strconv.Itoa(*product.Count)
	}
	inputs[4].SetValue(countStr)
	inputs[4].Width = 10

	inputs[5].SetValue(product.ShopLink)
	inputs[5].Width = 50

	notesInput := textinput.New()
	notesInput.PromptStyle = lipgloss.NewStyle().Foreground(sunOrange)
	notesInput.TextStyle = lipgloss.NewStyle().Foreground(warmWhite)
	notesInput.SetValue(product.Notes)
	notesInput.Width = 60

	categoryIndex := 0
	for i, cat := range categories {
		if cat == product.Category {
			categoryIndex = i
			break
		}
	}

	return &ProductEditView{
		product:       product,
		categories:    categories,
		categoryIndex: categoryIndex,
		formInputs:    inputs,
		focusedIndex:  0,
		notesInput:    notesInput,
		showCategory:  false,
	}
}

func (p *ProductEditView) Init() tea.Cmd {
	return textinput.Blink
}

func (p *ProductEditView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if p.showCategory {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "up":
				if p.categoryIndex > 0 {
					p.categoryIndex--
				}
			case "down":
				if p.categoryIndex < len(p.categories)-1 {
					p.categoryIndex++
				}
			case "enter", "tab":
				p.showCategory = false
				p.focusedIndex = 2
				return p, p.updateFocus()
			case "esc":
				p.showCategory = false
			}
		}
		return p, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			p.focusedIndex = (p.focusedIndex + 1) % (len(p.formInputs) + 1)
			return p, p.updateFocus()
		case "shift+tab":
			p.focusedIndex = (p.focusedIndex - 1 + (len(p.formInputs) + 1)) % (len(p.formInputs) + 1)
			return p, p.updateFocus()
		case "up":
			if p.focusedIndex > 0 {
				p.focusedIndex--
				return p, p.updateFocus()
			}
		case "down":
			if p.focusedIndex < len(p.formInputs) {
				p.focusedIndex++
				return p, p.updateFocus()
			}
		case "enter":
			if p.focusedIndex == 1 {
				p.showCategory = true
				return p, nil
			}
		case "ctrl+d":
			return p, func() tea.Msg {
				return DeleteProductMsg{}
			}
		case "ctrl+s":
			return p, p.saveProduct()
		case "esc":
			return p, func() tea.Msg {
				return CancelEditMsg{}
			}
		}
	}

	var cmd tea.Cmd
	if p.focusedIndex < len(p.formInputs) {
		p.formInputs[p.focusedIndex], cmd = p.formInputs[p.focusedIndex].Update(msg)
		cmds = append(cmds, cmd)
	}

	p.notesInput, cmd = p.notesInput.Update(msg)
	cmds = append(cmds, cmd)

	return p, tea.Batch(cmds...)
}

func (p *ProductEditView) View() string {
	var sb strings.Builder

	sb.WriteString("\n")
	sb.WriteString(editTitleStyle.Render("Produkt bearbeiten"))
	sb.WriteString("\n\n")

	sb.WriteString(fmt.Sprintf("%-*s", 20, "Name:"))
	if p.focusedIndex == 0 {
		sb.WriteString(focusedLabelStyle.Render(p.formInputs[0].View()))
	} else {
		sb.WriteString(p.formInputs[0].View())
	}
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf("%-*s", 20, "Kategorie:"))
	categoryName := p.categories[p.categoryIndex]
	if p.showCategory {
		sb.WriteString(categorySelectStyle.Render("▼ " + categoryName + " [▲/▼ auswählen, Enter bestätigen]"))
	} else if p.focusedIndex == 1 {
		sb.WriteString(focusedLabelStyle.Render("▶ " + categoryName + " [Enter für Auswahl]"))
	} else {
		sb.WriteString(categoryName)
	}
	sb.WriteString("\n")

	if p.showCategory {
		sb.WriteString("\n")
		for i, cat := range p.categories {
			if i == p.categoryIndex {
				sb.WriteString(categorySelectedStyle.Render("  > " + cat))
			} else {
				sb.WriteString("    " + cat)
			}
			sb.WriteString("\n")
		}
	}

	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%-*s", 20, "Geschätzte Kosten:"))
	if p.focusedIndex == 2 {
		sb.WriteString(focusedLabelStyle.Render(p.formInputs[1].View()))
	} else {
		sb.WriteString(p.formInputs[1].View())
	}
	sb.WriteString(" €\n")

	sb.WriteString(fmt.Sprintf("%-*s", 20, "Tatsächliche Kosten:"))
	if p.focusedIndex == 3 {
		sb.WriteString(focusedLabelStyle.Render(p.formInputs[2].View()))
	} else {
		sb.WriteString(p.formInputs[2].View())
	}
	sb.WriteString(" €\n")

	sb.WriteString(fmt.Sprintf("%-*s", 20, "Versandkosten:"))
	if p.focusedIndex == 4 {
		sb.WriteString(focusedLabelStyle.Render(p.formInputs[3].View()))
	} else {
		sb.WriteString(p.formInputs[3].View())
	}
	sb.WriteString(" €\n")

	sb.WriteString(fmt.Sprintf("%-*s", 20, "Anzahl:"))
	if p.focusedIndex == 5 {
		sb.WriteString(focusedLabelStyle.Render(p.formInputs[4].View()))
	} else {
		sb.WriteString(p.formInputs[4].View())
	}
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf("%-*s", 20, "Shop-Link:"))
	if p.focusedIndex == 6 {
		sb.WriteString(focusedLabelStyle.Render(p.formInputs[5].View()))
	} else {
		sb.WriteString(p.formInputs[5].View())
	}
	sb.WriteString("\n")

	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%-*s", 20, "Notizen:"))
	if p.focusedIndex == 7 {
		sb.WriteString(focusedLabelStyle.Render(p.notesInput.View()))
	} else {
		sb.WriteString(p.notesInput.View())
	}
	sb.WriteString("\n\n")

	sb.WriteString(shortcutStyle.Render("[CTRL-S] Speichern"))
	sb.WriteString("  ")
	sb.WriteString(deleteStyle.Render("[CTRL-D] Löschen"))
	sb.WriteString("  ")
	sb.WriteString(shortcutStyle.Render("[Esc] Abbrechen"))

	return sb.String()
}

func (p *ProductEditView) updateFocus() tea.Cmd {
	var cmds []tea.Cmd
	for i := range p.formInputs {
		if i == p.focusedIndex {
			cmds = append(cmds, p.formInputs[i].Focus())
		} else {
			p.formInputs[i].Blur()
		}
	}

	if p.focusedIndex == len(p.formInputs)+1 {
		cmds = append(cmds, p.notesInput.Focus())
	} else {
		p.notesInput.Blur()
	}

	return tea.Batch(cmds...)
}

func (p *ProductEditView) saveProduct() tea.Cmd {
	return func() tea.Msg {
		product := p.product
		product.Name = p.formInputs[0].Value()
		product.Category = p.categories[p.categoryIndex]

		if estimatedCost, err := strconv.ParseFloat(p.formInputs[1].Value(), 64); err == nil {
			product.EstimatedCost = estimatedCost
		}

		if actualCost, err := strconv.ParseFloat(p.formInputs[2].Value(), 64); err == nil {
			product.ActualCost = actualCost
		}

		if shippingCost, err := strconv.ParseFloat(p.formInputs[3].Value(), 64); err == nil {
			product.ShippingCost = shippingCost
		}

		if countStr := p.formInputs[4].Value(); countStr != "" {
			if count, err := strconv.Atoi(countStr); err == nil {
				product.Count = &count
			}
		} else {
			product.Count = nil
		}

		product.ShopLink = p.formInputs[5].Value()
		product.Notes = p.notesInput.Value()

		return SaveProductMsg{Product: product}
	}
}

func (p *ProductEditView) GetProduct() model.Product {
	return p.product
}

func (p *ProductEditView) GetShortcuts() string {
	return "[Tab] Nächstes Feld [↑]↑ [↓]↓ [Enter] Kategorie wählen [CTRL-S] Speichern [CTRL-D] Löschen [Esc] Abbrechen"
}

type SaveProductMsg struct {
	Product model.Product
}

type CancelEditMsg struct{}

type DeleteProductMsg struct{}
