// Code generated by MockGen. DO NOT EDIT.
// Source: project/e-commerce/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	sql "database/sql"
	sqlc "project/e-commerce/db/sqlc"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateCart mocks base method.
func (m *MockStore) CreateCart(arg0 context.Context, arg1 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockStoreMockRecorder) CreateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockStore)(nil).CreateCart), arg0, arg1)
}

// CreateCartItem mocks base method.
func (m *MockStore) CreateCartItem(arg0 context.Context, arg1 sqlc.CreateCartItemParams) (sqlc.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCartItem", arg0, arg1)
	ret0, _ := ret[0].(sqlc.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCartItem indicates an expected call of CreateCartItem.
func (mr *MockStoreMockRecorder) CreateCartItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCartItem", reflect.TypeOf((*MockStore)(nil).CreateCartItem), arg0, arg1)
}

// CreateComment mocks base method.
func (m *MockStore) CreateComment(arg0 context.Context, arg1 sqlc.CreateCommentParams) (sqlc.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockStoreMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockStore)(nil).CreateComment), arg0, arg1)
}

// CreateOrder mocks base method.
func (m *MockStore) CreateOrder(arg0 context.Context, arg1 sqlc.CreateOrderParams) (sqlc.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockStoreMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockStore)(nil).CreateOrder), arg0, arg1)
}

// CreateOrderItem mocks base method.
func (m *MockStore) CreateOrderItem(arg0 context.Context, arg1 sqlc.CreateOrderItemParams) (sqlc.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderItem", arg0, arg1)
	ret0, _ := ret[0].(sqlc.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderItem indicates an expected call of CreateOrderItem.
func (mr *MockStoreMockRecorder) CreateOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderItem", reflect.TypeOf((*MockStore)(nil).CreateOrderItem), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 sqlc.CreateProductParams) (sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreateProductCategory mocks base method.
func (m *MockStore) CreateProductCategory(arg0 context.Context, arg1 sqlc.CreateProductCategoryParams) (sqlc.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductCategory", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductCategory indicates an expected call of CreateProductCategory.
func (mr *MockStoreMockRecorder) CreateProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductCategory", reflect.TypeOf((*MockStore)(nil).CreateProductCategory), arg0, arg1)
}

// CreateProductRate mocks base method.
func (m *MockStore) CreateProductRate(arg0 context.Context, arg1 sqlc.CreateProductRateParams) (sqlc.ProductRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductRate", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ProductRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductRate indicates an expected call of CreateProductRate.
func (mr *MockStoreMockRecorder) CreateProductRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductRate", reflect.TypeOf((*MockStore)(nil).CreateProductRate), arg0, arg1)
}

// CreateSeller mocks base method.
func (m *MockStore) CreateSeller(arg0 context.Context, arg1 sqlc.CreateSellerParams) (sqlc.Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSeller", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSeller indicates an expected call of CreateSeller.
func (mr *MockStoreMockRecorder) CreateSeller(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSeller", reflect.TypeOf((*MockStore)(nil).CreateSeller), arg0, arg1)
}

// CreateSellerRate mocks base method.
func (m *MockStore) CreateSellerRate(arg0 context.Context, arg1 sqlc.CreateSellerRateParams) (sqlc.SellerRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSellerRate", arg0, arg1)
	ret0, _ := ret[0].(sqlc.SellerRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSellerRate indicates an expected call of CreateSellerRate.
func (mr *MockStoreMockRecorder) CreateSellerRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSellerRate", reflect.TypeOf((*MockStore)(nil).CreateSellerRate), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 sqlc.CreateSessionParams) (sqlc.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 sqlc.CreateUserParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserDeliver mocks base method.
func (m *MockStore) CreateUserDeliver(arg0 context.Context, arg1 sqlc.CreateUserDeliverParams) (sqlc.UserDeliver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserDeliver", arg0, arg1)
	ret0, _ := ret[0].(sqlc.UserDeliver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserDeliver indicates an expected call of CreateUserDeliver.
func (mr *MockStoreMockRecorder) CreateUserDeliver(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserDeliver", reflect.TypeOf((*MockStore)(nil).CreateUserDeliver), arg0, arg1)
}

// DeleteCart mocks base method.
func (m *MockStore) DeleteCart(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCart", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCart indicates an expected call of DeleteCart.
func (mr *MockStoreMockRecorder) DeleteCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockStore)(nil).DeleteCart), arg0, arg1)
}

// DeleteCartItem mocks base method.
func (m *MockStore) DeleteCartItem(arg0 context.Context, arg1 sqlc.DeleteCartItemParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCartItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCartItem indicates an expected call of DeleteCartItem.
func (mr *MockStoreMockRecorder) DeleteCartItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCartItem", reflect.TypeOf((*MockStore)(nil).DeleteCartItem), arg0, arg1)
}

// DeleteComment mocks base method.
func (m *MockStore) DeleteComment(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockStoreMockRecorder) DeleteComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockStore)(nil).DeleteComment), arg0, arg1)
}

