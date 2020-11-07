package banco

import (

	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"

)

func Conectar () (*sql.DB , error ) {

	stringConexao := "root:TeReVe!3@@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {

		log.Fatal(erro)
		return nil, erro

	}

	if erro = db.Ping(); erro != nil {

		return nil, erro

	}

	return db, nil 
}