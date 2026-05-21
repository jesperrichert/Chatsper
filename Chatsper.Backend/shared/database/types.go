package database

const (
	SqliteDataBase   DatabaseType = iota // 0
	PostgresDatabase                     // 1
)

type DatabaseType int
