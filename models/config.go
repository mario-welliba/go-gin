package models

import (
	"gin/db"
	"os"
)

var server = os.Getenv("DATABASE")

// Database name
var databaseName = os.Getenv("DATABASE_NAME")

// Create a connection
var dbConnect = db.NewConnection(server)
