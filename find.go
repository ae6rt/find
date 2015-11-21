package find

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Find finds recursively under a root directory those regular files whose base name matches the provided regular expression.  Files residing
// at a depth other than atDepth are ignored.  For a root R, this file F resides at depth 2: R/t/F, for some subdirectory t of R.
// If atDepth is -1, depth is not considered.  ignoreDirectories is an array of strings representing base directory names that are ignored in the search.
// For example, a directory name in .git will be ignored if ".git" is an element in ignoreDirectories.
func Find(root string, expression *regexp.Regexp, atDepth int, ignoreDirectories []string) ([]string, error) {
	if strings.HasSuffix(root, "/") {
		return nil, fmt.Errorf("Root must not end in /: %s\n", root)
	}

	checkDepth := atDepth != -1

	var files []string
	markFn := func(path string, info os.FileInfo, err error) error {

		normalizedPath := path[len(root):]

		parts := strings.Count(normalizedPath, "/")

		if err != nil {
			return err
		}

		for _, v := range ignoreDirectories {
			if strings.HasSuffix(normalizedPath, v) {
				return filepath.SkipDir
			}
		}

		if checkDepth && info.IsDir() && strings.Count(normalizedPath, "/") > parts {
			return filepath.SkipDir
		}

		if info.Mode().IsRegular() {
			ok := expression.MatchString(info.Name())
			if ok {
				if checkDepth {
					if atDepth == parts {
						files = append(files, path)
					}
				} else {
					files = append(files, path)
				}
			}
		}
		return nil
	}

	err := filepath.Walk(root, markFn)
	if err != nil {
		return nil, err
	}

	return files, nil
}
