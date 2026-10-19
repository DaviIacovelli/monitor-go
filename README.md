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

```text
monitor-go/
├── config/        # Configurações gerais da aplicação e variáveis de ambiente
├── database/      # Conexão, inicialização e gerenciamento do banco de dados
├── models/        # Definição das estruturas de dados (Entities/Structs)
├── repository/    # Camada de persistência e acesso a dados (Query/Storage)
├── service/       # Camada de regras de negócio e lógica da aplicação
├── main.go        # Ponto de entrada da aplicação (Entrypoint / Router)
├── main_test.go   # Testes unitários e de integração do sistema
└── go.mod         # Gerenciamento de módulos e dependências Go
