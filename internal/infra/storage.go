package infra

type FileSystem interface {
	Save(path string, content []byte) error
	Delete(path string) error
	Load(path string) ([]byte, error)
}

type FileSystemImpl struct {
}

func (f *FileSystemImpl) Save(path string, content []byte) error {

	return nil
}

func (f *FileSystemImpl) Delete(path string) error {

	return nil
}

func (f *FileSystemImpl) Load(path string) ([]byte, error) {

	return nil, nil
}
