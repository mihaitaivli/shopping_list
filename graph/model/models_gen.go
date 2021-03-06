// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type EditItem struct {
	ID        string  `json:"id"`
	List      *string `json:"list"`
	Content   *string `json:"content"`
	Purchased *bool   `json:"purchased"`
	Recurring *bool   `json:"recurring"`
	Postponed *bool   `json:"postponed"`
}

type EditList struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EditSharedList struct {
	ListID  string   `json:"listId"`
	UsersID []string `json:"usersId"`
}

type Item struct {
	ID        string    `json:"id"`
	List      *List     `json:"list"`
	Content   string    `json:"content"`
	Purchased bool      `json:"purchased"`
	Recurring bool      `json:"recurring"`
	Postponed bool      `json:"postponed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LinkUserInput struct {
	LinkedUserID string   `json:"linkedUserId"`
	SharedLists  []string `json:"sharedLists"`
}

type List struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Items     []*Item   `json:"items"`
	Users     []*User   `json:"users"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewItem struct {
	List      string `json:"list"`
	Content   string `json:"content"`
	Purchased *bool  `json:"purchased"`
	Recurring *bool  `json:"recurring"`
	Postponed *bool  `json:"postponed"`
}

type NewUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	AccessToken string `json:"accessToken"`
}

type User struct {
	ID string `json:"id"`
	// users are linked in order to share lists
	LinkedUsers []*User `json:"linkedUsers"`
	List        []*List `json:"list"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	// Can be a password, pin, etc...
	AccessToken string    `json:"accessToken"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
