package util

import "os"

// ReadFileForce If file is not exists, will create it.
func ReadFileForce(path string) (*os.File, error) {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		newFile, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		return newFile, nil
	}
	return file, nil
}

// ReadDirForce If dir is not exists, will create it.
func ReadDirForce(path string) (*os.File, error) {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
		return os.Open(path)
	}
	return file, nil
}
