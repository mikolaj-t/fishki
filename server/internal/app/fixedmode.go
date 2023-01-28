package app

import (
	"fishki/pkg/core"
)

type fixedModeService struct {
	repo        core.FixedModeRepository
	deckService core.DeckService
}

func NewFixedModeService(repo core.FixedModeRepository, deckService core.DeckService) core.FixedModeService {
	return &fixedModeService{repo, deckService}
}

func (f fixedModeService) Create(deckID string, intervals []uint16) (*core.FixedMode, error) {
	deck, err := f.deckService.Get(deckID)

	if err != nil {
		return nil, err
	}

	fixedMode := &core.FixedMode{}
	fixedMode.Intervals = intervals

	fixedMode.Levels = make(map[string]uint8)
	fixedMode.Dates = make(map[string]core.Date)

	for _, s := range deck.CardsID {
		fixedMode.Levels[s] = 0
		fixedMode.Dates[s] = core.Today()
	}

	return fixedMode, nil
}

func (f fixedModeService) HandleAnswer(answer core.Answer) error {
	mode, err := f.repo.Get(answer.ReviewID)

	if err != nil {
		return err
	}

	if answer.Correct {
		mode.Levels[answer.CardID]++
		level := mode.Levels[answer.CardID]

		intervalsSize := len(mode.Intervals)

		// card has finished progression
		if level > uint8(intervalsSize) {
			return nil
		}

		date := core.Today() + core.Date(mode.Intervals[level-1])
		f.setCardDate(mode, date, answer.CardID)
	} else {
		// reset the card's progress if the answer is incorrect
		mode.Levels[answer.CardID] = 0
		f.setCardDate(mode, core.Today(), answer.CardID)
	}

	return f.repo.Update(answer.ReviewID, mode)
}

func (f fixedModeService) setCardDate(mode *core.FixedMode, date core.Date, cardID string) {
	if mode.Dates == nil {
		mode.Dates = make(map[string]core.Date)
	}

	mode.Dates[cardID] = date
}
