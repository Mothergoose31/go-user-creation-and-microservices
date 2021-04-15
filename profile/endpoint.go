package account

import(
	"context"

	"github.com/go-kit/kit/endpoint"
)

type endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser endpoint.Endpoint
}
func MakeEndpoints(s Service)Endpoints{
	return endpoints{
		CreateUser:MakeCreateUserEndpoint(s),
		GetUser:MakeGetUserEndpoint(s)
	}
}
fun MakeCreateUserEndpoint(s Service) endpoint.Endpoint{
	return func(ctx context.Context,request (interface{})interface{},error){
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)

		return GetUserResponse{
			Email: email,
		}, err
	}
}