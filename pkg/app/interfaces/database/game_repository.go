package database

import (
	"UnityWebPlayer/domain"
)

type GameRepository struct {
	SqlHandler
}

func (repo *GameRepository) FindById(identifier int) (game domain.Game, err error) {
	row, err := repo.Query(
		"select id, title, thumbneil from games where id=?;", identifier,
	)

	if err != nil {
		return
	}
	defer row.Close()

	var id int
	var title string
	var thumbneil string
	if err = row.Scan(&id, &title, &thumbneil); err != nil {
		return
	}

	game.ID = id
	game.Title = title
	game.ThumbneilPath = thumbneil
	return
}

func (repo *GameRepository) FindAll() (games domain.Games, err error) {
	rows, err := repo.Query(
		"select id, title, thumbneil from games;",
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var thumbneil string

		if err := rows.Scan(&id, &title, &thumbneil); err != nil {
			continue
		}
		game := domain.Game{
			ID:            id,
			Title:         title,
			ThumbneilPath: thumbneil,
		}

		games = append(games, game)
	}

	return
}
