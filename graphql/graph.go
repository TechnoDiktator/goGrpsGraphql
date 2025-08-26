package main

//import "github.com/99designs/gqlgen/graphql"

//type Server struct {
// accountClient *account.Client
// catalogClient *catalog.Client
// orderClient   *order.Client
//}

//from the graphql gateway we want to be able t o call three clients
//account client, product client and the order client

//These cliemts are in turn connected to GRPC servers running as seperate microservices
//The servers are then connected to their respective databases

// takes in the urls of the three services and dials them to create the clients
// func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
// 	accountClient, err := account.NewClient(accountUrl)
// 	if err != nil {
// 		accountClient.Close()
// 		return nil, err
// 	}
// 	catalogClient, err := catalog.NewClient(catalogUrl)
// 	if err != nil {
// 		catalogClient.Close()
// 		return nil, err
// 	}
// 	orderClient, err := order.NewClient(orderUrl)
// 	if err != nil {
// 		orderClient.Close()
// 		return nil, err
// 	}

// 	//all these clients point to grpc servers running as microservices
// 	return &Server{
// 		accountClient: accountClient,
// 		catalogClient: catalogClient,
// 		orderClient:   orderClient,
// 	}, nil

// }

// func (s *Server) Close() {
// 	s.accountClient.Close()
// 	s.catalogClient.Close()
// 	s.orderClient.Close()
// }

// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{server: s}
// }

// func (s *Server) Query() QueryResolver {

// 	return &queryResolver{server: s}

// }

// func (s *Server) Account() AccountResolver {

// 	return &accountResolver{server: s}

// }

// func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
// 	return NewExecutableSchema(Config{Resolvers: s})
// }
