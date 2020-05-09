package main

import "errors"

// ErrNotFound if not found
var ErrNotFound = errors.New("Err: could not find the word you were looking for")

// Search returns index of str in dict
func (d Dictionary) Search(str string) (string, error) {
	def, found := d[str]
	if !found {
		return "", ErrNotFound
	}
	return def, nil
}

// Add d[key] = val
func (d Dictionary) Add(key, val string) {
	d[key] = val
}

// Dictionary is map of strings giving strings
type Dictionary map[string]string

func main() {

}
