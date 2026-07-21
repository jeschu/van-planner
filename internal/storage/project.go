package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/jens/van-planner/internal/model"
)

type ProjectStorage struct {
	projectsDir string
}

func NewProjectStorage(projectsDir string) *ProjectStorage {
	return &ProjectStorage{
		projectsDir: projectsDir,
	}
}

func (s *ProjectStorage) Load(name string) (*model.Project, error) {
	path := filepath.Join(s.projectsDir, name+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var project model.Project
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *ProjectStorage) Save(name string, project *model.Project) error {
	path := filepath.Join(s.projectsDir, name+".json")
	data, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (s *ProjectStorage) List() ([]string, error) {
	entries, err := os.ReadDir(s.projectsDir)
	if err != nil {
		return nil, err
	}

	var projects []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == ".json" && entry.Name() != "config.json" {
			projects = append(projects, entry.Name()[:len(entry.Name())-5])
		}
	}

	return projects, nil
}
