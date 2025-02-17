package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type DisplayIdentifier struct {
	DisplayName    string
	DisplayNameSet bool
}

func parseDisplayIdentifier(tok *tokenGenerator) (DisplayIdentifier, error) {
	di := DisplayIdentifier{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("displayIdentifier could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("displayIdentifier could not be parsed")
	} else if !di.DisplayNameSet {
		di.DisplayName = tok.current()
		di.DisplayNameSet = true
		log.Info().Msg("displayIdentifier displayName successfully parsed")
	}
	return di, err
}
