# 🚀 Monitor-Go — Sistema Backend de Monitoramento em Go

Aplicação backend de alta performance desenvolvida em **Go (Golang)**, estruturada com padrões de arquitetura limpa em camadas (Clean/Layered Architecture), separação de responsabilidades e foco em testes automatizados.

---

## 🛠️ Tecnologias e Stack
* **Linguagem:** Go (Golang)
* **Arquitetura:** Padrão em camadas (Config, Database, Models, Repository, Service)
* **Testes:** Testes unitários integrados (`net/http/httptest` / testing nativo do Go)
* **Gerenciamento de Dependências:** Go Modules (`go.mod` / `go.sum`)

---

## 📂 Arquitetura do Projeto
O projeto segue uma divisão modular limpa para facilitar a manutenção, escalabilidade e testabilidade do código:


monitor-go/
├── config/        # Configurações gerais da aplicação e variáveis de ambiente
├── database/      # Conexão, inicialização e gerenciamento do banco de dados
├── models/        # Definição das estruturas de dados (Entities/Structs)
├── repository/    # Camada de persistência e acesso a dados (Query/Storage)
├── service/       # Camada de regras de negócio e lógica da aplicação
├── main.go        # Ponto de entrada da aplicação (Entrypoint / Router)
├── main_test.go   # Testes unitários e de integração do sistema
└── go.mod         # Gerenciamento de módulos e dependências Go
⚙️ Como Executar o Projeto
Certifique-se de ter o Go (Golang) instalado na sua máquina (versão 1.18+ recomendada).

Clone este repositório:

Bash
git clone [https://github.com/Davilacovelli/monitor-go.git](https://github.com/Davilacovelli/monitor-go.git)
cd monitor-go
Baixe as dependências do projeto:

Bash
go mod tidy
Execute a aplicação:

Bash
go run main.go
🧪 Executando os Testes Automatizados
O projeto conta com testes unitários implementados no arquivo main_test.go. Para rodar a suíte de testes com relatório de cobertura, utilize:

Bash
go test -v -cover ./...
👤 Autor
Desenvolvido por Davi Iacovelli

Estudante de Ciência da Computação — Universidade Veiga de Almeida (UVA)
