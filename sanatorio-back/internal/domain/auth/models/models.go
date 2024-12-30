package models

import "github.com/google/uuid"

type CheckPermission struct {
	AccountID  uuid.UUID
	Permission int
}
