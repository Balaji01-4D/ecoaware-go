package models

type Status string

const (
	PENDING     Status = "PENDING"
	IN_PROGRESS Status = "IN_PROGRESS"
	RESOLVED    Status = "RESOLVED"
	REJECTED    Status = "REJECTED"
)
