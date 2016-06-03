package mailchimp

import (
	"net/url"

	"github.com/stretchr/testify/mock"
)

// ClientMock is a mocked object implementing MailchimpInterface
type ClientMock struct {
	mock.Mock
}

// Subscribe ...
func (_m *ClientMock) Subscribe(email string, listID string) (*MemberResponse, error) {
	ret := _m.Called(email, listID)

	var r0 *MemberResponse
	if rf, ok := ret.Get(0).(func(string, string) *MemberResponse); ok {
		r0 = rf(email, listID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*MemberResponse)
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

// SetBaseURL ...
func (_m *ClientMock) SetBaseURL(baseURL *url.URL) {
	_m.Called(baseURL)
}

// GetBaseURL ...
func (_m *ClientMock) GetBaseURL() *url.URL {
	ret := _m.Called()

	var r0 *url.URL
	if rf, ok := ret.Get(0).(func() *url.URL); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	return r0
}
