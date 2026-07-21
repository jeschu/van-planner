package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/jens/van-planner/internal/model"
)

type AppState int

const (
	StateProjectView AppState = iota
	StateProjectList
	StateHelp
	StateProductEdit
	StateDeleteDialog
)

type App struct {
	width            int
	height           int
	state            AppState
	projectName      string
	projectView      *ProjectView
	projectList      *ProjectListView
	helpView         *HelpView
	productEditView  *ProductEditView
	deleteDialog     *DeleteDialogView
	cursorIndex      int
	viewportOffset   int
	horizontalOffset int
	saveProjectCmd   func() tea.Cmd
}

func NewApp(projectName string, projectView *ProjectView) *App {
	return &App{
		state:            StateProjectView,
		projectName:      projectName,
		projectView:      projectView,
		helpView:         NewHelpView(),
		cursorIndex:      0,
		viewportOffset:   0,
		horizontalOffset: 0,
		saveProjectCmd:   nil,
	}
}

func (a *App) Init() tea.Cmd {
	return nil
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return a, tea.Quit
		case "ctrl+o":
			if a.state == StateProjectView {
				a.cursorIndex = a.projectView.cursorIndex
				a.viewportOffset = a.projectView.viewport.YOffset
				a.horizontalOffset = a.projectView.horizontalOffset
				a.state = StateProjectList
				if a.projectList == nil {
					a.projectList = NewProjectListView()
				}
				return a, nil
			}
		case "?":
			if a.state == StateProjectView {
				a.state = StateHelp
				return a, nil
			}
			if a.state == StateHelp {
				a.state = StateProjectView
				return a, nil
			}
		case "esc":
			if a.state == StateHelp {
				a.state = StateProjectView
				return a, nil
			}
			if a.state == StateProjectList {
				a.state = StateProjectView
				return a, nil
			}
			if a.state == StateProductEdit {
				a.state = StateProjectView
				a.restoreProjectViewState()
				return a, nil
			}
			if a.state == StateDeleteDialog {
				a.state = StateProductEdit
				return a, nil
			}
		case "enter":
			if a.state == StateProjectList {
				a.state = StateProjectView
				return a, nil
			}
		case "e":
			if a.state == StateProjectView {
				a.cursorIndex = a.projectView.cursorIndex
				a.viewportOffset = a.projectView.viewport.YOffset
				a.horizontalOffset = a.projectView.horizontalOffset
				product := a.projectView.project.Products[a.projectView.cursorIndex]
				a.productEditView = NewProductEditView(product, a.projectView.project.Categories)
				a.state = StateProductEdit
				return a, nil
			}
		case "n":
			if a.state == StateProjectView {
				a.addNewProduct()
				return a, nil
			}
		}

	case SaveProductMsg:
		a.saveEditedProduct(msg.Product)
		return a, nil

	case CancelEditMsg:
		if a.state == StateProductEdit {
			a.state = StateProjectView
			a.restoreProjectViewState()
		}
		return a, nil

	case DeleteProductMsg:
		if a.state == StateProductEdit {
			a.deleteDialog = NewDeleteDialogView()
			a.state = StateDeleteDialog
		}
		return a, nil

	case ConfirmDeleteMsg:
		a.deleteCurrentProduct()
		return a, nil

	case CancelDeleteMsg:
		if a.state == StateDeleteDialog {
			a.state = StateProductEdit
		}
		return a, nil

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
	}

	switch a.state {
	case StateProjectView:
		if a.projectView != nil {
			model, cmd := a.projectView.Update(msg)
			a.projectView = model.(*ProjectView)
			return a, cmd
		}
	case StateProjectList:
		if a.projectList != nil {
			model, cmd := a.projectList.Update(msg)
			a.projectList = model.(*ProjectListView)
			return a, cmd
		}
	case StateHelp:
		if a.helpView != nil {
			model, cmd := a.helpView.Update(msg)
			a.helpView = model.(*HelpView)
			return a, cmd
		}
	case StateProductEdit:
		if a.productEditView != nil {
			model, cmd := a.productEditView.Update(msg)
			a.productEditView = model.(*ProductEditView)
			return a, cmd
		}
	case StateDeleteDialog:
		if a.deleteDialog != nil {
			model, cmd := a.deleteDialog.Update(msg)
			a.deleteDialog = model.(*DeleteDialogView)
			return a, cmd
		}
	}

	return a, nil
}

func (a *App) View() string {
	var content string

	switch a.state {
	case StateProjectView:
		if a.projectView != nil {
			content = a.projectView.View()
		}
	case StateProjectList:
		if a.projectList != nil {
			content = a.projectList.View()
		}
	case StateHelp:
		if a.helpView != nil {
			content = a.helpView.View()
		}
	case StateProductEdit:
		if a.productEditView != nil {
			content = a.productEditView.View()
		}
	case StateDeleteDialog:
		if a.deleteDialog != nil {
			content = a.deleteDialog.View()
		}
	}

	header := a.renderHeader()
	footer := a.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Top, header, content+"\n", footer)
}

