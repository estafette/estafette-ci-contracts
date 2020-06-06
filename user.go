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
	Organizations   []*UserOrganization    `json:"organizations,omitempty"`
	Groups          []*UserGroup           `json:"groups,omitempty"`
	Roles           []*string              `json:"roles,omitempty"`
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

// UserOrganization represents an organization that uses a multi-tenancy installation
type UserOrganization struct {
	Name string `json:"name,omitempty"`
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

// HasRole returns true if a user has the parameterized role
func (u *User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r != nil && *r == role {
			return true
		}
	}
	return false
}

// AddRole adds a role if it's not present
func (u *User) AddRole(role string) {
	if !u.HasRole(role) {
		u.Roles = append(u.Roles)
	}
}

// RemoveRole removes a role if it's present
func (u *User) RemoveRole(role string) {
	remainingRoles := []*string{}
	for _, r := range u.Roles {
		if r != nil && *r != role {
			remainingRoles = append(remainingRoles, r)
		}
	}

	u.Roles = remainingRoles
}
