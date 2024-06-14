package maps

// cannot preduct the order of the keys, but they will all appear in the list
func mapKeys(map_in map[string]string) []string {
	var return_keys []string
	for k := range map_in {
		return_keys = append(return_keys, k)
	}
	return return_keys
}
