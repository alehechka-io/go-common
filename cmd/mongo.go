package cmd

import (
	"github.com/urfave/cli/v2"
)

// MongoConnectionStringArg is the CLI argument for the MongoDB connection string
const MongoConnectionStringArg string = "mongo-connection-string"

// MongoConnectionStringFlag is the urfave/cli Flag configuration for the MongoDB connection string
var MongoConnectionStringFlag = &cli.StringFlag{
	Name:     MongoConnectionStringArg,
	Usage:    "Specifies the MongoDB connection string.",
	EnvVars:  []string{"MONGO_CONNECTION_STRING"},
	Required: true,
}

// MongoDatabaseArg is the CLI argument for the MongoDB database
const MongoDatabaseArg string = "mongo-database"

// MongoDatabaseFlag is the urfave/cli Flag configuration for the MongoDB database
var MongoDatabaseFlag = &cli.StringFlag{
	Name:     MongoDatabaseArg,
	Usage:    "Specifies the MongoDB database",
	EnvVars:  []string{"MONGO_DATABASE"},
	Required: true,
}
