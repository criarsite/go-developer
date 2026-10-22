package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Estrutura do cliente
type Cliente struct {
	ID      int     `json:"id"`
	Nome    string  `json:"nome"`
	Empresa string  `json:"empresa"`
	CNPJ    string  `json:"cnpj"`
	Email   string  `json:"email"`
	Preco   float64 `json:"preco"`
}

// Armazenamento em memória
var (
	clientes  = make(map[int]Cliente)
	proximoID = 1
	mu        sync.RWMutex
)

// Responde com JSON
func responderJSON(w http.ResponseWriter, status int, dados any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dados)
}

// Valida os dados recebidos
func validarCliente(c Cliente) bool {
	return c.Nome != "" &&
		c.Empresa != "" &&
		c.CNPJ != "" &&
		c.Preco >= 0
}

// GET /clientes
// Lista todos os clientes
func listarClientes(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	lista := make([]Cliente, 0, len(clientes))

	for _, cliente := range clientes {
		lista = append(lista, cliente)
	}

	mu.RUnlock()

	responderJSON(w, http.StatusOK, lista)
}

// GET /clientes/{id}
// Busca um cliente pelo ID
func buscarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		responderJSON(w, http.StatusBadRequest,
			map[string]string{"erro": "ID inválido"})
		return
	}

	mu.RLock()
	cliente, existe := clientes[id]
	mu.RUnlock()

	if !existe {
		responderJSON(w, http.StatusNotFound,
			map[string]string{"erro": "Cliente não encontrado"})
		return
	}

	responderJSON(w, http.StatusOK, cliente)
}

// POST /clientes
// Cadastra um novo cliente
func criarCliente(w http.ResponseWriter, r *http.Request) {
	var cliente Cliente

	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		responderJSON(w, http.StatusBadRequest,
			map[string]string{"erro": "JSON inválido"})
		return
	}

	if !validarCliente(cliente) {
		responderJSON(w, http.StatusBadRequest,
			map[string]string{
				"erro": "Nome, empresa, CNPJ e preço válido são obrigatórios",
			})
		return
	}

	mu.Lock()
	cliente.ID = proximoID
	proximoID++
	clientes[cliente.ID] = cliente
	mu.Unlock()

	responderJSON(w, http.StatusCreated, cliente)
}

// PUT /clientes/{id}
// Atualiza um cliente existente
func atualizarCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		responderJSON(w, http.StatusBadRequest,
			map[string]string{"erro": "ID inválido"})
		return
	}

	var cliente Cliente

	if err := json.NewDecoder(r.Body).Decode(&cliente); err != nil {
		responderJSON(w, http.StatusBadRequest,
			map[string]string{"erro": "JSON inválido"})
		return
	}

	if !validarCliente(cliente) {
		responderJSON(w, http.StatusBadRequest,
			map[string]string{
				"erro": "Nome, empresa, CNPJ e preço válido são obrigatórios",
			})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, existe := clientes[id]; !existe {
		responderJSON(w, http.StatusNotFound,
			map[string]string{"erro": "Cliente não encontrado"})
		return
	}

	cliente.ID = id
	clientes[id] = cliente

	responderJSON(w, http.StatusOK, cliente)
}

// DELETE /clientes/{id}
// Exclui um cliente
func excluirCliente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		responderJSON(w, http.StatusBadRequest,
			map[string]string{"erro": "ID inválido"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, existe := clientes[id]; !existe {
		responderJSON(w, http.StatusNotFound,
			map[string]string{"erro": "Cliente não encontrado"})
		return
	}

	delete(clientes, id)

	responderJSON(w, http.StatusOK,
		map[string]string{"mensagem": "Cliente excluído com sucesso"})
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/clientes", listarClientes).
		Methods(http.MethodGet)

	router.HandleFunc("/clientes/{id}", buscarCliente).
		Methods(http.MethodGet)

	router.HandleFunc("/clientes", criarCliente).
		Methods(http.MethodPost)

	router.HandleFunc("/clientes/{id}", atualizarCliente).
		Methods(http.MethodPut)

	router.HandleFunc("/clientes/{id}", excluirCliente).
		Methods(http.MethodDelete)

	log.Println("API rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
