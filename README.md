# Auction API

Este projeto implementa uma API para gerenciamento de **leilões**, incluindo funcionalidades de criação, listagem, lances e consulta do vencedor.  

---

## 🚀 Build and Run with Docker

Para construir e rodar a aplicação:

```bash
docker-compose up --build
```

---

## Testar fechamento automatico de leilão

O sistema está configurado para fechar leilões automaticamente após um intervalo definido (10 minutos por padrão). Para testar essa funcionalidade, você pode ajustar o valor da variável de ambiente `AUCTION_INTERVAL` no arquivo `.env` para um tempo menor, como `1m` (1 minuto), e reiniciar os containers Docker.

## 📌 API Endpoints

### 🔹 Auctions

- **GET** `/auction` - Lista todos os leilões com filtros opcionais  

  **Query parameters:**
  - `status` → Filtrar por status do leilão (`0 = Ativo`, `1 = Concluído`) (required)
  - `category` → Filtrar por categoria(opcional)
  - `productName` → Filtrar por nome do produto(opcional)  

- **GET** `/auction/:auctionId` - Detalhes de um leilão específico  

- **POST** `/auction` - Criar um novo leilão  

  **Request body:**
  ```json
  {
    "product_name": "Example Product",
    "category": "Electronics",
    "description": "This is a detailed description of the product",
    "condition": 1
  }
  ```

  **Condition values:**
  - `0`: New  
  - `1`: Used  
  - `2`: Refurbished  

- **GET** `/auction/winner/:auctionId` - Obter o lance vencedor de um leilão  

---

### 🔹 Bids

- **POST** `/bid` - Fazer um lance  

  **Request body:**
  ```json
  {
    "user_id": "user-uuid",
    "auction_id": "auction-uuid",
    "amount": 100.50
  }
  ```

- **GET** `/bid/:auctionId` - Listar todos os lances de um leilão  

---

### 🔹 Users

- **GET** `/user/:userId` - Obter detalhes de um usuário  

---

## 🧪 Testing the System

### 1. Criar um leilão
```bash
curl -X POST http://localhost:8080/auction   -H "Content-Type: application/json"   -d '{
    "product_name": "Vintage Camera",
    "category": "Photography",
    "description": "A beautiful vintage camera in excellent condition",
    "condition": 1
  }'
```

### 2. Listar leilões
```bash
curl -X GET http://localhost:8080/auction
```

### 3. Obter detalhes de um leilão específico
Primeiro, recupere o `auctionId` da resposta do passo anterior, depois execute:
```bash
curl -X GET http://localhost:8080/auction/{auction-id}
```

### 4. Fazer um lance
```bash
curl -X POST http://localhost:8080/bid   -H "Content-Type: application/json"   -d '{
    "user_id": "00000000-0000-0000-0000-000000000001",
    "auction_id": "{auction-id}",
    "amount": 150.00
  }'
```

### 5. Listar lances de um leilão
```bash
curl -X GET http://localhost:8080/bid/{auction-id}
```

### 6. Checar lance vencedor
```bash
curl -X GET http://localhost:8080/auction/winner/{auction-id}
```

---