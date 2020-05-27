package contracts

// User represents a user of Estafette
type User struct {
	Name       string         `json:"name,omitempty"`
	Active     bool           `json:"active,omitempty"`
	Identities []UserIdentity `json:"identities,omitempty"`
	Groups     []UserGroup    `json:"groups,omitempty"`
}

// UserIdentity represents the various identities a user can have in different systems
type UserIdentity struct {
	Source   string `json:"source,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

// UserGroup represents a group of users as configured in different systems
type UserGroup struct {
	Source string `json:"source,omitempty"`
	Name   string `json:"name,omitempty"`
}
