package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/susilo001/golang-advance/session-6-db-pgx/entity"
)

func TestCreateUser(t *testing.T) {
	mockRepo := &MockUserRepository{}
	userService := NewUserService(mockRepo)

	t.Run("Create User Service - Success", func(t *testing.T) {
		user := &entity.User{Name: "test", Email: "test@test.com", Password: "password"}
		createdUser := userService.CreateUser(user)

		assert.Equal(t, 0, createdUser.ID)
		assert.Equal(t, "test", createdUser.Name)
		assert.Equal(t, "test@test.com", createdUser.Email)
		assert.Equal(t, "password", createdUser.Password)
	})
}
func TestUpdateUser(t *testing.T) {
	mockRepo := &MockUserRepository{}
	userService := NewUserService(mockRepo)

	user := &entity.User{Name: "test", Email: "test@test.com", Password: "password"}
	createdUser := userService.CreateUser(user)

	t.Run("Update User - Success", func(t *testing.T) {
		updatedUser := entity.User{Name: "Bambang", Email: "bambang@bambang.com", Password: "password"}
		result, err := userService.UpdateUser(createdUser.ID, updatedUser)

		assert.NoError(t, err)
		assert.Equal(t, "Bambang", result.Name)
		assert.Equal(t, "bambang@bambang.com", result.Email)
	})

	t.Run("UpdateUser - UserNotFound", func(t *testing.T) {
		updatedUser := entity.User{Name: "Updated", Email: "updated@example.com", Password: "password"}
		_, err := userService.UpdateUser(99, updatedUser)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}
func TestDeleteUser(t *testing.T) {
	mockRepo := &MockUserRepository{}
	userService := NewUserService(mockRepo)

	user := &entity.User{Name: "Test1", Email: "test1@example.com", Password: "password"}

	createdUser := userService.CreateUser(user)

	t.Run("Delete User - Success", func(t *testing.T) {
		err := userService.DeleteUser(createdUser.ID)

		assert.NoError(t, err)
	})
	t.Run("Delete User - Not Found", func(t *testing.T) {
		err := userService.DeleteUser(99)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}
func TestGetAllUsers(t *testing.T) {
	mockRepo := &MockUserRepository{}
	userService := NewUserService(mockRepo)

	user1 := &entity.User{Name: "Test1", Email: "test1@example.com", Password: "password"}
	user2 := &entity.User{Name: "Test2", Email: "test2@example.com", Password: "password"}

	userService.CreateUser(user1)
	userService.CreateUser(user2)

	t.Run("GetAllUsers - Success", func(t *testing.T) {
		users := userService.GetAllUsers()
		assert.Equal(t, 2, len(users))
		assert.Equal(t, "Test1", users[0].Name)
		assert.Equal(t, "Test2", users[1].Name)
	})
}
func TestGetUserByID(t *testing.T) {
	mockRepo := &MockUserRepository{}
	userService := NewUserService(mockRepo)

	user := &entity.User{Name: "test", Email: "test@test.com", Password: "password"}
	createdUser := userService.CreateUser(user)

	t.Run("Get User - Success", func(t *testing.T) {
		result, err := userService.GetUserByID(createdUser.ID)

		assert.NoError(t, err)
		assert.Equal(t, "test", result.Name)
		assert.Equal(t, "test@test.com", result.Email)
	})

	t.Run("UpdateUser - UserNotFound", func(t *testing.T) {
		_, err := userService.GetUserByID(99)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}
