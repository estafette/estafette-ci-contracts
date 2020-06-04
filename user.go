package contracts

import "time"

// User represents a user of Estafette
type User struct {
	ID     string `json:"id,omitempty"`
	Active bool   `json:"active,omitempty"`
	// Name is derived from the first identity with a name
	Name string `json:"name,omitempty"`
	// Email is derived from the first identity with an email address
	Email           string                 `json:"email,omitempty"`
	Identities      []*UserIdentity        `json:"identities,omitempty"`
	Groups          []*UserGroup           `json:"groups,omitempty"`
	Preferences     map[string]interface{} `json:"preferences,omitempty"`
	FirstVisit      *time.Time             `json:"firstVisit,omitempty"`
	LastVisit       *time.Time             `json:"lastVisit,omitempty"`
	CurrentProvider string                 `json:"currentProvider,omitempty"`
}

// UserIdentity represents the various identities a user can have in different systems
type UserIdentity struct {
	Provider string `json:"provider,omitempty"`
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

// UserGroup represents a group of users as configured in different systems
type UserGroup struct {
	Provider string `json:"provider,omitempty"`
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
}

// GetEmail returns the first identity email address
func (u *User) GetEmail() string {
	if u.Identities != nil && len(u.Identities) > 0 {
		for _, i := range u.Identities {
			if i.Email != "" {
				return i.Email
			}
		}
	}

	return ""
}

// GetProvider returns the first identity provider
func (u *User) GetProvider() string {
	if u.Identities != nil && len(u.Identities) > 0 {
		for _, i := range u.Identities {
			if i.Provider != "" {
				return i.Provider
			}
		}
	}

	return ""
}

// GetName returns the first identity name
func (u *User) GetName() string {
	if u.Identities != nil && len(u.Identities) > 0 {
		for _, i := range u.Identities {
			if i.Name != "" {
				return i.Name
			}
		}
	}

	return ""
}
