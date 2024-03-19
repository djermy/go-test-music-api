package database

var migrationQuery []string = []string{
	`
		CREATE TABLE IF NOT EXISTS song (
			id SERIAL PRIMARY KEY, 
			title VARCHAR NOT NULL, 
			author VARCHAR NOT NULL, 
			album VARCHAR NOT NULL, 
			genre VARCHAR NOT NULL, 
			created_at TIMESTAMP NOT NULL 
		);
	`,
	`
		ALTER TABLE song DROP COLUMN IF EXISTS created_at;
	`,
}
