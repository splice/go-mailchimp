package mailchimp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AreaHQ/mailchimp/status"
	"github.com/stretchr/testify/assert"
)

var alreadySubscribedErrorResponse = `{
    "type": "http://developer.mailchimp.com/documentation/mailchimp/guides/error-glossary/",
    "title": "Member Exists",
    "status": 400,
    "detail": " is already a list member. Use PUT to insert or update list members.",
    "instance": ""
}`

func TestSubscribeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(400)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, alreadySubscribedErrorResponse)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, err := NewClient("the_api_key-us13", &http.Client{Transport: transport})
	assert.NoError(t, err)

	baseURL, _ := url.Parse("http://localhost/")
	client.SetBaseURL(baseURL)

	memberResponse, err := client.Subscribe("john@reese.com", "list_id")
	assert.Nil(t, memberResponse)
	assert.Equal(t, "Error 400 Member Exists ( is already a list member. Use PUT to insert or update list members.)", err.Error())

	errResponse, ok := err.(*ErrorResponse)
	assert.True(t, ok)
	assert.Equal(t, "Member Exists", errResponse.Title)
	assert.Equal(t, 400, errResponse.Status)
	assert.Equal(t, " is already a list member. Use PUT to insert or update list members.", errResponse.Detail)
}

func TestSubscribeMalformedError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(500)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, err := NewClient("the_api_key-us13", &http.Client{Transport: transport})
	assert.NoError(t, err)

	baseURL, _ := url.Parse("http://localhost/")
	client.SetBaseURL(baseURL)

	memberResponse, err := client.Subscribe("john@reese.com", "list_id")
	assert.Nil(t, memberResponse)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}

var successResponse = `{
    "id": "11bf13d1eb58116eba1de370b2bd796b",
    "email_address": "john@reese.com",
    "unique_email_id": "1b757e82a3",
    "email_type": "html",
    "status": "subscribed",
    "merge_fields": {
        "FNAME": "",
        "LNAME": "",
        "MMERGE4": "",
        "MMERGE5": "",
        "MMERGE3": ""
    },
    "stats": {
        "avg_open_rate": 0,
        "avg_click_rate": 0
    },
    "ip_signup": "",
    "timestamp_signup": "",
    "ip_opt": "101.8.90.86",
    "timestamp_opt": "2016-06-03T07:13:07+00:00",
    "member_rating": 2,
    "last_changed": "2016-06-03T07:13:07+00:00",
    "language": "",
    "vip": false,
    "email_client": "",
    "location": {
        "latitude": 0,
        "longitude": 0,
        "gmtoff": 0,
        "dstoff": 0,
        "country_code": "",
        "timezone": ""
    },
    "list_id": "0f6b836652",
    "_links": [
        {
            "rel": "self",
            "href": "https://us13.api.mailchimp.com/3.0/lists/0f6b836652/members/11bf13d1eb58116eba1de370b2bd796b",
            "method": "GET",
            "targetSchema": "https://us13.api.mailchimp.com/schema/3.0/Lists/Members/Instance.json"
        },
        {
            "rel": "parent",
            "href": "https://us13.api.mailchimp.com/3.0/lists/0f6b836652/members",
            "method": "GET",
            "targetSchema": "https://us13.api.mailchimp.com/schema/3.0/Lists/Members/Collection.json",
            "schema": "https://us13.api.mailchimp.com/schema/3.0/CollectionLinks/Lists/Members.json"
        },
        {
            "rel": "update",
            "href": "https://us13.api.mailchimp.com/3.0/lists/0f6b836652/members/11bf13d1eb58116eba1de370b2bd796b",
            "method": "PATCH",
            "schema": "https://us13.api.mailchimp.com/schema/3.0/Lists/Members/Instance.json"
        },
        {
            "rel": "upsert",
            "href": "https://us13.api.mailchimp.com/3.0/lists/0f6b836652/members/11bf13d1eb58116eba1de370b2bd796b",
            "method": "PUT",
            "schema": "https://us13.api.mailchimp.com/schema/3.0/Lists/Members/Instance.json"
        },
        {
            "rel": "delete",
            "href": "https://us13.api.mailchimp.com/3.0/lists/0f6b836652/members/11bf13d1eb58116eba1de370b2bd796b",
            "method": "DELETE"
        },
        {
            "rel": "activity",
            "href": "https://us13.api.mailchimp.com/3.0/lists/0f6b836652/members/11bf13d1eb58116eba1de370b2bd796b/activity",
            "method": "GET",
            "targetSchema": "https://us13.api.mailchimp.com/schema/3.0/Lists/Members/Activity/Collection.json"
        },
        {
            "rel": "goals",
            "href": "https://us13.api.mailchimp.com/3.0/lists/0f6b836652/members/11bf13d1eb58116eba1de370b2bd796b/goals",
            "method": "GET",
            "targetSchema": "https://us13.api.mailchimp.com/schema/3.0/Lists/Members/Goals/Collection.json"
        },
        {
            "rel": "notes",
            "href": "https://us13.api.mailchimp.com/3.0/lists/0f6b836652/members/11bf13d1eb58116eba1de370b2bd796b/notes",
            "method": "GET",
            "targetSchema": "https://us13.api.mailchimp.com/schema/3.0/Lists/Members/Notes/Collection.json"
        }
    ]
}`

func TestSubscribe(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, successResponse)
	}))
	defer server.Close()

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client, err := NewClient("the_api_key-us13", &http.Client{Transport: transport})
	assert.NoError(t, err)

	baseURL, _ := url.Parse("http://localhost/")
	client.SetBaseURL(baseURL)

	memberResponse, err := client.Subscribe("john@reese.com", "list_id")
	assert.NoError(t, err)

	assert.Equal(t, "11bf13d1eb58116eba1de370b2bd796b", memberResponse.ID)
	assert.Equal(t, "john@reese.com", memberResponse.EmailAddress)
	assert.Equal(t, status.Subscribed, memberResponse.Status)
}
