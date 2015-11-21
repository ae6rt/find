package find

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Find finds regular files recursively under root whose base name matches the provided regular expression.  Files residing
// at a depth other than atDepth are ignored.  If atDepth is -1, depth is not considered.  A depth of 1 means find files directly beneath
// root.  ignoreDirectories is an array of prefixes that cause a directory to be recursively ignored if its name starts with the prefix.
func Find(root string, expression *regexp.Regexp, atDepth int, ignoreDirectories []string) ([]string, error) {
	if strings.HasSuffix(root, "/") {
		return nil, fmt.Errorf("Root must not end in /: %s\n", root)
	}

	// files of interest reside at a fixed depth of 3 below root
	slashOffset := strings.Count(root, "/") + atDepth - 1
	checkDepth := atDepth != -1

	var files []string
	markFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, v := range ignoreDirectories {
			if strings.HasSuffix(path, v) {
				return filepath.SkipDir
			}
		}

		// Ignore paths too deep if atDepth is specified
		if atDepth != -1 && info.IsDir() && strings.Count(path, "/") > slashOffset {
			return filepath.SkipDir
		}

		if info.Mode().IsRegular() {
			ok := expression.MatchString(info.Name())
			if ok {
				if checkDepth {
					if strings.Count(path, "/") == slashOffset {
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
