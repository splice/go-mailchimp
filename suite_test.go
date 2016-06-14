package mailchimp

var notFoundErrorResponse = `{
    "type": "http://developer.mailchimp.com/documentation/mailchimp/guides/error-glossary/",
    "title": "Resource Not Found",
    "status": 404,
    "detail": "The requested resource could not be found.",
    "instance": ""
}`

var alreadySubscribedErrorResponse = `{
    "type": "http://developer.mailchimp.com/documentation/mailchimp/guides/error-glossary/",
    "title": "Member Exists",
    "status": 400,
    "detail": " is already a list member. Use PUT to insert or update list members.",
    "instance": ""
}`

var invalidMergeFieldsErrorResponse = `{
    "type": "http://developer.mailchimp.com/documentation/mailchimp/guides/error-glossary/",
    "title": "Invalid Resource",
    "status": 400,
    "detail": "Your merge fields were invalid.",
    "instance": ""
}`

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
