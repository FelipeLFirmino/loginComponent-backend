# Login Component - Backend

Este é o backend para o sistema de login e cadastro, desenvolvido em Go (Golang). Ele utiliza o `httprouter` para roteamento e manipula requisições HTTP para autenticação e registro de usuários.

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem principal para o backend.
- **httprouter**: Gerenciamento de rotas HTTP.
- **JSON**: Utilizado para armazenar e transferir os dados dos usuários.

## Funcionalidades

1. **Login**: Verifica se o nome de usuário e senha correspondem a um usuário registrado.
2. **Cadastro**: Permite o registro de novos usuários.
3. **Armazenamento de Dados**: Os dados dos usuários são armazenados em um arquivo JSON.

## Estrutura do Projeto

```plaintext
backend/
│
├── handlers/
│   └── signup_handler.go    # Handler para cadastro de usuários
├── models/
│   └── user.go              # Estrutura do modelo de usuário
├── utils/
│   ├── load_users.go        # Função para carregar usuários do arquivo JSON
│   ├── save_user.go         # Função para salvar usuários no arquivo JSON
│   └── check_and_save.go    # Verifica e salva novos usuários
├── data/
│   └── users.json           # Arquivo JSON com os dados dos usuários
├── main.go                  # Arquivo principal para inicializar o servidor
└── README.md                # Documentação do backend
