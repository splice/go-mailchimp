package mailchimp

import (
	"net/url"

	"github.com/stretchr/testify/mock"
)

// ClientMock is a mocked object implementing MailchimpInterface
type ClientMock struct {
	mock.Mock
}

// CheckSubscription ...
func (_m *ClientMock) CheckSubscription(listID string, email string) (*MemberResponse, error) {
	ret := _m.Called(listID, email)

	var r0 *MemberResponse
	if rf, ok := ret.Get(0).(func(string, string) *MemberResponse); ok {
		r0 = rf(listID, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*MemberResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(listID, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe ...
func (_m *ClientMock) Subscribe(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error) {
	ret := _m.Called(listID, email, mergeFields)

	var r0 *MemberResponse
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}) *MemberResponse); ok {
		r0 = rf(listID, email, mergeFields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*MemberResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, map[string]interface{}) error); ok {
		r1 = rf(listID, email, mergeFields)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSubscription ...
func (_m *ClientMock) UpdateSubscription(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error) {
	ret := _m.Called(listID, email, mergeFields)

	var r0 *MemberResponse
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}) *MemberResponse); ok {
		r0 = rf(listID, email, mergeFields)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*MemberResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, map[string]interface{}) error); ok {
		r1 = rf(listID, email, mergeFields)
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
