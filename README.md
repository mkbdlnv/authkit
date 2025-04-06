# authkit
This is universal go module for implementing authorization for go projects

go-authkit/
├── go.mod
├── go.sum
├── README.md
├── LICENSE
├── .gitignore
│
├── pkg/                      # Публичный API
│   └── authkit.go            # Точка входа, экспорт интерфейса Authenticator
│
├── internal/                 # Внутренние реализации
│   ├── auth/                 # Основная бизнес-логика (реализация Authenticator)
│   │   ├── service.go
│   │   ├── interface.go      # Интерфейс Authenticator
│   │   ├── models.go         # RegisterInput, LoginInput, User, TokenPair и т.п.
│   │   └── errors.go         # Предопределённые ошибки
│   │
│   ├── token/                # Работа с access/refresh токенами
│   │   ├── interface.go      # TokenManager interface
│   │   ├── jwt.go            # Реализация на JWT (HS256)
│   │   └── jwt_rs256.go      # (опционально) Реализация на RS256
│   │
│   ├── hasher/               # Хэширование паролей
│   │   ├── interface.go      # Hasher interface
│   │   └── bcrypt.go         # Реализация через bcrypt
│   │
│   ├── userstore/            # Абстракция хранилища пользователей
│   │   ├── interface.go      # UserStore interface
│   │   ├── inmemory.go       # Простая реализация in-memory
│   │   └── postgres.go       # (позже) Реализация на PostgreSQL
│   │
│   └── logger/               # Логгер
│       ├── interface.go      # Logger interface
│       └── stdlog.go         # Простая реализация через стандартный логгер
│
├── examples/                 # Примеры интеграции
│   └── gin-server/           # Пример REST API с Gin
│       ├── main.go
│       └── handler.go
│
├── test/                     # Тесты
│   ├── auth_test.go
│   ├── token_test.go
│   └── integration_test.go
│
└── mocks/                    # Моки для тестов
    ├── userstore_mock.go
    ├── token_mock.go
    └── hasher_mock.go

## 📌 Итого: порядок модулей
Очередь	Модуль	Что делает
1	hasher	Хэширование паролей (bcrypt)
2	token	JWT или RS256
3	userstore	In-memory storage
4	logger	Интерфейс + простая реализация
5	auth	Бизнес-логика, интерфейс API
6	pkg/	Точка входа в модуль
7	examples/	Примеры интеграции
8	test/	Тесты, моки