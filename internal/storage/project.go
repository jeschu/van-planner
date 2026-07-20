package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/jens/van-planner/internal/model"
)

const ProjectsDir = "projekte"

type ProjectInfo struct {
	Name     string `json:"name"`
	FilePath string `json:"-"`
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
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
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
	return os.Remove(filePath)
}

func (pm *ProjectManager) ProjectExists(name string) bool {
	filePath := filepath.Join(pm.baseDir, name+".json")
	_, err := os.Stat(filePath)
	return err == nil
}
