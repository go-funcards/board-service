package board

import (
	"context"
	"github.com/go-funcards/board-service/proto/v1"
	"github.com/go-funcards/slice"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1.BoardServer = (*server)(nil)

type server struct {
	v1.UnimplementedBoardServer
	storage Storage
}

func NewBoardServer(storage Storage) *server {
	return &server{storage: storage}
}

func (s *server) CreateBoard(ctx context.Context, in *v1.CreateBoardRequest) (*emptypb.Empty, error) {
	err := s.storage.Save(ctx, CreateBoard(in))

	return s.empty(err)
}

func (s *server) UpdateBoard(ctx context.Context, in *v1.UpdateBoardRequest) (*emptypb.Empty, error) {
	err := s.storage.Save(ctx, UpdateBoard(in))

	return s.empty(err)
}

func (s *server) DeleteBoard(ctx context.Context, in *v1.DeleteBoardRequest) (*emptypb.Empty, error) {
	err := s.storage.Delete(ctx, in.GetBoardId())

	return s.empty(err)
}

func (s *server) GetBoards(ctx context.Context, in *v1.BoardsRequest) (*v1.BoardsResponse, error) {
	filter := CreateFilter(in)

	data, err := s.storage.Find(ctx, filter, in.GetPageIndex(), in.GetPageSize())
	if err != nil {
		return nil, err
	}

	total := uint64(len(data))
	if len(in.GetBoardIds()) == 0 && uint64(in.GetPageSize()) == total {
		if total, err = s.storage.Count(ctx, filter); err != nil {
			return nil, err
		}
	}

	return &v1.BoardsResponse{
		Boards: slice.Map(data, func(item Board) *v1.BoardsResponse_Board {
			return item.toProto()
		}),
		Total: total,
	}, nil
}

func (s *server) empty(err error) (*emptypb.Empty, error) {
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
