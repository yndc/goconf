package utils

func FlattenMap(source map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	TraverseMap(source, func(path *Path, value interface{}) {
		newMap[path.String()] = value
	})
	return newMap
}
