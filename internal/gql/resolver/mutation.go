package resolver

import graph "github.com/Sanchir01/SandjmaBack_Golang/internal/gql"

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
