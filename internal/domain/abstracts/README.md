# Abstracts

This directory contains the abstracts for the application. The abstracts provide a high-level overview of the application's architecture, design, and implementation details. They are intended to give readers a quick understanding of the application without diving into the source code.

## Example Abstracts

```go
package abstracts

type UserRepo interface {
	Save(ctx context.Context, user *entities.User) error
	FindByToken(ctx context.Context, token string) (*entities.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByPhone(ctx context.Context, phone string) (*entities.User, error)
	Filter(ctx context.Context, req *list.PagiRequest, search string, isActive string) (*list.PagiResponse[*entities.User], error)
	FilterByRoles(ctx context.Context, req *list.PagiRequest, roles []string) (*list.PagiResponse[*entities.User], error)
	ListByIds(ctx context.Context, userIds []uuid.UUID) ([]*entities.User, error)
}

type Repositories struct {
	ContactRepo ContactRepo
	VerifyRepo  VerifyRepo
	SessionRepo SessionRepo
	UserRepo    UserRepo
}
```

This abstract defines the `UserRepo` interface, which specifies the methods for interacting with the user data in the application. The `Repositories` struct aggregates all the repository interfaces into a single struct for easy access and dependency injection.
