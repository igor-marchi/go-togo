# API de Tarefas em Go

Uma API simples de tarefas (Todo) construída com Go, SQLite e Docker. Este projeto fornece endpoints para gerenciar uma lista de tarefas, incluindo a criação e a recuperação de tarefas.

## Funcionalidades

- Criar uma nova tarefa
- Recuperar todas as tarefas
- Armazenamento persistente usando SQLite
- Dockerizado para fácil implantação

## Requisitos

- [Docker](https://www.docker.com/) (para implantação em contêiner)

## Configuração

### Usando Docker Compose

1. Inicie os serviços:

   ```bash
   docker-compose up
   ```

2. Acesse a API em `http://localhost:8080`.

## Endpoints da API

### Recuperar Tarefas

- **URL:** `/todos`
- **Método:** `GET`
- **Resposta:**
  ```json
  [
    {
      "id": 1,
      "title": "Tarefa Exemplo",
      "done": false
    }
  ]
  ```

### Criar Tarefa

- **URL:** `/todos`
- **Método:** `POST`
- **Corpo da Requisição:**
  ```json
  {
    "title": "Nova Tarefa"
  }
  ```
- **Resposta:**
  ```json
  {
    "id": 2,
    "title": "Nova Tarefa",
    "done": false
  }
  ```

## Estrutura do Projeto

- `cmd/api`: Ponto de entrada da aplicação.
- `internal/todo`: Contém a lógica de negócios, modelos e handlers da API de tarefas.
- `internal/db`: Lógica de conexão com o banco de dados.
- `internal/config`: Lógica de carregamento de configurações.

## Licença

Este projeto está licenciado sob a Licença MIT.
