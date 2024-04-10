package resolver

import (
	"context"
	"fmt"

	"github.com/Sanchir01/SandjmaBack_Golang/internal/gql/model"
)

func (r *queryResolver) GetAllColors(ctx context.Context) ([]*model.ReturnColors, error) {
	fmt.Println("GetAllColors")
    return []*model.ReturnColors{
        
    }, nil
}
