package controller

import (
	"github.com/jens/van-planner/internal/model"
	"github.com/jens/van-planner/internal/storage"
	"github.com/jens/van-planner/internal/ui"
)

type Controller struct {
	configStorage  *storage.ConfigStorage
	projectStorage *storage.ProjectStorage
	config         *storage.Config
	currentProject *model.Project
	projectName    string
	app            *ui.App
}

func NewController(configStorage *storage.ConfigStorage, projectStorage *storage.ProjectStorage) *Controller {
	return &Controller{
		configStorage:  configStorage,
		projectStorage: projectStorage,
	}
}

func (c *Controller) Initialize() error {
	config, err := c.configStorage.Load()
	if err != nil {
		config = &storage.Config{
			DefaultCategories: []string{
				"Fahrzeug", "Werkzeug", "Dämmung", "Heizung",
				"Wasser", "Fenster", "Elektrik",
			},
			Projects:    []string{},
			LastProject: 0,
		}
	}
	c.config = config

	if err := c.loadCurrentProject(); err != nil {
		return err
	}

	return nil
}

func (c *Controller) loadCurrentProject() error {
	projectName := c.getCurrentProjectName()
	c.projectName = projectName

	project, err := c.projectStorage.Load(projectName)
	if err != nil {
		project = &model.Project{
			Categories: c.config.DefaultCategories,
			Products:   []model.Product{},
		}
	}
	c.currentProject = project

	return nil
}

func (c *Controller) getCurrentProjectName() string {
	if len(c.config.Projects) > 0 {
		if c.config.LastProject >= 0 && c.config.LastProject < len(c.config.Projects) {
			return c.config.Projects[c.config.LastProject]
		}
		return c.config.Projects[0]
	}
	return "Neues Projekt"
}

func (c *Controller) CreateApp() *ui.App {
	projectView := ui.NewProjectView(c.currentProject)
	app := ui.NewApp(c.projectName, projectView)
	c.app = app
	return app
}

func (c *Controller) GetApp() *ui.App {
	return c.app
}

func (c *Controller) SaveCurrentState() error {
	if c.app == nil {
		return nil
	}

	if c.app.GetCurrentState() == ui.StateProjectView {
		if pv := c.app.GetProjectView(); pv != nil {
			c.currentProject = pv.GetProject()
			if err := c.projectStorage.Save(c.projectName, c.currentProject); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Controller) SwitchProject(projectName string) error {
	c.projectName = projectName

	project, err := c.projectStorage.Load(projectName)
	if err != nil {
		return err
	}
	c.currentProject = project

	if c.app != nil {
		c.app.SetProjectName(projectName)
		if pv := c.app.GetProjectView(); pv != nil {
			pv.SetProject(c.currentProject)
		}
	}

	for i, p := range c.config.Projects {
		if p == projectName {
			c.config.LastProject = i
			break
		}
	}

	return c.configStorage.Save(c.config)
}

func (c *Controller) GetProjectList() ([]string, error) {
	return c.projectStorage.List()
}

func (c *Controller) GetCurrentProject() *model.Project {
	return c.currentProject
}

func (c *Controller) GetProjectName() string {
	return c.projectName
}

func (c *Controller) HandleProjectSelection(projectName string) error {
	return c.SwitchProject(projectName)
}
