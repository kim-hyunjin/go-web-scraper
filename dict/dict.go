package dict

import "errors"

type Dictionary map[string]string

var (
	errNotFound = errors.New("Not found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Can't update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to dictionary
func (d Dictionary) Add(word, definition string) error { // map is a hashmap in go. by default a hashmap use pointer. so we don't need to use *
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word in dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word in dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}