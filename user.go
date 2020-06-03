package contracts

import "time"

// User represents a user of Estafette
type User struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Active      bool                   `json:"active,omitempty"`
	Identities  []*UserIdentity        `json:"identities,omitempty"`
	Groups      []*UserGroup           `json:"groups,omitempty"`
	Preferences map[string]interface{} `json:"preferences,omitempty"`
	FirstVisit  *time.Time             `json:"firstVisit,omitempty"`
	LastVisit   *time.Time             `json:"lastVisit,omitempty"`
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
