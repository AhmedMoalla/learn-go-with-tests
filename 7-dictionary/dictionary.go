package dictionary

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	if definition, ok := d[word]; ok {
		return definition, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(word string, definition string) error {
	if _, ok := d[word]; ok {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	if _, ok := d[word]; ok {
		d[word] = definition
		return nil
	}
	return ErrWordDoesNotExist
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