// DeleteOrder mocks base method.
func (m *MockStore) DeleteOrder(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockStoreMockRecorder) DeleteOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockStore)(nil).DeleteOrder), arg0, arg1)
}

// DeleteOrderItem mocks base method.
func (m *MockStore) DeleteOrderItem(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderItem indicates an expected call of DeleteOrderItem.
func (mr *MockStoreMockRecorder) DeleteOrderItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderItem", reflect.TypeOf((*MockStore)(nil).DeleteOrderItem), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockStore) DeleteProduct(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockStoreMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockStore)(nil).DeleteProduct), arg0, arg1)
}

// DeleteProductCategory mocks base method.
func (m *MockStore) DeleteProductCategory(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductCategory indicates an expected call of DeleteProductCategory.
func (mr *MockStoreMockRecorder) DeleteProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductCategory", reflect.TypeOf((*MockStore)(nil).DeleteProductCategory), arg0, arg1)
}

// DeleteProductRate mocks base method.
func (m *MockStore) DeleteProductRate(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductRate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductRate indicates an expected call of DeleteProductRate.
func (mr *MockStoreMockRecorder) DeleteProductRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductRate", reflect.TypeOf((*MockStore)(nil).DeleteProductRate), arg0, arg1)
}

// DeleteSeller mocks base method.
func (m *MockStore) DeleteSeller(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSeller", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSeller indicates an expected call of DeleteSeller.
func (mr *MockStoreMockRecorder) DeleteSeller(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSeller", reflect.TypeOf((*MockStore)(nil).DeleteSeller), arg0, arg1)
}

// DeleteSellerRate mocks base method.
func (m *MockStore) DeleteSellerRate(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSellerRate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSellerRate indicates an expected call of DeleteSellerRate.
func (mr *MockStoreMockRecorder) DeleteSellerRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSellerRate", reflect.TypeOf((*MockStore)(nil).DeleteSellerRate), arg0, arg1)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteUserDeliver mocks base method.
func (m *MockStore) DeleteUserDeliver(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserDeliver", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserDeliver indicates an expected call of DeleteUserDeliver.
func (mr *MockStoreMockRecorder) DeleteUserDeliver(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserDeliver", reflect.TypeOf((*MockStore)(nil).DeleteUserDeliver), arg0, arg1)
}

// GetCommentByProductId mocks base method.
func (m *MockStore) GetCommentByProductId(arg0 context.Context, arg1 int64) (sqlc.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentByProductId", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentByProductId indicates an expected call of GetCommentByProductId.
func (mr *MockStoreMockRecorder) GetCommentByProductId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentByProductId", reflect.TypeOf((*MockStore)(nil).GetCommentByProductId), arg0, arg1)
}

// GetCommentByUserId mocks base method.
func (m *MockStore) GetCommentByUserId(arg0 context.Context, arg1 int64) (sqlc.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentByUserId", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentByUserId indicates an expected call of GetCommentByUserId.
func (mr *MockStoreMockRecorder) GetCommentByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentByUserId", reflect.TypeOf((*MockStore)(nil).GetCommentByUserId), arg0, arg1)
}

// GetOrder mocks base method.
func (m *MockStore) GetOrder(arg0 context.Context, arg1 uuid.UUID) (sqlc.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockStoreMockRecorder) GetOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockStore)(nil).GetOrder), arg0, arg1)
}

