package artefact_store

import (
	"fmt"
	"os"
	"path/filepath"

	im "github.com/lejeunel/go-image-annotator-v2/entities/image"
)

// FileArtefactRepo implements ArtefactRepo and ArtefactReadRepo using the local file system
type FileArtefactRepo struct {
	baseDir string
}

// NewFileArtefactRepo creates a new file-based artefact repo
func NewFileArtefactRepo(baseDir string) *FileArtefactRepo {
	// Ensure the directory exists
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create base directory: %v", err))
	}
	return &FileArtefactRepo{baseDir: baseDir}
}

// filePath returns the full path for a given ImageId
func (r *FileArtefactRepo) filePath(id im.ImageId) string {
	return filepath.Join(r.baseDir, fmt.Sprintf("%s", id.String()))
}

// Store writes the byte slice to a file
func (r *FileArtefactRepo) Store(id im.ImageId, data []byte) error {
	path := r.filePath(id)
	return os.WriteFile(path, data, 0644)
}

// Delete removes the file corresponding to the ImageId
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

// Get reads the contents of the file
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
