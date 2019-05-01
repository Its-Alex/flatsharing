package main

// UserAuth is struct used to authentification
type UserAuth struct {
	OrganisationID string
	UserID         string
}

// ContextKey define the type of context key used in context.WithValue
type ContextKey string
