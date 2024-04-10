package resolver

import graph "github.com/Sanchir01/SandjmaBack_Golang/internal/gql"

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
