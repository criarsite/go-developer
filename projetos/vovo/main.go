package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

// BancoMock simula um banco de dados em memória, somente para consultas.
type BancoMock struct {
	clientes []Cliente
}

func (b BancoMock) Listar() []Cliente {
	return append([]Cliente(nil), b.clientes...)
}

func (b BancoMock) BuscarPorID(id int) (Cliente, bool) {
	for _, cliente := range b.clientes {
		if cliente.ID == id {
			return cliente, true
		}
	}

	return Cliente{}, false
}

var banco = BancoMock{
	clientes: []Cliente{
		{ID: 1, Nome: "Ana Souza", Empresa: "Acme Tecnologia", CNPJ: "00.000.000/0001-01", Email: "ana@example.com", Preco: 1200},
		{ID: 2, Nome: "Bruno Lima", Empresa: "Beta Comercio", CNPJ: "00.000.000/0001-02", Email: "bruno@example.com", Preco: 850.50},
		{ID: 3, Nome: "Carla Mendes", Empresa: "Gama Servicos", CNPJ: "00.000.000/0001-03", Email: "carla@example.com", Preco: 2400},
	},
}

// Responde com JSON
func responderJSON(w http.ResponseWriter, status int, dados any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dados)
}

// GET /clientes
// Lista todos os clientes
func listarClientes(w http.ResponseWriter, r *http.Request) {
	responderJSON(w, http.StatusOK, banco.Listar())
}

var paginaClientes = template.Must(template.New("clientes").Parse(`<!doctype html>
<html lang="pt-BR">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Clientes de exemplo</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 2rem auto; max-width: 1000px; padding: 0 1rem; color: #222; }
		h1 { margin-bottom: 0.5rem; }
		.note { color: #555; margin-bottom: 1.5rem; }
		table { border-collapse: collapse; width: 100%; }
		th, td { border: 1px solid #ccc; padding: 0.65rem; text-align: left; }
		th { background: #f2f2f2; }
		tbody tr:nth-child(even) { background: #fafafa; }
		code { background: #f2f2f2; padding: 0.15rem 0.3rem; }
	</style>
</head>
<body>
	<h1>Clientes de exemplo</h1>
	<p class="note">Esta tabela mostra os dados do banco mock em memória. Eles são reiniciados quando a aplicação para.</p>
	<p>Para consultar os dados em JSON, acesse <a href="/clientes"><code>GET /clientes</code></a>. Para ver um cliente específico, use <code>GET /clientes/{id}</code>, por exemplo <a href="/clientes/1"><code>/clientes/1</code></a>.</p>
	<table>
		<thead>
			<tr><th>ID</th><th>Nome</th><th>Empresa</th><th>CNPJ</th><th>E-mail</th><th>Preço</th></tr>
		</thead>
		<tbody>
			{{range .}}
			<tr>
				<td>{{.ID}}</td><td>{{.Nome}}</td><td>{{.Empresa}}</td><td>{{.CNPJ}}</td><td>{{.Email}}</td><td>{{printf "%.2f" .Preco}}</td>
			</tr>
			{{else}}
			<tr><td colspan="6">Nenhum cliente cadastrado.</td></tr>
			{{end}}
		</tbody>
	</table>
</body>
</html>`))

func mostrarClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := paginaClientes.Execute(w, banco.Listar()); err != nil {
		log.Printf("erro ao renderizar a página de clientes: %v", err)
	}
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

	cliente, existe := banco.BuscarPorID(id)
	if !existe {
		responderJSON(w, http.StatusNotFound,
			map[string]string{"erro": "Cliente não encontrado"})
		return
	}

	responderJSON(w, http.StatusOK, cliente)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", mostrarClientes).Methods(http.MethodGet)

	router.HandleFunc("/clientes", listarClientes).
		Methods(http.MethodGet)

	router.HandleFunc("/clientes/{id}", buscarCliente).
		Methods(http.MethodGet)

	log.Println("API rodando em http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
