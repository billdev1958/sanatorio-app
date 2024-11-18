package validation

import (
	"errors"
)

/*type UserErrors int

const (

	ErrUserNotFound = errors.New("Usuario no encontrado")
	ErrPasswordInvalid = errors.New("Contraseña invalida")
	ErrInvalidInput = errors.New("Permiso denegado")
)*/

func ValidateLoginData(identifier, pass string) error {
	if identifier == "" {
		return errors.New("debe proporcionar un email")
	}
	if pass == "" {
		return errors.New("falta campo contraseña")
	}

	return nil
}

/*func CheckPermissionUser(ctx context.Context, permissionID int) error {
	claims := auth.ExtractClaims(ctx)
	if claims == nil {
		return "", fmt.Errorf("unauthorized: no claims found in context")
	}

	roleID := claims.
      }
*/
