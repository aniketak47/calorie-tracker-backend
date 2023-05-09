package models

var migrationModels = []interface{} {
	&Entry{},
}

func GetMigrationModel() []interface{} {
	return migrationModels
}