package app

import (
	"errors"
	"fishki/pkg/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type reviewService struct {
	repo         core.ReviewRepository
	userService  *core.UserService
	modeServices map[core.ReviewModeID]*core.ReviewModeService
}

func NewReviewService(repo core.ReviewRepository, userService *core.UserService) core.ReviewService {
	mmm := make(map[core.ReviewModeID]*core.ReviewModeService)
	return &reviewService{repo: repo, modeServices: mmm, userService: userService}
}

func (r reviewService) CreateReview(review *core.Review, deckID string, userID string) error {
	review.ID = primitive.NewObjectID().Hex()
	//review.Mode = core.ReviewModeFromID(review.ModeID)
	/*mode, err := (*r.modeServices[review.ModeID]).Create(deckID, nil)

	if err != nil {
		return err
	}

	review.Mode = mode*/
	if err := (*r.userService).AddReview(userID, review.ID); err != nil {
		return err
	}

	return r.repo.Create(review)
}

func (r reviewService) Get(id string) (*core.Review, error) {
	return r.repo.Get(id)
}

func (r reviewService) Update(review *core.Review) error {
	return r.repo.Update(review)
}

func (r reviewService) Delete(id string) error {
	return r.repo.Delete(id)
}

func (r reviewService) RegisterModeService(id core.ReviewModeID, service core.ReviewModeService) {
	r.modeServices[id] = &service
}

func (r reviewService) GetModeService(id core.ReviewModeID) (*core.ReviewModeService, error) {
	service, present := r.modeServices[id]
	if !present {
		return nil, errors.New("not found")
	}
	return service, nil
}
