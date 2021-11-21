package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"test_task/graph/generated"
	"test_task/graph/model"
	"time"
)

func (r *mutationResolver) RequestSignInCode(ctx context.Context, input model.RequestSignInCodeInput) (*model.ErrorPayload, error) {
	num, err := strconv.Atoi(input.Phone)
	if err != nil {
		return &model.ErrorPayload{Message: "phone number must be only digits"}, nil
	}
	if (num < 0) || (len(input.Phone) != 9) {
		return &model.ErrorPayload{Message: "phone number must be 9 digits, without \"+\",\"-\""}, nil
	}
	exists := r.Resolver.phonesCodes[input.Phone]
	if exists != 0 {
		fmt.Printf("Code: %d", exists)
		return nil, nil
	}
	rand.Seed(time.Now().UnixNano())
	code := 1000 + rand.Intn(9999-1000)
	fmt.Printf("Code: %d", code)
	r.Resolver.phonesCodes[input.Phone] = code
	return nil, nil
}

func (r *mutationResolver) SignInByCode(ctx context.Context, input model.SignInByCodeInput) (model.SignInOrErrorPayload, error) {
	num, err := strconv.Atoi(input.Phone)
	if err != nil {
		return &model.ErrorPayload{Message: "phone number must be only digits"}, nil
	}

	if (num < 0) || (len(input.Phone) != 9) {
		return &model.ErrorPayload{Message: "phone number must be 9 digits, without \"+\",\"-\""}, nil
	}

	code, err := strconv.Atoi(input.Code)
	if (code < 0) || (len(input.Code) != 4) || (err != nil) {
		return &model.ErrorPayload{Message: "code must be 4 digits"}, nil
	}

	exists := r.Resolver.phonesCodes[input.Phone]
	if code != exists {
		return &model.ErrorPayload{Message: "Wrong code, or number phone"}, nil
	}

	err = r.Authorization.SignIn(input.Phone)
	if err != nil {
		return &model.ErrorPayload{Message: err.Error()}, nil
	}

	token, err := r.Authorization.GenerateToken(input.Phone)
	if err != nil {
		return &model.ErrorPayload{Message: err.Error()}, nil
	}

	return &model.SignInPayload{
		Token: token,
		Viewer: &model.Viewer{
			User: &model.User{
				Phone: input.Phone,
			},
		},
	}, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	products, err := r.Resolver.Product.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*model.Viewer, error) {
	token := ForContext(ctx)
	if token == "" {
		return nil, errors.New("not found auth token")
	}
	phone, err := r.Resolver.Authorization.ParseToken(token)
	if err != nil {
		return nil, errors.New("wrong auth token")
	}
	return &model.Viewer{User: &model.User{Phone: phone}}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
