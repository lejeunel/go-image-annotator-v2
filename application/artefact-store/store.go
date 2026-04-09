package artefact_store

import (
	"fmt"
	"os"
	"path/filepath"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

type FileArtefactRepo struct {
	baseDir string
}

func NewFileArtefactRepo(baseDir string) *FileArtefactRepo {
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create base directory: %v", err))
	}
	return &FileArtefactRepo{baseDir: baseDir}
}

func (r *FileArtefactRepo) filePath(id im.ImageId) string {
	return filepath.Join(r.baseDir, fmt.Sprintf("%s", id.String()))
}

func (r *FileArtefactRepo) Store(id im.ImageId, data []byte) error {
	path := r.filePath(id)
	return os.WriteFile(path, data, 0644)
}

func (r *FileArtefactRepo) Delete(id im.ImageId) error {
	path := r.filePath(id)
	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("artefact not found: %w", err)
		}
		return err
	}
	return nil
}

func (r *FileArtefactRepo) Get(id im.ImageId) ([]byte, error) {
	path := r.filePath(id)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("artefact not found: %w", err)
		}
		return nil, err
	}
	return data, nil
}
