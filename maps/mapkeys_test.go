package maps

import (
	"slices"
	"testing"
)

func TestMapKeys(t *testing.T) {
	testMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	keysList := mapKeys(testMap)
	wanted := "key2"
	//use 'contains' below because the order of keys is indeterminate, thus can't look at a particular location for a particular key
	assertion := slices.Contains(keysList, wanted)

	if !assertion {
		t.Errorf("wanted: %v, but its not found in %v", wanted, keysList)
	}
}
