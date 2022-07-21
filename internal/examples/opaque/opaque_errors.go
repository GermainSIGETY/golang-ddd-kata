package opaque

import "errors"

func generic_error() error {
	x, err := something()
	if err != nil {
		return err
	}
	print(x)
	return nil
}

func something() (string, error) {
	return "", errors.New("error")
}

// Avantage / inconvénient ?

// ++++ Pas de dépendance entre les packages
// ---- Pas de contexte

type temporary interface {
	Temporary() bool
}

// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

// ---- Plus de contexte
