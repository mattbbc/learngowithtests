package maps

type Dictionary map[string]string

// Create our own error
// First declare our own type
type DictionaryErr string

// Create what we need
const (
	ErrWordNotFound      = DictionaryErr("no find word")
	ErrWordExistsAlready = DictionaryErr("word exists already")
	ErrWordDoesNotExist  = DictionaryErr("word does not exist")
)

// Then implement error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

// Maps are already (sort of) reference types, so no need to
// use pointer receiver
// They're not quite reference types though?
// They are pointers to runtime types.
// https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
func (d Dictionary) Search(term string) (string, error) {
	result, ok := d[term]

	if !ok {
		return "", ErrWordNotFound
	}

	return result, nil
}

func (d Dictionary) Add(term, definition string) error {
	_, err := d.Search(term)

	switch err {
	case ErrWordNotFound:
		d[term] = definition
	case nil:
		return ErrWordExistsAlready
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(term, newDef string) error {
	_, err := d.Search(term)

	switch err {
	case ErrWordNotFound:
		// We could reuse ErrWordNotFound but it is better to have specific errors for specific cases
		return ErrWordDoesNotExist
	case nil:
		d[term] = newDef
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(term string) {
	delete(d, term)
}
