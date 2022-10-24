package app

import (
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"fishki/pkg/core"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
)

type userService struct {
	repo core.UserRepository
}

func NewUserService(repo core.UserRepository) core.UserService {
	return &userService{repo}
}

func (u userService) Login(username string, password string) (*core.User, error) {
	user, err := u.GetByUsername(username)

	if err != nil {
		return nil, err
	}

	goodPwd := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))

	if goodPwd != nil {
		return nil, errors.New("incorrect password")
	}

	sessionId := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, sessionId); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, nil
	}
	user.Session = sessionId
	if err = u.repo.Update(user); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	return user, nil
}

func (u userService) Logout(username string, session []byte) error {
	user, err := u.GetByUsername(username)

	if err != nil {
		return err
	}

	if subtle.ConstantTimeCompare(user.Session, []byte(session)) != 1 {
		return errors.New("invalid session id")
	}

	user.Session = nil

	return nil
}

func (u userService) Register(username string, email string, password string) (*core.User, error) {
	usu, _ := u.GetByUsername(username)
	sus, _ := u.GetByEmail(email)

	if usu.Username == username || sus.Email == email {
		return nil, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	if err != nil {
		return nil, err
	}

	user := &core.User{}
	user.ID = primitive.NewObjectID().Hex()
	user.Username = username
	user.Email = email
	user.PasswordHash = hash
	user.DecksID = make(core.Set[string])
	user.DecksID["aaa"] = struct{}{}

	err = u.repo.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userService) Get(id string) (*core.User, error) {
	return u.repo.Get(id)
}

func (u userService) GetByUsername(username string) (*core.User, error) {
	return u.repo.GetByUsername(username)
}

func (u userService) GetByEmail(email string) (*core.User, error) {
	return u.repo.GetByEmail(email)
}

func (u userService) Update(user *core.User) error {
	return u.repo.Update(user)
}

func (u userService) Delete(username string) error {
	return u.repo.Delete(username)
}

func (u userService) AddDeck(id string, deckID string) error {
	return u.repo.AddDeck(id, deckID)
}

func (u userService) RemoveDeck(id string, deckID string) error {
	return u.repo.RemoveDeck(id, deckID)
}

func (u userService) AddReview(id string, reviewID string) error {
	return u.repo.AddReview(id, reviewID)
}

func (u userService) RemoveReview(id string, reviewID string) error {
	return u.repo.RemoveReview(id, reviewID)
}
