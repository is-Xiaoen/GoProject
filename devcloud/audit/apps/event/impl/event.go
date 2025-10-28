package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/types"
	"github.com/is-Xiaoen/GoProject/devcloud/audit/apps/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SaveEvent 存储
func (s *EventServiceImpl) SaveEvent(ctx context.Context, in *types.Set[*event.Event]) error {
	s.log.Debug().Msgf("events: %s", in)

	_, err := s.col.InsertMany(ctx, in.ToAny())
	if err != nil {
		return err
	}
	return nil
}

// QueryEvent 存储
func (s *EventServiceImpl) QueryEvent(ctx context.Context, in *event.QueryEventRequest) (*types.Set[*event.Event], error) {
	set := types.NewSet[*event.Event]()

	// 查询条件
	filter := bson.M{}

	opt := options.Find()
	opt.SetLimit(int64(in.PageSize))
	opt.SetSkip(in.ComputeOffset())

	cursor, err := s.col.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		e := event.NewEvent()
		if err := cursor.Decode(e); err != nil {
			return nil, err
		}
		set.Add(e)
	}
	return set, nil
}
