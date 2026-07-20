package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/jens/van-planner/internal/model"
)

const (
	ProjectsDir   = "projekte"
	AppConfigFile = "van-planner.config.json"
)

type ProjectInfo struct {
	Name     string `json:"name"`
	FilePath string `json:"-"`
}

type AppConfig struct {
	LastProject string `json:"lastProject,omitempty"`
}

type ProjectManager struct {
	baseDir string
}

func NewProjectManager() *ProjectManager {
	return &ProjectManager{
		baseDir: ProjectsDir,
	}
}

func (pm *ProjectManager) ensureDir() error {
	return os.MkdirAll(pm.baseDir, 0755)
}

func (pm *ProjectManager) ListProjects() ([]ProjectInfo, error) {
	if err := pm.ensureDir(); err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(pm.baseDir)
	if err != nil {
		return nil, err
	}

	var projects []ProjectInfo
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") && !strings.HasSuffix(entry.Name(), ".config.json") {
			name := strings.TrimSuffix(entry.Name(), ".json")
			projects = append(projects, ProjectInfo{
				Name:     name,
				FilePath: filepath.Join(pm.baseDir, entry.Name()),
			})
		}
	}

	return projects, nil
}

func (pm *ProjectManager) LoadProject(name string) (model.Data, error) {
	filePath := filepath.Join(pm.baseDir, name+".json")
	storage := NewJSONStorage(filePath)
	return storage.Load()
}

func (pm *ProjectManager) SaveProject(name string, data model.Data) error {
	if err := pm.ensureDir(); err != nil {
		return err
	}

	filePath := filepath.Join(pm.baseDir, name+".json")
	storage := NewJSONStorage(filePath)
	return storage.Save(data)
}

func (pm *ProjectManager) CreateProject(name string) error {
	if err := pm.ensureDir(); err != nil {
		return err
	}

	filePath := filepath.Join(pm.baseDir, name+".json")

	if _, err := os.Stat(filePath); err == nil {
		return os.ErrExist
	}

	data := model.NewData()
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, dataJSON, 0644)
}

func (pm *ProjectManager) DeleteProject(name string) error {
	filePath := filepath.Join(pm.baseDir, name+".json")
	configPath := filepath.Join(pm.baseDir, name+".config.json")

	os.Remove(configPath)
	return os.Remove(filePath)
}

func (pm *ProjectManager) ProjectExists(name string) bool {
	filePath := filepath.Join(pm.baseDir, name+".json")
	_, err := os.Stat(filePath)
	return err == nil
}

func (pm *ProjectManager) LoadProjectConfig(name string) (model.ProjectConfig, error) {
	data, err := pm.LoadProject(name)
	if err != nil {
		return model.NewProjectConfig(), nil
	}

	return model.ProjectConfig{
		LastProductID: data.LastProductID,
	}, nil
}

func (pm *ProjectManager) SaveProjectConfig(name string, config model.ProjectConfig) error {
	data, err := pm.LoadProject(name)
	if err != nil {
		return err
	}

	data.LastProductID = config.LastProductID

	return pm.SaveProject(name, data)
}

func (pm *ProjectManager) LoadAppConfig() (AppConfig, error) {
	filePath := AppConfigFile
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return AppConfig{}, nil
		}
		return AppConfig{}, err
	}

	var config AppConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return AppConfig{}, err
	}

	return config, nil
}

func (pm *ProjectManager) SaveAppConfig(config AppConfig) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(AppConfigFile, data, 0644)
}
