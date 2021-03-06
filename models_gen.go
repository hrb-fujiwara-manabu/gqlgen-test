// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen_test

type Profile struct {
	ID          string `json:"id"`
	Owner       *User  `json:"owner"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ChangedBy   *User  `json:"changedBy"`
}

type ProfileInput struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

type Purchase struct {
	ID        string `json:"id"`
	Owner     *User  `json:"owner"`
	Has       bool   `json:"has"`
	StartedAt string `json:"startedAt"`
}

type PurchaseInput struct {
	Has bool `json:"has"`
}

type User struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Phone     string     `json:"phone"`
	Profiles  []*Profile `json:"profiles"`
	Purchase  *Purchase  `json:"purchase"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
	DeletedAt *string    `json:"deletedAt"`
}

type UserInput struct {
	Email     string          `json:"email"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Phone     *string         `json:"phone"`
	Profiles  []*ProfileInput `json:"profiles"`
	Purchase  *PurchaseInput  `json:"purchase"`
}
