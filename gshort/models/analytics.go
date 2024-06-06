package models

import (
	"gshort/db"
	"time"
)

type Analytics struct {
	ID        int64
	Views     int64
	DateTime  time.Time
	UserAgent string
	IP        string
	UrlID     int64
}

func (a *Analytics) Save() error {
	id, err := db.PrepareAndExec(`
		INSERT INTO analytics (views, date_time, user_agent, ip, url_id)
		VALUES(?,?,?,?,?)
	`, a.Views, a.DateTime, a.UserAgent, a.IP, a.UrlID)
	if err != nil {
		return err
	}

	a.ID = id

	return nil
}

func FindAllViewsOfUrl(urlId string) ([]Analytics, error) {
	query := `SELECT * FROM analytics WHERE url_id=?`
	rows, err := db.DB.Query(query, urlId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var analytics []Analytics

	for rows.Next() {
		var url Analytics

		err := rows.Scan(&url.ID, &url.Views, &url.DateTime, &url.UserAgent, &url.IP, &url.UrlID)
		if err != nil {
			return nil, err
		}

		analytics = append(analytics, url)
	}

	return analytics, nil
}
