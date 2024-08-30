package user

import "context"

type Repository interface {
	RegisterUser(ctx context.Context) (string, error)
}
