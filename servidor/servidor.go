package servidor


import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"crud/banco"
)


type usuario struct {

	ID 		uint32 	`json:"id"`
	Nome 	string 	`json:"nome"`
	Email 	string 	`json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request)  {

	corpoRequisicao , erro := ioutil.ReadAll( r.Body )

	if erro != nil {

		w.Write( []byte("Falha ao leer o corpo da requisição") )
		return 
	}


	var usuario usuario 
	if erro = json.Unmarshal( corpoRequisicao, &usuario ); erro != nil {

		w.Write( []byte("Erro ao converter o usuário para struct") )
		return
	}


	//fmt.Println(usuario)

	db, erro := banco.Conectar() 
	if erro != nil {
		w.Write( []byte("Erro ao Conectar no banco de dados") )
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome,email) values (?,?) ")
	if erro != nil {

		w.Write( []byte("Erro ao criar statement!") )
		return
	}
	defer statement.Close() 



	insercao, erro := statement.Exec( usuario.Nome, usuario.Email )
	if erro != nil {

		w.Write( []byte("Erro ao executar o statement!") )
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write( []byte("Erro ao obter o ID inserido") )		
		return 
	}


	w.WriteHeader(http.StatusCreated)
	w.Write( []byte( fmt.Sprintf( "Usuário inserido com sucesso! Id: %d", idInserido) ) )



}
