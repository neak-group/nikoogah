package dto

import "github.com/neak-group/nikoogah/utils/uuid"

type AddRepresentativeParams struct {
	CharityID uuid.UUID
	UserID    uuid.UUID
}

type RemoveRepresentativeParams struct {
	CharityID uuid.UUID
	UserID    uuid.UUID
}

type CheckRepresentativeAccessParams struct {
	CharityID uuid.UUID
	UserID    uuid.UUID
	AccessKey string
}
