package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo"))
}

func usuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregando usuários!"))
}

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICACAO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// REQUEST - RESPONSE

	// ROTAS
		// URI - INDENTIFICADOR DO RECURSO
		// MÉTODO - GET, POST, PUT, DELETE
		// GET - BUSCA
		// POST - CADASTRAR
		// PUT ATUALIZAR DADOS
		// DELETE - DELETAR DADOS

		http.HandleFunc("/home", home)
		http.HandleFunc("/usuario", usuario)

		log.Fatal(http.ListenAndServe(":8000", nil))
}