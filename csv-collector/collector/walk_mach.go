package collector

import (
	"os"
	"path/filepath"
)

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		matched, err := filepath.Match(pattern, filepath.Base(path))

		if err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return matches, nil
}
