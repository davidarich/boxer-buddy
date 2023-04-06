package path

import "golang.org/x/sys/windows/registry"

// interpolates environment variables in the path with their values
func Expand(path string) (string, error) {
	return registry.ExpandString(path)
}
