package mock_accountservice

import (
	"github.com/Kimenzo/any-sync/accountservice"
	"github.com/Kimenzo/any-sync/commonspace/object/accountdata"
	"go.uber.org/mock/gomock"
)

func NewAccountServiceWithAccount(ctrl *gomock.Controller, acc *accountdata.AccountKeys) *MockService {
	mock := NewMockService(ctrl)
	mock.EXPECT().Name().Return(accountservice.CName).AnyTimes()
	mock.EXPECT().Init(gomock.Any()).AnyTimes()
	mock.EXPECT().Account().Return(acc).AnyTimes()
	return mock
}
