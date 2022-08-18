# /internal/service

**Directory for business logic**

## **Example**

Edit `service.go`. Add new interface `Something`
```go 
package service

import (
    "context"
    "{{MODULE_NAME}}/internal/entity"
    "{{MODULE_NAME}}/internal/storage"
)

type Something interface {
    CreateSomething(ctx context.Context, something *entity.CreateSomethingReq) (*entity.GetSomethingResp, error)
}

type Service struct {
    Something *SomethingService
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
        Something: NewSomethingService(storage.psql.Something, storage.redis.Something),
	}
}
```

Create file `something.go`
```go
package service

import (
	"{{SERVICE_NAME}}/internal/entity"
	"{{SERVICE_NAME}}/internal/storage/psql"
	"{{SERVICE_NAME}}/internal/storage/redis"
)

type SomethingService struct {
	storage psql.Something
	cache redis.Something
}

func NewSomethingService(storage psql.Example, cache redis.Example) *SomethingService {
	return &SomethingService{storage: storage, cache: cache}
}

func (s *SomethingService) CreateSomething(something *entity.CreateSomethingReq) (*entity.GetSomethingResp, error) {
    // business logic 
    something.ID = "123-321-123-321" // It is more logical to generate ID at the database level
    
    somethingID, err := s.storage.CreateSomething(something)
    if err != nil {
        // Return original error from storage to handler. You can handle errors yourself with errors.New("CreateSomething: something went wrong")
        return nil, err 
    }

    getNewSomething, err := s.storage.GetSomethingByID(somethingID)
    if err != nil {
        // Return original error from storage to handler. You can handle errors yourself with errors.New("something not found")
        return nil, err
    }
    
    return getNewSomething, nil
}
```