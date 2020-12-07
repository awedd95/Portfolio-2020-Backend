package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"errors"
	"server/auth"
	"server/graph/model"

	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))

}

func (r *mutationResolver) CreateBlogPost(ctx context.Context, input model.NewBlogPost) (*model.BlogPost, error) {
    user := auth.ForContext(ctx); 
    if user != nil{
	blogPost := model.BlogPost{
		Title: input.Title,
		Body:  input.Body,
	}
	_, err := r.DB.Model(&blogPost).Insert()
	if err != nil {
		return nil, err
	}
	return &blogPost, nil
} else {
    err := errors.New( "you must be logged in to create a post")
    return nil, err 
    }
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (*model.Token, error) {
	user := new(model.User)
	err := r.DB.Model(user).Where("email = ?" , input.Email).Select()
	if err != nil {
        fmt.Print("A db error ocurred", err)
		return nil, err
	}
	if auth.ComparePasswords(user.Password, input.Password) {
		genToken, err := auth.CreateToken(user.ID)
		if err != nil {
			return nil, err
		}
        authToken := model.Token{
            Auth: genToken,
        }
		return &authToken, nil
	}

	return nil, err
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterUser) (*model.Token, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  string(hashedPass),
	}
	_, err = r.DB.Model(&user).Insert()
	fmt.Println(user.ID)
    token, err := auth.CreateToken(user.ID)
    retToken := model.Token{
        Auth: token,
    }
	return &retToken, err
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var projects []*model.Project
	err := r.DB.Model(&projects).Select()
	if err != nil {
	    return nil, err
	}

	return projects, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.BlogPost, error) {
	var posts []*model.BlogPost
	err := r.DB.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	return posts, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
    var users []*model.User
    err := r.DB.Model(&users).Select()
    if err != nil{
        return nil, err
    }
    return users, nil
}


