package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jens/van-planner/internal/model"
)

type FormState int

const (
	StateName FormState = iota
	StateCategory
	StateEstimatedCost
	StateActualCost
	StateShopLink
	StateNotes
)

type formModel struct {
	inputs      []textinput.Model
	state       FormState
	product     model.Product
	categoryIdx int
	submitting  bool
	err         error
}

func newFormModel(data model.Data, editProduct *model.Product, category string) formModel {
	inputs := make([]textinput.Model, 6)

	inputs[0] = textinput.New()
	inputs[0].Placeholder = "Produktname"
	inputs[0].Focus()
	inputs[0].CharLimit = 100
	inputs[0].Width = 40

	inputs[1] = textinput.New()
	inputs[1].Placeholder = "Kategorie"
	inputs[1].CharLimit = 50
	inputs[1].Width = 40

	inputs[2] = textinput.New()
	inputs[2].Placeholder = "Kosten geschätzt (€)"
	inputs[2].CharLimit = 15
	inputs[2].Width = 25

	inputs[3] = textinput.New()
	inputs[3].Placeholder = "Kosten tatsächlich (€)"
	inputs[3].CharLimit = 15
	inputs[3].Width = 25

	inputs[4] = textinput.New()
	inputs[4].Placeholder = "Shop-Link"
	inputs[4].CharLimit = 200
	inputs[4].Width = 40

	inputs[5] = textinput.New()
	inputs[5].Placeholder = "Notizen"
	inputs[5].CharLimit = 500
	inputs[5].Width = 40

	if editProduct != nil {
		inputs[0].SetValue(editProduct.Name)
		inputs[1].SetValue(editProduct.Category)
		if editProduct.EstimatedCost > 0 {
			inputs[2].SetValue(fmt.Sprintf("%.2f", editProduct.EstimatedCost))
		}
		if editProduct.ActualCost > 0 {
			inputs[3].SetValue(fmt.Sprintf("%.2f", editProduct.ActualCost))
		}
		inputs[4].SetValue(editProduct.ShopLink)
		inputs[5].SetValue(editProduct.Notes)
	} else if category != "" {
		inputs[1].SetValue(category)
	} else if len(data.Categories) > 0 {
		inputs[1].SetValue(data.Categories[0])
	}

	return formModel{
		inputs:      inputs,
		state:       StateName,
		product:     model.NewProduct("", ""),
		categoryIdx: 0,
	}
}

func (m formModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m formModel) Update(msg tea.Msg) (formModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.submitting = true
			return m, tea.Quit
		case "tab", "shift+tab", "enter":
			if msg.String() == "enter" && m.state == StateNotes {
				m.submitting = true
				return m, tea.Quit
			}

			s := int(m.state)

			if msg.String() == "shift+tab" {
				s--
				if s < 0 {
					s = len(m.inputs) - 1
				}
			} else {
				s++
				if s >= len(m.inputs) {
					s = 0
				}
			}

			m.state = FormState(s)

			for i := range m.inputs {
				m.inputs[i].Blur()
			}

			m.inputs[m.state].Focus()

			return m, textinput.Blink
		}
	}

	for i := range m.inputs {
		var cmd tea.Cmd
		m.inputs[i], cmd = m.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m formModel) View() string {
	var b strings.Builder

	b.WriteString(TitleStyle.Render("Produkt ") + "\n")
	b.WriteString("\n")

	labels := []string{"Name:", "Kategorie:", "Kosten geschätzt:", "Kosten tatsächlich:", "Shop-Link:", "Notizen:"}

	for i := range m.inputs {
		b.WriteString(labels[i] + "\n")
		b.WriteString(m.inputs[i].View() + "\n\n")
	}

	b.WriteString(HelpStyle.Render("Tab/Enter: Weiter | Shift+Tab: Zurück | Enter bei Notizen: Speichern | Esc: Abbrechen"))
	b.WriteString("\n")

	return b.String()
}

func (m formModel) GetProduct(data model.Data) (model.Product, error) {
	name := strings.TrimSpace(m.inputs[0].Value())
	if name == "" {
		return model.Product{}, fmt.Errorf("Name ist erforderlich")
	}

	category := strings.TrimSpace(m.inputs[1].Value())
	if category == "" {
		category = data.Categories[0]
	}

	var estimatedCost float64
	estimatedStr := strings.TrimSpace(m.inputs[2].Value())
	if estimatedStr != "" {
		var err error
		estimatedCost, err = strconv.ParseFloat(estimatedStr, 64)
		if err != nil {
			return model.Product{}, fmt.Errorf("Ungültige geschätzte Kosten: %v", err)
		}
	}

	var actualCost float64
	actualStr := strings.TrimSpace(m.inputs[3].Value())
	if actualStr != "" {
		var err error
		actualCost, err = strconv.ParseFloat(actualStr, 64)
		if err != nil {
			return model.Product{}, fmt.Errorf("Ungültige tatsächliche Kosten: %v", err)
		}
	}

	product := model.NewProduct(name, category)
	product.EstimatedCost = estimatedCost
	product.ActualCost = actualCost
	product.ShopLink = strings.TrimSpace(m.inputs[4].Value())
	product.Notes = strings.TrimSpace(m.inputs[5].Value())

	return product, nil
}
