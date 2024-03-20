package psqlstore

type Migration struct {
	Up   string
	Down string
}

var migrations = []Migration{
	{
		Up: `CREATE TABLE IF NOT EXISTS song (
				id SERIAL PRIMARY KEY, 
				title VARCHAR NOT NULL, 
				author VARCHAR NOT NULL, 
				album VARCHAR NOT NULL, 
				genre VARCHAR NOT NULL, 
				created_at TIMESTAMP NOT NULL 
			);`,
		Down: `ALTER TABLE song DROP COLUMN IF EXISTS created_at;
			`,
	},
}
