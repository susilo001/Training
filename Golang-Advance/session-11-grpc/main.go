package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/susilo001/golang-advance/session-11-grpc/proto"
	"google.golang.org/grpc"
)

const address = "localhost:8080"

type server struct {
	proto.UnimplementedUserServiceServer
	users map[int64]*proto.User
	mu    sync.Mutex
}

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &server{
		users: make(map[int64]*proto.User),
	})

	log.Println("Server is running on port:", address)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}

func (s *server) GetAllUsers(ctx context.Context, req *proto.GetUsersResponse) (*proto.GetUsersResponse, error) {
	s.mu.Lock()	
	defer s.mu.Unlock()

	var users []*proto.User
	for _, user := range s.users {
		users = append(users, user)
	}
	return &proto.GetUsersResponse{Users: users}, nil
}

func (s *server) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[req.GetId()]
	if !exists {
		return &proto.GetUserResponse{Err: "User not found"}, nil
	}
	return &proto.GetUserResponse{Data: user}, nil
}

func (s *server) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := req.GetUser()
	s.users[user.GetId()] = user
	return &proto.CreateUserResponse{Data: user}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := req.Data
	if _, exists := s.users[user.GetId()]; !exists {
		return &proto.UpdateUserResponse{Err: "User not found"}, nil
	}
	s.users[user.GetId()] = user
	return &proto.UpdateUserResponse{Data: user}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[req.GetId()]; !exists {
		return &proto.DeleteUserResponse{Err: "User not found"}, nil
	}
	delete(s.users, req.GetId())
	return &proto.DeleteUserResponse{Id: req.GetId()}, nil
}
