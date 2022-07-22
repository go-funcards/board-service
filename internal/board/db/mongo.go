package db

import (
	"context"
	"fmt"
	"github.com/go-funcards/board-service/internal/board"
	"github.com/go-funcards/mongodb"
	"github.com/go-funcards/slice"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var _ board.Storage = (*storage)(nil)

const (
	timeout    = 5 * time.Second
	collection = "boards"
)

type storage struct {
	c   *mongo.Collection
	log logrus.FieldLogger
}

func NewStorage(ctx context.Context, db *mongo.Database, log logrus.FieldLogger) *storage {
	s := &storage{
		c:   db.Collection(collection),
		log: log,
	}
	s.indexes(ctx)
	return s
}

func (s *storage) indexes(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	name, err := s.c.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{"owner_id", 1},
			{"created_at", 1},
			{"members.member_id", 1},
		},
	})
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"collection": collection,
			"error":      err,
		}).Fatal("index not created")
	}

	s.log.WithFields(logrus.Fields{
		"collection": collection,
		"name":       name,
	}).Info("index created")
}

func (s *storage) Save(ctx context.Context, model board.Board) error {
	var write []mongo.WriteModel
	data, err := mongodb.ToBson(model)
	if err != nil {
		return err
	}

	delete(data, "_id")
	delete(data, "owner_id")
	delete(data, "created_at")
	delete(data, "members")

	if deleteMembers := slice.Map(model.Members, func(item board.Member) string {
		return item.MemberID
	}); len(deleteMembers) > 0 {
		s.log.WithFields(logrus.Fields{
			"board_id": model.BoardID,
			"members":  deleteMembers,
		}).Info("delete board's members")

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
				"created_at": model.CreatedAt,
			},
			"$addToSet": bson.M{
				"members": bson.M{"$each": addMembers},
			},
		}),
	)

	s.log.WithField("board_id", model.BoardID).Info("board update")

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := s.c.BulkWrite(ctx, write)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("board update: %s", mongodb.ErrMsgQuery), err)
	}

	s.log.WithFields(logrus.Fields{
		"board_id": model.BoardID,
		"result":   result,
	}).Info("board updated")

	return nil
}

func (s *storage) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	s.log.WithField("id", id).Debug("document delete")
	result, err := s.c.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf(mongodb.ErrMsgQuery, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf(mongodb.ErrMsgQuery, mongo.ErrNoDocuments)
	}
	s.log.WithField("id", id).Debug("document deleted")

	return nil
}

func (s *storage) Find(ctx context.Context, filter board.Filter, index uint64, size uint32) ([]board.Board, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	opts := mongodb.FindOptions(index, size).SetSort(bson.D{{"created_at", -1}})
	cur, err := s.c.Find(ctx, s.build(filter), opts)
	if err != nil {
		return nil, fmt.Errorf(mongodb.ErrMsgQuery, err)
	}
	return mongodb.DecodeAll[board.Board](ctx, cur)
}

func (s *storage) Count(ctx context.Context, filter board.Filter) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	total, err := s.c.CountDocuments(ctx, s.build(filter))
	if err != nil {
		return 0, fmt.Errorf(mongodb.ErrMsgQuery, err)
	}
	return uint64(total), nil
}

func (s *storage) build(filter board.Filter) any {
	f := make(mongodb.Filter, 0)
	if len(filter.BoardIDs) > 0 {
		f = append(f, mongodb.In("_id", filter.BoardIDs))
	}
	if len(filter.OwnerIDs) > 0 && len(filter.MemberIDs) > 0 {
		f = append(f, mongodb.Or(
			mongodb.In("owner_id", filter.OwnerIDs),
			mongodb.In("members.member_id", filter.MemberIDs),
		))
	} else if len(filter.OwnerIDs) > 0 {
		f = append(f, mongodb.In("owner_id", filter.OwnerIDs))
	} else if len(filter.MemberIDs) > 0 {
		f = append(f, mongodb.In("members.member_id", filter.MemberIDs))
	}
	return f.Build()
}
