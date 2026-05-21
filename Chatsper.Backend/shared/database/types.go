package database

const (
	SqliteDataBase   DatabaseType = "SQLITE"   // 0
	PostgresDatabase DatabaseType = "POSTGRES" // 1
)

type DatabaseType string