func (a *App) renderHeader() string {
	header := headerStyle.Render("Van Planner - Projekt: " + a.projectName)

	return headerBorderStyle.
		Width(a.width).
		Render(header)
}

func (a *App) renderFooter() string {
	var shortcuts string

	switch a.state {
	case StateProjectView:
		if a.projectView != nil {
			shortcuts = a.projectView.GetShortcuts()
		}
	case StateProjectList:
		if a.projectList != nil {
			shortcuts = a.projectList.GetShortcuts()
		}
	case StateHelp:
		if a.helpView != nil {
			shortcuts = a.helpView.GetShortcuts()
		}
	}

	return footerStyle.
		Width(a.width).
		Render(shortcuts)
}

func (a *App) SetProjectName(name string) {
	a.projectName = name
}

func (a *App) GetCurrentState() AppState {
	return a.state
}

func (a *App) GetProjectView() *ProjectView {
	return a.projectView
}

func (a *App) GetProjectList() *ProjectListView {
	return a.projectList
}

func (a *App) SetProjectList(pl *ProjectListView) {
	a.projectList = pl
}

func (a *App) SetSaveProjectCmd(cmd func() tea.Cmd) {
	a.saveProjectCmd = cmd
}

func (a *App) restoreProjectViewState() {
	if a.projectView != nil {
		a.projectView.cursorIndex = a.cursorIndex
		a.projectView.viewport.YOffset = a.viewportOffset
		a.projectView.horizontalOffset = a.horizontalOffset
	}
}

func (a *App) saveEditedProduct(product model.Product) {
	if a.projectView == nil || a.projectView.project == nil {
		return
	}

	products := a.projectView.project.Products
	if a.cursorIndex < 0 || a.cursorIndex >= len(products) {
		return
	}

	products[a.cursorIndex] = product
	a.projectView.project.Products = products

	a.state = StateProjectView
	a.restoreProjectViewState()

	if a.saveProjectCmd != nil {
		a.saveProjectCmd()
	}
}

func (a *App) addNewProduct() {
	if a.projectView == nil || a.projectView.project == nil {
		return
	}

	products := a.projectView.project.Products
	currentIndex := a.projectView.cursorIndex

	newProduct := model.Product{
		Name:          "Neues Produkt",
		Category:      "",
		EstimatedCost: 0.0,
		ActualCost:    0.0,
		ShippingCost:  0.0,
		Count:         nil,
		ShopLink:      "",
		Notes:         "",
		Completed:     false,
	}

	if len(products) > 0 && currentIndex >= 0 && currentIndex < len(products) {
		newProduct.Category = products[currentIndex].Category
		newProduct.Index = products[currentIndex].Index + 1

		for i := currentIndex + 1; i < len(products); i++ {
			products[i].Index++
		}

		newProducts := make([]model.Product, len(products)+1)
		copy(newProducts, products[:currentIndex+1])
		newProducts[currentIndex+1] = newProduct
		copy(newProducts[currentIndex+2:], products[currentIndex+1:])

		products = newProducts
		a.cursorIndex = currentIndex + 1
	} else {
		if len(products) > 0 {
			newProduct.Index = products[len(products)-1].Index + 1
		} else {
			newProduct.Index = 1
		}
		products = append(products, newProduct)
		a.cursorIndex = len(products) - 1
	}

	a.projectView.project.Products = products
	a.productEditView = NewProductEditView(newProduct, a.projectView.project.Categories)
	a.viewportOffset = a.projectView.viewport.YOffset
	a.horizontalOffset = a.projectView.horizontalOffset
	a.state = StateProductEdit

	if a.saveProjectCmd != nil {
		a.saveProjectCmd()
	}
}

func (a *App) deleteCurrentProduct() {
	if a.projectView == nil || a.projectView.project == nil {
		return
	}

	products := a.projectView.project.Products
	if a.cursorIndex < 0 || a.cursorIndex >= len(products) {
		return
	}

	newProducts := make([]model.Product, len(products)-1)
	copy(newProducts, products[:a.cursorIndex])
	copy(newProducts[a.cursorIndex:], products[a.cursorIndex+1:])

	for i := a.cursorIndex; i < len(newProducts); i++ {
		newProducts[i].Index--
	}

	a.projectView.project.Products = newProducts

	if a.cursorIndex >= len(newProducts) && len(newProducts) > 0 {
		a.cursorIndex = len(newProducts) - 1
	}

	a.state = StateProjectView
	a.restoreProjectViewState()

	if a.saveProjectCmd != nil {
		a.saveProjectCmd()
	}
}
