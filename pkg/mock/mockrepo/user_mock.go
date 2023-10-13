// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/user.go

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	domain "jerseyhub/pkg/domain"
	models "jerseyhub/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserRepository) AddAddress(id int, address models.AddAddress, result bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", id, address, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserRepositoryMockRecorder) AddAddress(id, address, result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserRepository)(nil).AddAddress), id, address, result)
}

// ChangePassword mocks base method.
func (m *MockUserRepository) ChangePassword(id int, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", id, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserRepositoryMockRecorder) ChangePassword(id, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserRepository)(nil).ChangePassword), id, password)
}

// CheckIfFirstAddress mocks base method.
func (m *MockUserRepository) CheckIfFirstAddress(id int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfFirstAddress", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckIfFirstAddress indicates an expected call of CheckIfFirstAddress.
func (mr *MockUserRepositoryMockRecorder) CheckIfFirstAddress(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfFirstAddress", reflect.TypeOf((*MockUserRepository)(nil).CheckIfFirstAddress), id)
}

// CheckUserAvailability mocks base method.
func (m *MockUserRepository) CheckUserAvailability(email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserAvailability", email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckUserAvailability indicates an expected call of CheckUserAvailability.
func (mr *MockUserRepositoryMockRecorder) CheckUserAvailability(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserAvailability", reflect.TypeOf((*MockUserRepository)(nil).CheckUserAvailability), email)
}

// CreditReferencePointsToWallet mocks base method.
func (m *MockUserRepository) CreditReferencePointsToWallet(user_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditReferencePointsToWallet", user_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreditReferencePointsToWallet indicates an expected call of CreditReferencePointsToWallet.
func (mr *MockUserRepositoryMockRecorder) CreditReferencePointsToWallet(user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditReferencePointsToWallet", reflect.TypeOf((*MockUserRepository)(nil).CreditReferencePointsToWallet), user_id)
}

// EditEmail mocks base method.
func (m *MockUserRepository) EditEmail(id int, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditEmail", id, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditEmail indicates an expected call of EditEmail.
func (mr *MockUserRepositoryMockRecorder) EditEmail(id, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditEmail", reflect.TypeOf((*MockUserRepository)(nil).EditEmail), id, email)
}

// EditName mocks base method.
func (m *MockUserRepository) EditName(id int, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditName", id, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditName indicates an expected call of EditName.
func (mr *MockUserRepositoryMockRecorder) EditName(id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditName", reflect.TypeOf((*MockUserRepository)(nil).EditName), id, name)
}

// EditPhone mocks base method.
func (m *MockUserRepository) EditPhone(id int, phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditPhone", id, phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditPhone indicates an expected call of EditPhone.
func (mr *MockUserRepositoryMockRecorder) EditPhone(id, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditPhone", reflect.TypeOf((*MockUserRepository)(nil).EditPhone), id, phone)
}

// FindCartQuantity mocks base method.
func (m *MockUserRepository) FindCartQuantity(cart_id, inventory_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCartQuantity", cart_id, inventory_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCartQuantity indicates an expected call of FindCartQuantity.
func (mr *MockUserRepositoryMockRecorder) FindCartQuantity(cart_id, inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartQuantity", reflect.TypeOf((*MockUserRepository)(nil).FindCartQuantity), cart_id, inventory_id)
}

// FindCategory mocks base method.
func (m *MockUserRepository) FindCategory(inventory_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCategory", inventory_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCategory indicates an expected call of FindCategory.
func (mr *MockUserRepositoryMockRecorder) FindCategory(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCategory", reflect.TypeOf((*MockUserRepository)(nil).FindCategory), inventory_id)
}

// FindIdFromPhone mocks base method.
func (m *MockUserRepository) FindIdFromPhone(phone string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIdFromPhone", phone)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIdFromPhone indicates an expected call of FindIdFromPhone.
func (mr *MockUserRepositoryMockRecorder) FindIdFromPhone(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIdFromPhone", reflect.TypeOf((*MockUserRepository)(nil).FindIdFromPhone), phone)
}

// FindPrice mocks base method.
func (m *MockUserRepository) FindPrice(inventory_id int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPrice", inventory_id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPrice indicates an expected call of FindPrice.
func (mr *MockUserRepositoryMockRecorder) FindPrice(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPrice", reflect.TypeOf((*MockUserRepository)(nil).FindPrice), inventory_id)
}

// FindProductImage mocks base method.
func (m *MockUserRepository) FindProductImage(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductImage", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductImage indicates an expected call of FindProductImage.
func (mr *MockUserRepositoryMockRecorder) FindProductImage(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductImage", reflect.TypeOf((*MockUserRepository)(nil).FindProductImage), id)
}

// FindProductNames mocks base method.
func (m *MockUserRepository) FindProductNames(inventory_id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProductNames", inventory_id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProductNames indicates an expected call of FindProductNames.
func (mr *MockUserRepositoryMockRecorder) FindProductNames(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProductNames", reflect.TypeOf((*MockUserRepository)(nil).FindProductNames), inventory_id)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", user)
	ret0, _ := ret[0].(models.UserSignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), user)
}

// FindUserFromReference mocks base method.
func (m *MockUserRepository) FindUserFromReference(ref string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserFromReference", ref)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserFromReference indicates an expected call of FindUserFromReference.
func (mr *MockUserRepositoryMockRecorder) FindUserFromReference(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserFromReference", reflect.TypeOf((*MockUserRepository)(nil).FindUserFromReference), ref)
}

// FindofferPercentage mocks base method.
func (m *MockUserRepository) FindofferPercentage(category_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindofferPercentage", category_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindofferPercentage indicates an expected call of FindofferPercentage.
func (mr *MockUserRepositoryMockRecorder) FindofferPercentage(category_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindofferPercentage", reflect.TypeOf((*MockUserRepository)(nil).FindofferPercentage), category_id)
}

// GetAddresses mocks base method.
func (m *MockUserRepository) GetAddresses(id int) ([]domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddresses", id)
	ret0, _ := ret[0].([]domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddresses indicates an expected call of GetAddresses.
func (mr *MockUserRepositoryMockRecorder) GetAddresses(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockUserRepository)(nil).GetAddresses), id)
}

// GetCart mocks base method.
func (m *MockUserRepository) GetCart(id int) ([]models.GetCart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", id)
	ret0, _ := ret[0].([]models.GetCart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockUserRepositoryMockRecorder) GetCart(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockUserRepository)(nil).GetCart), id)
}

// GetCartID mocks base method.
func (m *MockUserRepository) GetCartID(id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartID", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartID indicates an expected call of GetCartID.
func (mr *MockUserRepositoryMockRecorder) GetCartID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartID", reflect.TypeOf((*MockUserRepository)(nil).GetCartID), id)
}

// GetPassword mocks base method.
func (m *MockUserRepository) GetPassword(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockUserRepositoryMockRecorder) GetPassword(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockUserRepository)(nil).GetPassword), id)
}

// GetProductsInCart mocks base method.
func (m *MockUserRepository) GetProductsInCart(cart_id int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsInCart", cart_id)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsInCart indicates an expected call of GetProductsInCart.
func (mr *MockUserRepositoryMockRecorder) GetProductsInCart(cart_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsInCart", reflect.TypeOf((*MockUserRepository)(nil).GetProductsInCart), cart_id)
}

// GetReferralCodeFromID mocks base method.
func (m *MockUserRepository) GetReferralCodeFromID(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReferralCodeFromID", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReferralCodeFromID indicates an expected call of GetReferralCodeFromID.
func (mr *MockUserRepositoryMockRecorder) GetReferralCodeFromID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReferralCodeFromID", reflect.TypeOf((*MockUserRepository)(nil).GetReferralCodeFromID), id)
}

// GetUserDetails mocks base method.
func (m *MockUserRepository) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", id)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails.
func (mr *MockUserRepositoryMockRecorder) GetUserDetails(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockUserRepository)(nil).GetUserDetails), id)
}

// RemoveFromCart mocks base method.
func (m *MockUserRepository) RemoveFromCart(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromCart", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromCart indicates an expected call of RemoveFromCart.
func (mr *MockUserRepositoryMockRecorder) RemoveFromCart(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromCart", reflect.TypeOf((*MockUserRepository)(nil).RemoveFromCart), id)
}

// UpdateQuantityAdd mocks base method.
func (m *MockUserRepository) UpdateQuantityAdd(id, inv_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantityAdd", id, inv_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuantityAdd indicates an expected call of UpdateQuantityAdd.
func (mr *MockUserRepositoryMockRecorder) UpdateQuantityAdd(id, inv_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantityAdd", reflect.TypeOf((*MockUserRepository)(nil).UpdateQuantityAdd), id, inv_id)
}

// UpdateQuantityLess mocks base method.
func (m *MockUserRepository) UpdateQuantityLess(id, inv_id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuantityLess", id, inv_id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuantityLess indicates an expected call of UpdateQuantityLess.
func (mr *MockUserRepositoryMockRecorder) UpdateQuantityLess(id, inv_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuantityLess", reflect.TypeOf((*MockUserRepository)(nil).UpdateQuantityLess), id, inv_id)
}

// UserBlockStatus mocks base method.
func (m *MockUserRepository) UserBlockStatus(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserBlockStatus", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserBlockStatus indicates an expected call of UserBlockStatus.
func (mr *MockUserRepositoryMockRecorder) UserBlockStatus(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserBlockStatus", reflect.TypeOf((*MockUserRepository)(nil).UserBlockStatus), email)
}

// UserSignUp mocks base method.
func (m *MockUserRepository) UserSignUp(user models.UserDetails, referal string) (models.UserDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", user, referal)
	ret0, _ := ret[0].(models.UserDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserRepositoryMockRecorder) UserSignUp(user, referal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserRepository)(nil).UserSignUp), user, referal)
}
