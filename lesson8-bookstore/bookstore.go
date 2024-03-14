package main

import (
	"boostore/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// 构建服务 bookstore grpc服务
type server struct {
	pb.UnimplementedBookstoreServer
	//定义数据库的类型引入，就能调用数据实现的方法了
	bs *bookstore
}

// 实现service
func (s *server) ListShelves(ctx context.Context, in *emptypb.Empty) (*pb.ListShelvesResponse, error) {
	//调用orm获取数据
	sl, err := s.bs.ListShelves(ctx)
	if err == gorm.ErrEmptySlice {
		return &pb.ListShelvesResponse{}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}
	//封装数据 因为从数据库获取过来的就是要给模型数据，现在要转换成grpc定义的message 响应类型
	nsl := make([]*pb.Shelf, 0, len(sl))
	for _, s := range sl {
		nsl = append(nsl, &pb.Shelf{
			Id:    s.ID,
			Theme: s.Theme,
			Size:  s.Size,
		})
	}
	//返回数据
	return &pb.ListShelvesResponse{Shelf: nsl}, nil //因为定义的这个服务的响应体是一个书架集合，所以要构建集合
}

// CreateShelf 根据id创建书架
func (s *server) CreateShelf(ctx context.Context, in *pb.CreateShelfRequest) (*pb.Shelf, error) {
	//参数校验
	if len(in.GetShelf().GetTheme()) == 0 {
		return nil, status.Error(codes.Internal, "create failed")
	}
	//准备数据
	data := Shelf{
		Theme: in.GetShelf().GetTheme(),
		Size:  in.GetShelf().GetSize(),
	}

	//去orm创建
	ns, err := s.bs.CreateShelf(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, "create failed ")
	}
	//返回响应
	return &pb.Shelf{Id: ns.ID, Theme: ns.Theme, Size: ns.Size}, nil

}

// GetShelf 创建书架
func (s *server) GetShelf(ctx context.Context, in *pb.GetShelfRequest) (*pb.Shelf, error) {
	//参数检查
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	//去orm层创建
	shelf, err := s.bs.GetShelf(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}
	//封装数据，返回响应
	return &pb.Shelf{Id: shelf.ID, Theme: shelf.Theme, Size: shelf.Size}, nil
}

// DeleteShelf 通过id删除数据
func (s *server) DeleteShelf(ctx context.Context, in *pb.DeleteShelfRequest) (*emptypb.Empty, error) {
	//参数校验
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id ")
	}
	err := s.bs.DeleShelf(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "delete failed")
	}
	return &emptypb.Empty{}, nil
}
