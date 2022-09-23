# /internal/entity

**Directory for business entities:**
1. Request and response models (entities)
2. Validation for handlers
3. etc...

## **Example**
```go
type CreateSomethingReq struct {
	ID          string      // Not for client (private field)
	CreatedAt   time.Time   // Not for client (private field)
	Name        string      `json:"name" binding:"required"`
	City        string      `form:"city"`
	Age         int         `json:"age"`
}

type GetSomethingResp struct {
	ID          string      `json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	Name        string      `json:"name"`
	City        string      `json:"city"`
	Age         int         `json:"age"`
}
```

Example request. Sending `name` and `age` on json format. However, `city` sending on form data format.
```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Alex","age":35}' \
  https://api.example.com/create?city=Moscow
```

Example response
```json
{
  "id": "123",
  "created_at": "2022-08-20T20:00:00.100000Z",
  "name": "Alex",
  "city": "Moscow",
  "age": 35,
}
```