package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("didn't find what you wanted")

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {

		return "", ErrNotFound
	}
	return result, nil
}

func (d Dictionary) Add(key string, data string) {
	d[key] = data
}
