package usecase

import (
	"UnityWebPlayer/domain"
)

type GameInteractor struct {
	GameRepository GameRepository
}

func (interactor *GameInteractor) GameById(identifier int) (game domain.Game, err error) {
	game, err = interactor.GameRepository.FindById(identifier)
	return
}

func (interactor *GameInteractor) Games() (games domain.Games, err error) {
	games, err = interactor.GameRepository.FindAll()
	return
}
