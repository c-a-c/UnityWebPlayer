package database

import (
	"UnityWebPlayer/domain"
	"time"
)

type GameRepository struct {
	SqlHandler
}

func (repo *GameRepository) FindById(identifier int) (game domain.Game, err error) {
	row, err := repo.Query(
		"select id, title, auther, created_at from games where id=?;", identifier,
	)
	if err != nil {
		return
	}
	defer row.Close()

	var id int
	var title string
	var auther string
	var createdAt *time.Time
	var deletedAt *time.Time

	row.Next()
	if err = row.Scan(&id, &title, &auther, &createdAt, &deletedAt); err != nil {
		return
	}

	game.ID = id
	game.Title = title
	game.Auther = auther
	game.CreatedAt = createdAt
	game.DeletedAt = deletedAt
	return
}

func (repo *GameRepository) FindAll() (games domain.Games, err error) {
	rows, err := repo.Query(
		"select id, title, auther, created_at from games;",
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var auther string
		var createdAt *time.Time
		var deletedAt *time.Time

		if err := rows.Scan(&id, &title, &auther, &createdAt, &deletedAt); err != nil {
			continue
		}
		game := domain.Game{
			ID:        id,
			Title:     title,
			Auther:    auther,
			CreatedAt: createdAt,
			DeletedAt: deletedAt,
		}

		games = append(games, game)
	}

	return
}
