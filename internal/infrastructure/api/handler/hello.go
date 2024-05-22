package handler

// Importação de pacotes necessários
import (
    "fmt" // Pacote para formatação de strings e impressão
    "net/http" // Pacote para manipulação de requisições HTTP
)

// Definição de um tipo struct chamado HelloHandler
// Este struct não contém campos, mas é utilizado para agrupar métodos relacionados.
type HelloHandler struct {

}

// Definição de um método chamado Hello associado ao tipo HelloHandler
// Este método manipula requisições HTTP e tem dois parâmetros:
// - w: um http.ResponseWriter, que é utilizado para escrever a resposta HTTP
// - r: um *http.Request, que representa a requisição HTTP recebida
func (h *HelloHandler) Hello(w http.ResponseWriter, r *http.Request) {
    // Verificação se o método HTTP da requisição é GET
    if r.Method == http.MethodGet {
        // Se o método for GET, escreve "Hello, World!" na resposta HTTP
        fmt.Fprintf(w, "Hello, World!")
    } else {
        // Se o método não for GET, retorna um erro 405 Method Not Allowed
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
