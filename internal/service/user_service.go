package service

/*
 * not sure the import items could change to from local rather than github.com/zinan-c/howGO
 * because the go.mod module name is changed to howgo
 */
import (
    "context"
    "github.com/zinan-c/howGO/internal/model"
    "github.com/zinan-c/howGO/internal/repository"
)

type UserService interface {
    RegisterUser(ctx context.Context, user *model.User) error
    GetUser(ctx context.Context, id int) (*model.User, error)
}

type userService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
    return &userService{userRepo: userRepo}
}