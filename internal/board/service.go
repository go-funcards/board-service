package board

import (
	"context"
	"github.com/go-funcards/board-service/proto/v1"
	"github.com/go-funcards/slice"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ v1.BoardServer = (*boardService)(nil)

type boardService struct {
	v1.UnimplementedBoardServer
	storage Storage
	log     *zap.Logger
}

func NewBoardService(storage Storage, logger *zap.Logger) *boardService {
	return &boardService{
		storage: storage,
		log:     logger,
	}
}

func (s *boardService) CreateBoard(ctx context.Context, in *v1.CreateBoardRequest) (*emptypb.Empty, error) {
	if err := s.storage.Save(ctx, CreateBoard(in)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *boardService) UpdateBoard(ctx context.Context, in *v1.UpdateBoardRequest) (*emptypb.Empty, error) {
	if err := s.storage.Save(ctx, UpdateBoard(in)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *boardService) DeleteBoard(ctx context.Context, in *v1.DeleteBoardRequest) (*emptypb.Empty, error) {
	if err := s.storage.Delete(ctx, in.GetBoardId()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *boardService) GetBoards(ctx context.Context, in *v1.BoardsRequest) (*v1.BoardsResponse, error) {
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
			return item.toResponse()
		}),
		Total: total,
	}, nil
}
