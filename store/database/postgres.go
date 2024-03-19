package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func InitDB() (*pgx.Conn, error) {
	conn, err := dbConnect()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = migrate(conn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

func dbConnect() (*pgx.Conn, error) {
	url := "postgres://postgres:password@localhost:5432/music"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func migrate(conn *pgx.Conn) error {
	for _, query := range migrationQuery {
		_, err := conn.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
	commandTag, err := conn.Exec(context.Background(), "delete from widgets where id=$1", 42)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("No row found to delete")
	}

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
*/
