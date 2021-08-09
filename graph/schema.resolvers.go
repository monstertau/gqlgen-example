package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graph-ql-example/graph/generated"
	"graph-ql-example/graph/model"
	"math/rand"
)

func (r *mutationResolver) CreateStatistic(ctx context.Context, input model.NewStatistic) (*model.Statistic, error) {
	var points []*model.DataPoint
	if input.Data != nil {
		for _, point := range input.Data.Points {
			p := &model.DataPoint{
				Count:     point.Count,
				Timestamp: point.Timestamp,
				Attribute: point.Attribute,
			}
			points = append(points, p)
		}
	}

	statistic := &model.Statistic{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Duration: input.Duration,
		Interval: input.Interval,
		Mode:     input.Mode,
		Timezone: input.Timezone,
		Data:     &model.Data{Points: points},
	}
	r.statistics = append(r.statistics, statistic)
	return statistic, nil
}

func (r *queryResolver) Statistic(ctx context.Context) ([]*model.Statistic, error) {
	return r.statistics, nil
}

func (r *queryResolver) Search(ctx context.Context, id string) (*model.Statistic, error) {
	for _, statistic := range r.statistics {
		if statistic.ID == id {
			return statistic, nil
		}
	}
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
