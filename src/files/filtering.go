package files

// FilterOnName ...
func FilterOnName(fileMap map[string]File, f func(name string) bool) {
	for i, v := range fileMap {
		if f(v.Name) {
			delete(fileMap, i)
		}
	}
}
