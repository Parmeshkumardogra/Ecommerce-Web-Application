package mysqlservices

import (
	"database/sql"
	"log"

	"github.com/BMS/database"
)

func SelectOne(query string, values ...interface{}) (*sql.Row){
	log.Printf("Executing query : %s with values : %v ",query,values);
	sqlRes := database.DB.QueryRow(query,values...);
	return sqlRes;
}