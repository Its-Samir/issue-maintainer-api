package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	var err error

	/* database url */
	connString := "....."

	/* create a new pool */
	DB, err = pgxpool.New(context.Background(), connString)

	if err != nil {
		panic("Could not connect to DB")
	}

	/* the orders are matter, because of foreign key relations */
	err = createStatusType()

	if err != nil {
		panic("Could not create status type")
	}

	err = createUserTable()

	if err != nil {
		panic("Could not create user table")
	}

	err = createGroupTable()

	if err != nil {
		panic("Could not create group table")
	}

	err = createIssueTable()

	if err != nil {
		panic("Could not create issue table")
	}
}

func createStatusType() error {
	query := `
		DO $$ BEGIN
			CREATE TYPE issue_status AS ENUM ('open', 'in_progress', 'closed');
		EXCEPTION
			WHEN duplicate_object THEN NULL;
		END $$;
	`
	_, err := DB.Exec(context.Background(), query)

	return err
}

func createUserTable() error {
	query := `
		-- can join only one group for now, we can change that also, but i decided to choose single group for here

		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			email VARCHAR(255) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			group_id INTEGER
		)
	`
	_, err := DB.Exec(context.Background(), query)

	return err
}

func createGroupTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS groups(
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL UNIQUE,
			description TEXT NOT NULL,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
		)
	`
	_, err := DB.Exec(context.Background(), query)

	return err
}

func createIssueTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS issues(
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			status issue_status DEFAULT 'open'::issue_status,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			assigned_to_user_id INTEGER NOT NULL,
			group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE
		)
	`
	_, err := DB.Exec(context.Background(), query)

	return err
}
