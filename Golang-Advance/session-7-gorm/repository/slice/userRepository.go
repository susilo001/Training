package slice

import (
	"context"
	"errors"
	"time"

	"github.com/susilo001/golang-advance/session-7-gorm/entity"
	"github.com/susilo001/golang-advance/session-7-gorm/service"
)

type UserRepository struct {
	db     []entity.User
	nextID int
}

func NewUserRepository(db []entity.User) service.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	user.ID = r.nextID
	r.nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	r.db = append(r.db, *user)
	return *user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	for _, user := range r.db {
		if user.ID == id {
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	for i, u := range r.db {
		if u.ID == id {
			user.ID = id
			user.CreatedAt = u.CreatedAt
			user.UpdatedAt = time.Now()
			r.db[i] = user
			return user, nil
		}
	}
	return entity.User{}, errors.New("user not found")
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	for i, user := range r.db {
		if user.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return r.db, nil
}