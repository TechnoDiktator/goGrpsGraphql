package main

type Account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// Orders []*Order `json:"orders"`  // commented out because we want autogen to generate a resolver for this
}

/*


type Account{
    id: String!
    name: String!
    orders: [Order!]!
}


*/
