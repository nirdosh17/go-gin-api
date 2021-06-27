# Go Rest API sample using Gin framework

## Features:
- CRUD APIs
- List API with pagination
- Postgres for storage
- Error tracking in [Sentry](https://sentry.io/)
- API key header(`x-api-key`) based authorization
- Containerized | Approx image size: 25 MB

## APIs
### 1. Create a Product
- Endpoint: `POST /products`
- Request body:
  ```json
  {
  	"name": "Sony ANC Headphone",
  	"price": 30000,
  	"category": "electronics"
  }
  ```
- Response: `Product`
  ```json
  {
    "ID": 1,
    "Category": "electronics",
    "Name": "Sony ANC Headphone",
    "Price": 30000,
    "CreatedAt": "2021-06-27T09:04:51Z",
    "UpdatedAt": "2021-06-27T09:04:51Z",
    "DeletedAt": null
  }
  ```

### 2. Get a Product
- Endpoint: `GET /products/:id`
- Response: Product

### 3. Update Price of a Product
- Endpoint: `PUT /products/:id`
- Request body:
  ```json
  {
  	"price": 27000
  }
  ```
- Response: `Product`

### 4. Delete a Product
- Endpoint: `DELETE: /products/:id`
- Response: 204 | No Content

### 5. List All Products
- Endpoint: `GET: /products`
- Query params:
  - `page`
  - `per_page`: default 10 | max 100
- Response: Array of `Product`

## References:
Gin Framework:
- [QuickStart](https://gin-gonic.com/docs/quickstart/)

Sentry:
- [Sentry setup for Gin](https://docs.sentry.io/platforms/go/guides/gin)

GORM:
- [Queries](https://gorm.io/docs/query.html)
- [Migration](https://gorm.io/docs/migration.html)
- [Scopes for commonly used logic](https://gorm.io/docs/scopes.html)
