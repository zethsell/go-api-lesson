package repositories

import (
	"api/src/models"
	"database/sql"
)

type Post struct {
	db *sql.DB
}

func PostRepository(db *sql.DB) *Post {
	return &Post{db}
}

func (repository Post) List(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
	SELECT DISTINCT p.*, u.nick FROM posts p
	INNER JOIN users u on u.id = p.authorId
	INNER JOIN followers f on p.authorId = f.userId
	WHERE u.id = ? OR f.followerId = ?`,
		userID, userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Post) Store(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO posts (title,content,authorId) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorId)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil

}

func (repository Post) Show(ID uint64) (models.Post, error) {
	row, err := repository.db.Query(`
		SELECT p.*, u.nick FROM
		posts p INNER JOIN users u 
		on u.id = p.authorId 
		 WHERE p.id = ? 
	`, ID)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post

	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository Post) Update(ID uint64, post models.Post) error {
	statement, err := repository.db.Prepare("UPDATE posts SET title= ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		post.Title,
		post.Content,
		ID,
	); err != nil {
		return err
	}

	return nil
}

func (repository Post) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM posts  WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository Post) ListByAuthor(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
	SELECT  p.*, u.nick FROM posts p
	JOIN users u on p.authorId = u.id
	WHERE p.authorId = ? `, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Post) Like(ID uint64) error {
	statement, err := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository Post) UnLike(ID uint64) error {
	statement, err := repository.db.Prepare("UPDATE posts SET likes = IF(likes = 0, 0, likes -1) WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
