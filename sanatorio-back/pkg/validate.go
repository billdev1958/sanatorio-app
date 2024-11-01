package validation

import "errors"

func ValidateLoginData(identifier, pass string) error {
	if identifier == "" {
		return errors.New("debe proporcionar un email")
	}
	if pass == "" {
		return errors.New("falta campo contraseña")
	}

	return nil
}
