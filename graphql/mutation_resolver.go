package main

//What is a resolver ?
//A resolver is a function that resolves a value for a type or field in your schema
//What is a mutation resolver ?
//A mutation resolver is a function that resolves a value for a mutation in your schema . Basically it is used to create , update or delete data
//Why is there a server reference in the mutation resolver struct ?
//The server reference is used to access the accountClient which is used to communicate with the gRPC server
type mutationResolver struct {
	server *Server
}

// func (r *mutationResolver) createAccount(ctx context.Context , in AccountInput)(*Account , error){

// }
// func (r *mutationResolver) createProduct(ctx context.Context , in ProductInput)(*Product , error){

// }
// func (r *mutationResolver) createOrder(ctx context.Context ,in OrderInput )(*Account , error){

// }
