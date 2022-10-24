package core

import "github.com/gin-gonic/gin"

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Email    string `json:"-" bson:"email,omitepmty"`

	PasswordHash []byte `json:"-" bson:"pwd_hash,omitempty"`
	//PasswordSalt string `json:"-" bson:"pwd_salt,omitempty"`
	Session []byte `json:"-" bson:"session,omitempty"`

	DecksID   Set[string] `json:"decks" bson:"decks,omitempty"`
	ReviewsID Set[string] `json:"reviews" bson:"reviews,omitempty"`
}

type UserCredentials struct {
}

type UserRepository interface {
	Create(user *User) error
	Get(id string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)

	Update(user *User) error
	Delete(username string) error

	AddDeck(id string, deckID string) error
	RemoveDeck(username string, deckID string) error

	AddReview(id string, reviewID string) error
	RemoveReview(username string, reviewID string) error
}

type UserService interface {
	Login(username string, password string) (*User, error)
	Logout(username string, session []byte) error
	Register(username string, email string, password string) (*User, error)

	Get(id string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(username string) error

	AddDeck(id string, deckID string) error
	RemoveDeck(i string, deckID string) error

	AddReview(id string, reviewID string) error
	RemoveReview(id string, reviewID string) error
}

type UserHandler interface {
	Handler
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Register(ctx *gin.Context)
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
