package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/robotomize/gb-golang/homework/03-04-umanager/pkg/api/apiv1"
	"gitlab.com/robotomize/gb-golang/homework/03-04-umanager/pkg/pb"
	"google.golang.org/grpc"
)

type MockLinksClient struct {
	ctrl     *gomock.Controller
	recorder *MockLinksClientMockRecorder
}

type MockLinksClientMockRecorder struct {
	mock *MockLinksClient
}

func NewMockLinksClient(ctrl *gomock.Controller) *MockLinksClient {
	mock := &MockLinksClient{ctrl: ctrl}
	mock.recorder = &MockLinksClientMockRecorder{mock}
	return mock
}

func (m *MockLinksClient) EXPECT() *MockLinksClientMockRecorder {
	return m.recorder
}

func (m *MockLinksClient) CreateLink(ctx context.Context, in *pb.CreateLinkRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLink", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinksClientMockRecorder) CreateLink(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLink", reflect.TypeOf((*MockLinksClient)(nil).CreateLink), ctx, in)
}

func (m *MockLinksClient) DeleteLink(ctx context.Context, in *pb.DeleteLinkRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLink", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLinksClientMockRecorder) DeleteLink(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLink", reflect.TypeOf((*MockLinksClient)(nil).DeleteLink), ctx, in)
}

func TestPostLinks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockLinksClient(ctrl)
	handler := newLinksHandler(mockClient)

	link := apiv1.LinkCreate{
		Id:     "1",
		Title:  "Test",
		Url:    "http://test.com",
		Images: []string{"test1", "test2"},
		Tags:   []string{"tag1", "tag2"},
		UserId: "user1",
	}

	mockClient.EXPECT().CreateLink(gomock.Any(), &pb.CreateLinkRequest{
		Id:     link.Id,
		Title:  link.Title,
		Url:    link.Url,
		Images: link.Images,
		Tags:   link.Tags,
		UserId: link.UserId,
	}).Return(&pb.Empty{}, nil)

	body, _ := json.Marshal(link)
	req, _ := http.NewRequest("POST", "/links", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.PostLinks(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestDeleteLinksId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockLinksClient(ctrl)
	handler := newLinksHandler(mockClient)

	linkID := "1"

	mockClient.EXPECT().DeleteLink(gomock.Any(), &pb.DeleteLinkRequest{Id: linkID}).Return(&pb.Empty{}, nil)

	req, _ := http.NewRequest("DELETE", "/links/"+linkID, nil)
	rr := httptest.NewRecorder()

	handler.DeleteLinksId(rr, req, linkID)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}

type MockUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockUsersClientMockRecorder
}

type MockUsersClientMockRecorder struct {
	mock *MockUsersClient
}

func NewMockUsersClient(ctrl *gomock.Controller) *MockUsersClient {
	mock := &MockUsersClient{ctrl: ctrl}
	mock.recorder = &MockUsersClientMockRecorder{mock}
	return mock
}

func (m *MockUsersClient) EXPECT() *MockUsersClientMockRecorder {
	return m.recorder
}

func (m *MockUsersClient) CreateUser(ctx context.Context, in *pb.CreateUserRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersClientMockRecorder) CreateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsersClient)(nil).CreateUser), ctx, in)
}

func (m *MockUsersClient) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest, opts ...grpc.CallOption) (*pb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, in)
	ret0, _ := ret[0].(*pb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUsersClientMockRecorder) DeleteUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUsersClient)(nil).DeleteUser), ctx, in)
}

func TestPostUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockUsersClient(ctrl)
	handler := newUsersHandler(mockClient)

	user := apiv1.UserCreate{
		Id:       "1",
		Username: "test",
		Password: "123",
	}

	mockClient.EXPECT().CreateUser(gomock.Any(), &pb.CreateUserRequest{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
	}).Return(&pb.Empty{}, nil)

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.PostUsers(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestDeleteUsersId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockUsersClient(ctrl)
	handler := newUsersHandler(mockClient)

	userID := "1"

	mockClient.EXPECT().DeleteUser(gomock.Any(), &pb.DeleteUserRequest{Id: userID}).Return(&pb.Empty{}, nil)

	req, _ := http.NewRequest("DELETE", "/users/"+userID, nil)
	rr := httptest.NewRecorder()

	handler.DeleteUsersId(rr, req, userID)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
