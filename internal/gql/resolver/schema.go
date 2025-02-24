package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/Sanchir01/SandjmaBack_Golang/internal/gql/model"
)

// CrateInsolation is the resolver for the crateInsolation field.
func (r *mutationResolver) CrateInsolation(ctx context.Context, createInsolationInput model.CreateInsolationInput) (*model.Insolation, error) {
	panic(fmt.Errorf("not implemented: CrateInsolation - crateInsolation"))
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, createCategoryInput model.CreateCategoryInput) (*model.ResponseCategory, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// CreateColor is the resolver for the createColor field.
func (r *mutationResolver) CreateColor(ctx context.Context, createReturnColorsInput model.CreateColorInput) (*model.ReturnColors, error) {
	panic(fmt.Errorf("not implemented: CreateColor - createColor"))
}

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, createProductInput model.CreateProductInput) (*model.ReturnFieldByCreateProduct, error) {
	panic(fmt.Errorf("not implemented: CreateProduct - createProduct"))
}

// CreateProductColor is the resolver for the createProductColor field.
func (r *mutationResolver) CreateProductColor(ctx context.Context, productColorInput model.ProductColorInput) (*model.ProductColor, error) {
	panic(fmt.Errorf("not implemented: CreateProductColor - createProductColor"))
}

// CreateSize is the resolver for the createSize field.
func (r *mutationResolver) CreateSize(ctx context.Context, crateSizeInput model.CreateSizeInput) (*model.Size, error) {
	panic(fmt.Errorf("not implemented: CreateSize - createSize"))
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, id float64) (*model.ResponseCategory, error) {
	panic(fmt.Errorf("not implemented: DeleteCategory - deleteCategory"))
}

// DeleteColor is the resolver for the deleteColor field.
func (r *mutationResolver) DeleteColor(ctx context.Context, deleteReturnColorsInput model.CreateColorInput) (*model.ReturnColors, error) {
	panic(fmt.Errorf("not implemented: DeleteColor - deleteColor"))
}

// DeleteInsolation is the resolver for the deleteInsolation field.
func (r *mutationResolver) DeleteInsolation(ctx context.Context, deleteInsolationInput model.DeleteInsolationInput) (*model.Insolation, error) {
	panic(fmt.Errorf("not implemented: DeleteInsolation - deleteInsolation"))
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, deleteProductByID model.GetProductByID) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: DeleteProduct - deleteProduct"))
}

// DeleteProductColor is the resolver for the deleteProductColor field.
func (r *mutationResolver) DeleteProductColor(ctx context.Context, productColorInput model.ProductColorInput) (*model.ProductColor, error) {
	panic(fmt.Errorf("not implemented: DeleteProductColor - deleteProductColor"))
}

// DeleteSize is the resolver for the deleteSize field.
func (r *mutationResolver) DeleteSize(ctx context.Context, id float64) (*model.Size, error) {
	panic(fmt.Errorf("not implemented: DeleteSize - deleteSize"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, loginInput model.LoginInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: Logout - logout"))
}

// NewToken is the resolver for the newToken field.
func (r *mutationResolver) NewToken(ctx context.Context) (*model.NewTokensResponse, error) {
	panic(fmt.Errorf("not implemented: NewToken - newToken"))
}

// PlaceOrderOne is the resolver for the placeOrderOne field.
func (r *mutationResolver) PlaceOrderOne(ctx context.Context, createOrderInput model.CreateOrderInput) (string, error) {
	panic(fmt.Errorf("not implemented: PlaceOrderOne - placeOrderOne"))
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, authInput model.AuthInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, updateCategoryInput model.UpdateCategoryInput) (*model.ResponseCategory, error) {
	panic(fmt.Errorf("not implemented: UpdateCategory - updateCategory"))
}

// UpdateColor is the resolver for the updateColor field.
func (r *mutationResolver) UpdateColor(ctx context.Context, updateReturnColorsInput model.CreateColorInput) (*model.ReturnColors, error) {
	panic(fmt.Errorf("not implemented: UpdateColor - updateColor"))
}

// UpdateInsolation is the resolver for the updateInsolation field.
func (r *mutationResolver) UpdateInsolation(ctx context.Context, updateInsolationInput model.UpdateInsolationInput) (*model.Insolation, error) {
	panic(fmt.Errorf("not implemented: UpdateInsolation - updateInsolation"))
}

// UpdateProductColor is the resolver for the updateProductColor field.
func (r *mutationResolver) UpdateProductColor(ctx context.Context, productColorInput model.ProductColorInput) (*model.ProductColor, error) {
	panic(fmt.Errorf("not implemented: UpdateProductColor - updateProductColor"))
}

// UpdateProfile is the resolver for the updateProfile field.
func (r *mutationResolver) UpdateProfile(ctx context.Context, updateUserProfileInput model.UpdateUserProfileInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateProfile - updateProfile"))
}

// GetAllCategories is the resolver for the getAllCategories field.
func (r *queryResolver) GetAllCategories(ctx context.Context) ([]*model.ResponseCategory, error) {
	panic(fmt.Errorf("not implemented: GetAllCategories - getAllCategories"))
}

// GetAllColors is the resolver for the getAllColors field.


// GetAllInsolation is the resolver for the getAllInsolation field.
func (r *queryResolver) GetAllInsolation(ctx context.Context) ([]*model.Insolation, error) {
	panic(fmt.Errorf("not implemented: GetAllInsolation - getAllInsolation"))
}

// GetAllOrders is the resolver for the getAllOrders field.
func (r *queryResolver) GetAllOrders(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: GetAllOrders - getAllOrders"))
}

// GetAllProduct is the resolver for the getAllProduct field.
func (r *queryResolver) GetAllProduct(ctx context.Context) ([]*model.ProductColor, error) {
	panic(fmt.Errorf("not implemented: GetAllProduct - getAllProduct"))
}

// GetAllProducts is the resolver for the getAllProducts field.
func (r *queryResolver) GetAllProducts(ctx context.Context, getAllProductInput model.GetAllProductInput) (*model.AllProductsAndLength, error) {
	panic(fmt.Errorf("not implemented: GetAllProducts - getAllProducts"))
}

// GetAllSize is the resolver for the getAllSize field.
func (r *queryResolver) GetAllSize(ctx context.Context) ([]*model.Size, error) {
	panic(fmt.Errorf("not implemented: GetAllSize - getAllSize"))
}

// GetCategoryByID is the resolver for the getCategoryById field.
func (r *queryResolver) GetCategoryByID(ctx context.Context, getCategoryByIDInput model.GetCategoryByIDInput) (*model.ResponseCategory, error) {
	panic(fmt.Errorf("not implemented: GetCategoryByID - getCategoryById"))
}

// GetProductByColor is the resolver for the getProductByColor field.
func (r *queryResolver) GetProductByColor(ctx context.Context, getProductByColor model.GetProductByColor) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: GetProductByColor - getProductByColor"))
}

// GetProductByID is the resolver for the getProductById field.
func (r *queryResolver) GetProductByID(ctx context.Context, getProductByID model.GetProductByID) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: GetProductByID - getProductById"))
}

// GetProfile is the resolver for the getProfile field.
func (r *queryResolver) GetProfile(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetProfile - getProfile"))
}

// Mutation returns graph.MutationResolver implementation.

// Query returns graph.QueryResolver implementation.
