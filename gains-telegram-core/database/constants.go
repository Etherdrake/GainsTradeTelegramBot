package database

import "os"

const (
	databaseName   = "hooterdb"
	collectionName = "users"
)

var (
	mongoURI = os.Getenv("MONGO_URI")
)
