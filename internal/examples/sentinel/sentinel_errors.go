package sentinel

import (
	"database/sql"
	"errors"
	"io"
)

var myError = errors.New("EOF")

func init() {
	io.EOF = nil // haha!
	// sqlError := sql.ErrNoRows
}

func usage() {
	err := someWork()
	if err == io.EOF {
		print("handle error")
	}

	if err == sql.ErrNoRows {
		print("Handle error")
	}
}

func someWork() error {
	return errors.New("error")
}

// Avantage / inconvénient ?

// ++++ Une constance simple à utilser, réutiliser

// ---- Dépendances entre les packages
// ---- Comment ajouter plus de contexte d'erreur
