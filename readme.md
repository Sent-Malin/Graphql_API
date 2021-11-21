# GraphQL API

```bash
Для начала работы следует ввести команды:
docker-compose up -d
# команда накатывает миграции и тестовые данные
make migrateup
go run server.go
```

Примеры запросов
Получить список продуктов

```graphql
query {
  products {
    id
    name
  }
}
```

Отправить код в терминале. Запрос вернет null (все ок) или ошибку.

```graphql
mutation {
  requestSignInCode(input: { phone: "799999999" }) {
    message
  }
}
```

Авторизация с номер+код, результатом является токен или ошибка

```graphql
mutation {
  signInByCode(input: { phone: "799999999", code: "0000" }) {
    ... on SignInPayload {
      token
      viewer {
        user {
          phone
        }
      }
    }
    ... on ErrorPayload {
      message
    }
  }
}
```

С токеном можно получить данные пользователя. Токен передается через Authorization заголовок, для этого следует прописать в HTTP HEADERS:

```graphql
{
    "authorization": "(токен)"
}
```

```graphql
query {
  viewer {
    user {
      phone
    }
  }
}
```
