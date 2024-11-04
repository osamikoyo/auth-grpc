package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"

	"auth-grpc/db"
	"auth-grpc/db/models"
	pb "auth-grpc/proto"
)

type Server struct {
	pb.UnimplementedRegisterServer
}

func (s Server) Register(ctx context.Context, u *pb.UserReq) (*pb.UserResp, error) {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	user := models.User{
		Username: u.User.Username,
		Email:    u.User.Email,
		Password: u.User.Password,
	}

	if err := db.Register(user); err != nil {
		loger.Error(err.Error())
		return &pb.UserResp{Status: err.Error()}, err
	}

	return &pb.UserResp{Status: "succes"}, nil
}

func (s Server) Login(ctx context.Context, u *pb.UserLogReq) (*pb.UserResp, error) {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	user := models.User{
		Email: u.User.Email,
		Password: u.User.Password,
	}

	if err := db.Login_User(user); err != nil{
		loger.Error(err.Error())
		return &pb.UserResp{Status: err.Error()}, err
	}

	return &pb.UserResp{Status: "succes"}, nil
}

func (s Server) Auth(ctx context.Context, u *pb.UserAuthReq) (*pb.User, error) {

	user := models.User{
		Email: u.User.Email,
		Token: u.User.Token,
	}

	prof := db.Get_Prof(user)
	return &pb.User{
		Email: prof.Email,
		Password: prof.Password,
		Username: prof.Username,
	}, nil
}

func main() {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	lis, err := net.Listen("tcp", ":2020")
	if err != nil {
		loger.Error(err.Error())
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	pb.RegisterRegisterServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		loger.Error(err.Error())
	}
}
