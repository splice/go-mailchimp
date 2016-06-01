package mailchimp

import (
	"github.com/stretchr/testify/mock"
)

// ClientMock is a mocked object implementing MailchimpInterface
type ClientMock struct {
	mock.Mock
}

// Subscribe ...
func (_m *ClientMock) Subscribe(email string, listID string) (interface{}, error) {
	ret := _m.Called(email, listID)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, string) interface{}); ok {
		r0 = rf(email, listID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, listID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
