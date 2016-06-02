package models

import (
	"github.com/revel/revel"
	"time"
)

type User struct {
	UserID         uint      `json:"user_id" db:"user_id,omitempty"`
	Name           string    `json:"name" db:"name"`
	Password       string    `json:"password" db:"-"`
	Hash           []byte    `json:"-" db:"hash"`
	Salt           []byte    `json:"-" db:"salt"`
	DefaultAccount uint      `json:"default_account" db:"default_account"`
	Email          string    `json:"email" db:"email"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Session struct {
	UserID uint   `db:"user_id"`
	Token  string `db:"token"`
}

func (user *User) Validate(v *revel.Validation) {
	v.Required(user.Name)
	v.MinSize(user.Name, 4)
	v.Required(user.Password)
	v.MinSize(user.Password, 4)
	v.Required(user.Email)
	v.Email(user.Email)
	// v.Required(user.PasswordConfirm)
	// v.Required(user.PasswordConfirm == user.Password).
	//     Message("The passwords do not match.")
	// v.Required(user.EmailConfirm)
	// v.Required(user.EmailConfirm == user.Email).
	//     Message("The email addresses do not match")
	// v.Required(user.TermsOfUse)
}

func (c *User) NewFill() {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *User) UpdateFill() {
	c.UpdatedAt = time.Now()
}

// func (c *User) MarshalJSON() ([]byte, error) {
//     type Alias User
//     return json.Marshal(&struct {
//         *Alias
// 		CreatedAt    string `json:"created_at"`
// 		UpdatedAt    string `json:"updated_at"`
//     }{
//         Alias: (*Alias)(c),
// 		CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
// 		UpdatedAt: c.UpdatedAt.Format("2006-01-02 15:04:05"),
//     })
// }
