package dbstorefactory

import (
	"github.com/nikhil133/duplo/src/dbstore"
	"github.com/nikhil133/duplo/src/dbstoreinterface"
)

func NewMySql() dbstoreinterface.DbStoreInterface {
	return &dbstore.MySql{}
}
