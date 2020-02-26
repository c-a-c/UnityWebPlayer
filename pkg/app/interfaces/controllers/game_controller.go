package controllers

import (
	"UnityWebPlayer/interfaces/database"
	"UnityWebPlayer/usecase"
	"log"
	"strconv"
)

type GameController struct {
	Interactor usecase.GameInteractor
}

func NewGameController(sqlHandler database.SqlHandler) *GameController {
	return &GameController{
		Interactor: usecase.GameInteractor{
			GameRepository: &database.GameRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *GameController) Index(c Context) {
	games, err := controller.Interactor.Games()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(200, games)
}

func (controller *GameController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	game, err := controller.Interactor.GameById(id)
	log.Println("<", err.Error(), ">")
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(200, game)
}
