package board

import "context"

type Storage interface {
	Save(ctx context.Context, model Board) error
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, filter Filter, index uint64, size uint32) ([]Board, error)
	Count(ctx context.Context, filter Filter) (uint64, error)
}
