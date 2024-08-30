package user

import "context"

type Usecase interface {
	RegisterUser(ctx context.Context) (string, error)
}
