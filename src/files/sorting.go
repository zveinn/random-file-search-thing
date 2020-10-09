package files

import "sort"

func SortBySize(filemap map[string]File) (sizes []int64) {
	for _, v := range filemap {
		if v.IsDir {
			continue
		}
		sizes = append(sizes, v.Size)
	}

	sort.Slice(sizes, func(a, b int) bool { return sizes[a] > sizes[b] })

	return
}
