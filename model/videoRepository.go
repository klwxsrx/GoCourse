package model

import (
	"database/sql"
	"gocourse/database"
)

type VideoRepository struct {
	db *sql.DB
}

func (r *VideoRepository) GetByKey(key string) (*Video, error) {
	row := r.db.QueryRow(`
		SELECT
			*
		FROM
			video
		WHERE
			video_key = ?
			AND status = ?;`, key, StatusReady)

	var video = Video{}
	err := row.Scan(&video.id, &video.Key, &video.Name, &video.Status, &video.Duration, &video.Url, &video.ThumbnailUrl)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &video, nil
}

func (r *VideoRepository) GetAll() ([]*Video, error) {
	rows, err := r.db.Query(`
		SELECT
			*
		FROM
			video
		WHERE
		    status = ?;
	`, StatusReady)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []*Video
	for rows.Next() {
		var video = Video{}
		err := rows.Scan(&video.id, &video.Key, &video.Name, &video.Status, &video.Duration, &video.Url, &video.ThumbnailUrl)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}
	return videos, nil
}

func (r *VideoRepository) Save(video *Video) error {
	if video.id == 0 {
		return r.create(video)
	} else {
		return r.update(video)
	}
}

func (r *VideoRepository) create(video *Video) error {
	result, err := r.db.Exec(`
		INSERT INTO
			video
		SET
			id = ?,
			video_key = ?,
			title = ?,
			status = ?,
			duration = ?,
			url = ?,
			thumbnail_url = ?;	
	`, video.id, video.Key, video.Name, video.Status, video.Duration, video.Url, video.ThumbnailUrl)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	video.id = int(id)
	return nil
}

func (r *VideoRepository) update(video *Video) error {
	_, err := r.db.Exec(`
		UPDATE
			video
		SET 
			video_key = ?,
			title = ?,
			status = ?,
			duration = ?,
			url = ?,
			thumbnail_url = ?
		WHERE
			id = ?	
	`, video.Key, video.Name, video.Status, video.Duration, video.Url, video.ThumbnailUrl, video.id)

	return err
}

func NewVideoRepository() *VideoRepository {
	return &VideoRepository{database.GetConnection()}
}
