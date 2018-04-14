package pathupdate

import "regexp"

// PathUpdate updates the value identified by path within a nested map.
func PathUpdate(path string, m *interface{}, value interface{}) {
	matches := re.FindAllStringSubmatch(path, -1)
	components := make([]string, 0, len(matches))
	for _, m := range matches {
		components = append(components, m[1])
	}
	pathUpdate(components, m, value)
}

var re = regexp.MustCompile("/([^/]+)")

func pathUpdate(components []string, m *interface{}, value interface{}) {
	switch len(components) {
	case 0:
		*m = value
	default:
		curr, ok := (*m).(map[string]interface{})
		if !ok {
			curr = map[string]interface{}{}
			*m = curr
		}
		var v interface{}
		v = curr[components[0]]
		pathUpdate(components[1:len(components)], &v, value)
		if v == nil {
			delete(curr, components[0])
		} else {
			curr[components[0]] = v
		}
	}
}
