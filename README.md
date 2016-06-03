[![Codeship Status for AreaHQ/mailchimp](https://codeship.com/projects/7252c9a0-09f0-0134-e361-2adbeb910e90/status?branch=master)](https://codeship.com/projects/155402)

# Mailchimp

A golang SDK for Mailchimp API v3.

## Usage

```go
package main

import (
  "log"

  "github.com/AreaHQ/mailchimp"
)

func main() {
  client, err := mailchimp.NewClient("the_api_key-us13", nil)
  if err != nil {
    log.Fatal(err)
  }

  memberResponse, err := client.CheckSubscription("john@reese.com", "list_id")
	if err != nil {
    errResponse, ok := err.(*mailchimp.ErrorResponse)
    if !ok {
      log.Fatal(err)
    }
    // errResponse.Type
    // errResponse.Title
    // errResponse.Status
    // errResponse.Detail
		log.Fatal(errResponse)
	}

  memberResponse, err = client.Subscribe("john@reese.com", "list_id")
  if err != nil {
    errResponse, ok := err.(*mailchimp.ErrorResponse)
    if !ok {
      log.Fatal(err)
    }
    // errResponse.Type
    // errResponse.Title
    // errResponse.Status
    // errResponse.Detail
		log.Fatal(errResponse)
	}
}
```
