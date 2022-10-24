package app

import "fishki/pkg/core"

type answerService struct {
	reviewService core.ReviewService
}

func NewAnswerService(service core.ReviewService) core.AnswerService {
	return &answerService{reviewService: service}
}

func (a answerService) Handle(answer core.Answer) error {
	modeService, err := a.reviewService.GetModeService(answer.ModeID)
	if err != nil {
		return err
	}
	return (*modeService).HandleAnswer(answer)
}
