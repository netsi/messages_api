package model

type UserType int

const (
	AdminType UserType = iota
	AuthenticatedUserType
)
