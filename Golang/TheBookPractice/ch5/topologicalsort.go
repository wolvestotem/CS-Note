package ch5

func TopologicalSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	//sort.Strings(keys)
	visitAll(keys)
	/*var visitAll func(map[string]bool)
	visitAll = func(items map[string]bool) {
		for item, _ := range items {
			if !items[item] {
				items[item] = true
				for _, i := range m[item] {
					items[i] = items[i]
				}
				visitAll(items)
				order = append(order, item)
			}
		}
	}
	keys := make(map[string]bool)
	for key := range m {
		keys[key] = false
	}
	visitAll(keys)*/
	return order
}
