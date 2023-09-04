package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func UserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name,nick,email,password) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository Users) List(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repository.db.Query(
		"SELECT id, name,nick,email, createdAt FROM users WHERE name LIKE ? or nick LIKE ?", nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil
}

func (repository Users) Show(ID uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT id, name, nick, email, createdAt FROM users WHERE id = ?", ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) FindByEmail(email string) (models.User, error) {
	row, err := repository.db.Query("SELECT id,password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare("UPDATE users SET name = ?,nick = ?,email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository Users) Follow(userID uint64, followerID uint64) error {
	statement, err := repository.db.Prepare("INSERT IGNORE INTO followers (userId, followerId) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository Users) Unfollow(userID uint64, followerID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM followers WHERE userId = ? AND followerId =?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository Users) Followers(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT u.id, u.name,u.nick, u.email, u.createdAt
				FROM users u  
				INNER JOIN  followers on followers.followerId = u.id
				WHERE followers.userId = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var follower models.User

		if err = rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository Users) Following(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT u.id, u.name,u.nick, u.email, u.createdAt
				FROM users u  
				INNER JOIN  followers on followers.userId = u.id
				WHERE followers.followerId = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var follower models.User

		if err = rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository Users) GetCurrentPassword(ID uint64) (string, error) {
	row, err := repository.db.Query(
		"SELECT password FROM users WHERE id = ?", ID,
	)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(
			&user.Password,
		); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository Users) UpdatePassword(ID uint64, password string) error {
	statement, err := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(password, ID); err != nil {
		return err
	}

	return nil
}
