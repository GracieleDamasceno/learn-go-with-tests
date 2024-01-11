package maps

import "errors"

type Dictionary map[string]string

var (
	ErrorWordNotFound = errors.New("expected to find word")
	ErrorWordExists   = errors.New("word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrorWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorWordNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}
