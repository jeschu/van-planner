package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/jens/van-planner/internal/model"
)

type JSONStorage struct {
	filepath string
}

func NewJSONStorage(path string) *JSONStorage {
	return &JSONStorage{
		filepath: path,
	}
}

func (s *JSONStorage) Load() (model.Data, error) {
	data := model.NewData()

	content, err := os.ReadFile(s.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil
		}
		return data, err
	}

	if err := json.Unmarshal(content, &data); err != nil {
		return data, err
	}

	return data, nil
}

func (s *JSONStorage) Save(data model.Data) error {
	dir := filepath.Dir(s.filepath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filepath, content, 0644)
}
