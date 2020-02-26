package usecase

import (
	"UnityWebPlayer/domain"
)

type GameRepository interface {
	FindById(int) (domain.Game, error)
	FindAll() (domain.Games, error)
}
