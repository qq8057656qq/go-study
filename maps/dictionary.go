package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (dictionary Dictionary) Search(key string) (string, error) {
	definition, ok := dictionary[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dictionary Dictionary) Add(key, value string) error {
	_, err := dictionary.Search(key)
	switch err {
	case ErrNotFound:
		dictionary[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Update(word, definition string) error {
	_, err := dictionary.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[word] = definition
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)
}

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