// GetOrderItemByOrderId mocks base method.
func (m *MockStore) GetOrderItemByOrderId(arg0 context.Context, arg1 uuid.UUID) (sqlc.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItemByOrderId", arg0, arg1)
	ret0, _ := ret[0].(sqlc.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItemByOrderId indicates an expected call of GetOrderItemByOrderId.
func (mr *MockStoreMockRecorder) GetOrderItemByOrderId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItemByOrderId", reflect.TypeOf((*MockStore)(nil).GetOrderItemByOrderId), arg0, arg1)
}

// GetOrderItemByProductId mocks base method.
func (m *MockStore) GetOrderItemByProductId(arg0 context.Context, arg1 int64) ([]sqlc.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItemByProductId", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItemByProductId indicates an expected call of GetOrderItemByProductId.
func (mr *MockStoreMockRecorder) GetOrderItemByProductId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItemByProductId", reflect.TypeOf((*MockStore)(nil).GetOrderItemByProductId), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockStore) GetProduct(arg0 context.Context, arg1 int64) (sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStore)(nil).GetProduct), arg0, arg1)
}

// GetProductCategory mocks base method.
func (m *MockStore) GetProductCategory(arg0 context.Context, arg1 int64) (sqlc.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductCategory", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductCategory indicates an expected call of GetProductCategory.
func (mr *MockStoreMockRecorder) GetProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductCategory", reflect.TypeOf((*MockStore)(nil).GetProductCategory), arg0, arg1)
}

// GetProductRateByProductId mocks base method.
func (m *MockStore) GetProductRateByProductId(arg0 context.Context, arg1 sql.NullInt64) (sqlc.ProductRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductRateByProductId", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ProductRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductRateByProductId indicates an expected call of GetProductRateByProductId.
func (mr *MockStoreMockRecorder) GetProductRateByProductId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductRateByProductId", reflect.TypeOf((*MockStore)(nil).GetProductRateByProductId), arg0, arg1)
}

// GetProductRateByUserId mocks base method.
func (m *MockStore) GetProductRateByUserId(arg0 context.Context, arg1 sql.NullInt64) (sqlc.ProductRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductRateByUserId", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ProductRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductRateByUserId indicates an expected call of GetProductRateByUserId.
func (mr *MockStoreMockRecorder) GetProductRateByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductRateByUserId", reflect.TypeOf((*MockStore)(nil).GetProductRateByUserId), arg0, arg1)
}

// GetSeller mocks base method.
func (m *MockStore) GetSeller(arg0 context.Context, arg1 int64) (sqlc.Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeller", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSeller indicates an expected call of GetSeller.
func (mr *MockStoreMockRecorder) GetSeller(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeller", reflect.TypeOf((*MockStore)(nil).GetSeller), arg0, arg1)
}

// GetSellerRateByUserId mocks base method.
func (m *MockStore) GetSellerRateByUserId(arg0 context.Context, arg1 int64) (sqlc.SellerRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSellerRateByUserId", arg0, arg1)
	ret0, _ := ret[0].(sqlc.SellerRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSellerRateByUserId indicates an expected call of GetSellerRateByUserId.
func (mr *MockStoreMockRecorder) GetSellerRateByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSellerRateByUserId", reflect.TypeOf((*MockStore)(nil).GetSellerRateByUserId), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (sqlc.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 string) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserDeliver mocks base method.
func (m *MockStore) GetUserDeliver(arg0 context.Context, arg1 int64) (sqlc.UserDeliver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDeliver", arg0, arg1)
	ret0, _ := ret[0].(sqlc.UserDeliver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDeliver indicates an expected call of GetUserDeliver.
func (mr *MockStoreMockRecorder) GetUserDeliver(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDeliver", reflect.TypeOf((*MockStore)(nil).GetUserDeliver), arg0, arg1)
}

// ListAllProductCategory mocks base method.
func (m *MockStore) ListAllProductCategory(arg0 context.Context) ([]sqlc.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllProductCategory", arg0)
	ret0, _ := ret[0].([]sqlc.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllProductCategory indicates an expected call of ListAllProductCategory.
func (mr *MockStoreMockRecorder) ListAllProductCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllProductCategory", reflect.TypeOf((*MockStore)(nil).ListAllProductCategory), arg0)
}

// ListCartItem mocks base method.
func (m *MockStore) ListCartItem(arg0 context.Context, arg1 int64) ([]sqlc.ListCartItemRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCartItem", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.ListCartItemRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCartItem indicates an expected call of ListCartItem.
func (mr *MockStoreMockRecorder) ListCartItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCartItem", reflect.TypeOf((*MockStore)(nil).ListCartItem), arg0, arg1)
}

// ListHitProduct mocks base method.
func (m *MockStore) ListHitProduct(arg0 context.Context, arg1 int32) ([]sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHitProduct", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHitProduct indicates an expected call of ListHitProduct.
func (mr *MockStoreMockRecorder) ListHitProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHitProduct", reflect.TypeOf((*MockStore)(nil).ListHitProduct), arg0, arg1)
}

// ListOrder mocks base method.
func (m *MockStore) ListOrder(arg0 context.Context, arg1 int64) ([]sqlc.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrder", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrder indicates an expected call of ListOrder.
func (mr *MockStoreMockRecorder) ListOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrder", reflect.TypeOf((*MockStore)(nil).ListOrder), arg0, arg1)
}

// ListProduct mocks base method.
func (m *MockStore) ListProduct(arg0 context.Context, arg1 sqlc.ListProductParams) ([]sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProduct", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProduct indicates an expected call of ListProduct.
func (mr *MockStoreMockRecorder) ListProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProduct", reflect.TypeOf((*MockStore)(nil).ListProduct), arg0, arg1)
}

// ListProductByType mocks base method.
func (m *MockStore) ListProductByType(arg0 context.Context, arg1 sqlc.ListProductByTypeParams) ([]sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductByType", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductByType indicates an expected call of ListProductByType.
func (mr *MockStoreMockRecorder) ListProductByType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductByType", reflect.TypeOf((*MockStore)(nil).ListProductByType), arg0, arg1)
}

// ListProductCategory mocks base method.
func (m *MockStore) ListProductCategory(arg0 context.Context, arg1 int64) ([]sqlc.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductCategory", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductCategory indicates an expected call of ListProductCategory.
func (mr *MockStoreMockRecorder) ListProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductCategory", reflect.TypeOf((*MockStore)(nil).ListProductCategory), arg0, arg1)
}

// ListUserDeliverByUserId mocks base method.
func (m *MockStore) ListUserDeliverByUserId(arg0 context.Context, arg1 int64) ([]sqlc.UserDeliver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserDeliverByUserId", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.UserDeliver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserDeliverByUserId indicates an expected call of ListUserDeliverByUserId.
func (mr *MockStoreMockRecorder) ListUserDeliverByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserDeliverByUserId", reflect.TypeOf((*MockStore)(nil).ListUserDeliverByUserId), arg0, arg1)
}

// OrderTx mocks base method.
func (m *MockStore) OrderTx(arg0 context.Context, arg1 sqlc.OrderParams) (*sqlc.OrderResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderTx", arg0, arg1)
	ret0, _ := ret[0].(*sqlc.OrderResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrderTx indicates an expected call of OrderTx.
func (mr *MockStoreMockRecorder) OrderTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderTx", reflect.TypeOf((*MockStore)(nil).OrderTx), arg0, arg1)
}

// UpdateComment mocks base method.
func (m *MockStore) UpdateComment(arg0 context.Context, arg1 sqlc.UpdateCommentParams) (sqlc.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockStoreMockRecorder) UpdateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockStore)(nil).UpdateComment), arg0, arg1)
}

// UpdateCommentStatus mocks base method.
func (m *MockStore) UpdateCommentStatus(arg0 context.Context, arg1 sqlc.UpdateCommentStatusParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCommentStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCommentStatus indicates an expected call of UpdateCommentStatus.
func (mr *MockStoreMockRecorder) UpdateCommentStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommentStatus", reflect.TypeOf((*MockStore)(nil).UpdateCommentStatus), arg0, arg1)
}

// UpdateOrderInfo mocks base method.
func (m *MockStore) UpdateOrderInfo(arg0 context.Context, arg1 sqlc.UpdateOrderInfoParams) (sqlc.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderInfo", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrderInfo indicates an expected call of UpdateOrderInfo.
func (mr *MockStoreMockRecorder) UpdateOrderInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderInfo", reflect.TypeOf((*MockStore)(nil).UpdateOrderInfo), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockStore) UpdateProduct(arg0 context.Context, arg1 sqlc.UpdateProductParams) (sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockStoreMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStore)(nil).UpdateProduct), arg0, arg1)
}

// UpdateProductCategory mocks base method.
func (m *MockStore) UpdateProductCategory(arg0 context.Context, arg1 sqlc.UpdateProductCategoryParams) (sqlc.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductCategory", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductCategory indicates an expected call of UpdateProductCategory.
func (mr *MockStoreMockRecorder) UpdateProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductCategory", reflect.TypeOf((*MockStore)(nil).UpdateProductCategory), arg0, arg1)
}

// UpdateProductRateByUserId mocks base method.
func (m *MockStore) UpdateProductRateByUserId(arg0 context.Context, arg1 sqlc.UpdateProductRateByUserIdParams) (sqlc.ProductRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductRateByUserId", arg0, arg1)
	ret0, _ := ret[0].(sqlc.ProductRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductRateByUserId indicates an expected call of UpdateProductRateByUserId.
func (mr *MockStoreMockRecorder) UpdateProductRateByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductRateByUserId", reflect.TypeOf((*MockStore)(nil).UpdateProductRateByUserId), arg0, arg1)
}

// UpdateProductStockSell mocks base method.
func (m *MockStore) UpdateProductStockSell(arg0 context.Context, arg1 sqlc.UpdateProductStockSellParams) (sqlc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductStockSell", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductStockSell indicates an expected call of UpdateProductStockSell.
func (mr *MockStoreMockRecorder) UpdateProductStockSell(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductStockSell", reflect.TypeOf((*MockStore)(nil).UpdateProductStockSell), arg0, arg1)
}

// UpdateSellerRate mocks base method.
func (m *MockStore) UpdateSellerRate(arg0 context.Context, arg1 sqlc.UpdateSellerRateParams) (sqlc.SellerRate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSellerRate", arg0, arg1)
	ret0, _ := ret[0].(sqlc.SellerRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSellerRate indicates an expected call of UpdateSellerRate.
func (mr *MockStoreMockRecorder) UpdateSellerRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSellerRate", reflect.TypeOf((*MockStore)(nil).UpdateSellerRate), arg0, arg1)
}

// UpdateUserDeliver mocks base method.
func (m *MockStore) UpdateUserDeliver(arg0 context.Context, arg1 sqlc.UpdateUserDeliverParams) (sqlc.UserDeliver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDeliver", arg0, arg1)
	ret0, _ := ret[0].(sqlc.UserDeliver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserDeliver indicates an expected call of UpdateUserDeliver.
func (mr *MockStoreMockRecorder) UpdateUserDeliver(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDeliver", reflect.TypeOf((*MockStore)(nil).UpdateUserDeliver), arg0, arg1)
}

// UpdateUserInfo mocks base method.
func (m *MockStore) UpdateUserInfo(arg0 context.Context, arg1 sqlc.UpdateUserInfoParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo.
func (mr *MockStoreMockRecorder) UpdateUserInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockStore)(nil).UpdateUserInfo), arg0, arg1)
}

// UpdateUserPassword mocks base method.
func (m *MockStore) UpdateUserPassword(arg0 context.Context, arg1 sqlc.UpdateUserPasswordParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockStoreMockRecorder) UpdateUserPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockStore)(nil).UpdateUserPassword), arg0, arg1)
}
