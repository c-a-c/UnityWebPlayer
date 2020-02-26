package database

import (
	"UnityWebPlayer/domain"
)

type GameRepository struct {
	SqlHandler
}

func (repo *GameRepository) FindById(identifier int) (game domain.Game, err error) {
	row, err := repo.Query(
		"select id, title from games where id=?;", identifier,
	)
	if err != nil {
		return
	}
	defer row.Close()

	var id int
	var title string
	row.Next()
	if err = row.Scan(&id, &title); err != nil {
		return
	}

	game.ID = id
	game.Title = title
	return
}

func (repo *GameRepository) FindAll() (games domain.Games, err error) {
	rows, err := repo.Query(
		"select id, title from games;",
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string

		if err := rows.Scan(&id, &title); err != nil {
			continue
		}
		game := domain.Game{
			ID:    id,
			Title: title,
		}

		games = append(games, game)
	}

	return
}
