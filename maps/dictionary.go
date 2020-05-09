package main

const (
	// ErrNotFound if not found
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists word exists in dict
	ErrWordExists = DictionaryErr("cannot add word because it already exists")

	// ErrWordDoesNotExist for updates
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// Dictionary is map of strings giving strings
type Dictionary map[string]string

// DictionaryErr type of error
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search returns index of str in dict
func (d Dictionary) Search(str string) (string, error) {
	def, found := d[str]
	if !found {
		return "", ErrNotFound
	}
	return def, nil
}

// Add d[key] = val
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update key value in dict
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func main() {

}
