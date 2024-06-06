package models

import (
	"errors"
	"gshort/db"
	"time"
)

type Url struct {
	ID        int64
	LongUrl   string `form:"url"`
	ShortCode string
	DateTime  time.Time
	UserID    int64
}

func (u *Url) Save() error {
	query := `
	INSERT INTO urls(long_url, short_code, date_time, user_id)
	VALUES (?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.LongUrl, u.ShortCode, u.DateTime, u.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id

	return nil
}

func FindAllUrlsByUserId(userId int64) ([]Url, error) {
	query := `SELECT * FROM urls WHERE user_id = ? ORDER BY date_time DESC`
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []Url
	for rows.Next() {
		var url Url

		err := rows.Scan(&url.ID, &url.LongUrl, &url.ShortCode, &url.DateTime, &url.UserID)
		if err != nil {
			return nil, err
		}

		urls = append(urls, url)
	}

	return urls, nil
}

func FindUrlByShortCode(shortCode string) (Url, error) {
	query := `SELECT * FROM urls WHERE short_code=?`
	rows, err := db.DB.Query(query, shortCode)

	if err != nil {
		return Url{}, err
	}

	if !rows.Next() {
		return Url{}, errors.New("not found")
	}

	var url Url
	err = rows.Scan(&url.ID, &url.LongUrl, &url.ShortCode, &url.DateTime, &url.UserID)
	if err != nil {
		return Url{}, err
	}

	defer rows.Close()

	return url, nil
}
