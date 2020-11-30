package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"server/graph/model"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBlogPost(ctx context.Context, input model.NewBlogPost) (*model.BlogPost, error) {
    blogPost := model.BlogPost{
        Title: input.Title,
        Body : input.Body,
    }
    _,err := r.DB.Model(&blogPost).Insert()
	if err != nil {
		return nil, err
	}    
    return &blogPost, nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var projects []*model.Project
	dummyProject := model.Project{
		Title:       "test title 2",
		Description: "test description",
		Language:    "test language",
	}
	projects = append(projects, &dummyProject)
	return projects, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.BlogPost, error) {
	panic(fmt.Errorf("not implemented"))
}

