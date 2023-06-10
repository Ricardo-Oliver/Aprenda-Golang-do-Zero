package main

import (
	"log"
	"net/http"
)

func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários"))
}

func main() {
	//HTTP é um protocolo de comunicação - Base da comunicação Web
	//Cliente (Faz requisição) - Servidor (Processa a requisição e envia a resposta)
	//Request - Response
	//Rotas (endpoints)
	//URI - Identificador do Recurso
	//Método - GET, POST, PUT, DELETE...

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))

}