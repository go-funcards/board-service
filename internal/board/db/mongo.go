package db

import (
	"context"
	"fmt"
	"github.com/go-funcards/board-service/internal/board"
	"github.com/go-funcards/mongodb"
	"github.com/go-funcards/slice"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

var _ board.Storage = (*storage)(nil)

const (
	timeout    = 5 * time.Second
	collection = "boards"
)

type storage struct {
	c mongodb.Collection[board.Board]
}

func NewStorage(ctx context.Context, db *mongo.Database, logger *zap.Logger) (*storage, error) {
	s := &storage{c: mongodb.Collection[board.Board]{
		Inner: db.Collection(collection),
		Log:   logger,
	}}

	if err := s.indexes(ctx); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *storage) indexes(ctx context.Context) error {
	name, err := s.c.Inner.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{"owner_id", 1},
			{"type", 1},
			{"created_at", 1},
			{"members.member_id", 1},
		},
	})
	if err == nil {
		s.c.Log.Info("index created", zap.String("collection", collection), zap.String("name", name))
	}
	return err
}

func (s *storage) Save(ctx context.Context, model board.Board) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var write []mongo.WriteModel
	data, err := s.c.ToM(model)
	if err != nil {
		return err
	}

	delete(data, "_id")
	delete(data, "owner_id")
	delete(data, "type")
	delete(data, "created_at")
	delete(data, "members")

	if deleteMembers := slice.Map(model.Members, func(item board.Member) string {
		return item.MemberID
	}); len(deleteMembers) > 0 {
		s.c.Log.Info("delete members from board", zap.String("board_id", model.BoardID), zap.Strings("members", deleteMembers))

		write = append(write, mongo.
			NewUpdateOneModel().
			SetFilter(bson.M{"_id": model.BoardID}).
			SetUpdate(bson.M{
				"$pull": bson.M{
					"members": bson.M{
						"member_id": bson.M{
							"$in": deleteMembers,
						},
					},
				},
			}),
		)
	}

	addMembers := slice.Filter(model.Members, func(item board.Member) bool {
		return !item.Delete
	})

	write = append(write, mongo.
		NewUpdateOneModel().
		SetUpsert(true).
		SetFilter(bson.M{"_id": model.BoardID}).
		SetUpdate(bson.M{
			"$set": data,
			"$setOnInsert": bson.M{
				"owner_id":   model.OwnerID,
				"type":       model.Type,
				"created_at": model.CreatedAt,
			},
			"$addToSet": bson.M{
				"members": bson.M{"$each": addMembers},
			},
		}),
	)

	s.c.Log.Debug("bulk update")

	result, err := s.c.Inner.BulkWrite(ctx, write, options.BulkWrite())
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("bulk update: %s", mongodb.ErrMsgQuery), err)
	}

	s.c.Log.Info("document updated", zap.String("board_id", model.BoardID), zap.Any("result", result))

	return nil
}

func (s *storage) Delete(ctx context.Context, id string) error {
	return s.c.DeleteOne(ctx, bson.M{"_id": id})
}

func (s *storage) Find(ctx context.Context, filter board.Filter, index uint64, size uint32) ([]board.Board, error) {
	return s.c.Find(ctx, s.filter(filter), s.c.FindOptions(index, size).SetSort(bson.D{{"created_at", -1}}))
}

func (s *storage) Count(ctx context.Context, filter board.Filter) (uint64, error) {
	return s.c.CountDocuments(ctx, s.filter(filter))
}

func (s *storage) filter(filter board.Filter) bson.M {
	f := make(bson.M)
	if len(filter.Types) > 0 {
		f["type"] = bson.M{"$in": filter.Types}
	}
	if len(filter.BoardIDs) > 0 {
		f["_id"] = bson.M{"$in": filter.BoardIDs}
	}
	if len(filter.OwnerIDs) > 0 && len(filter.MemberIDs) > 0 {
		f["$or"] = bson.A{
			bson.M{"owner_id": bson.M{"$in": filter.OwnerIDs}},
			bson.M{"members.member_id": bson.M{"$in": filter.MemberIDs}},
		}
	} else if len(filter.OwnerIDs) > 0 {
		f["members.owner_id"] = bson.M{"$in": filter.OwnerIDs}
	} else if len(filter.MemberIDs) > 0 {
		f["members.member_id"] = bson.M{"$in": filter.MemberIDs}
	}
	return f
}
