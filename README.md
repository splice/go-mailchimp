[![Codeship Status for RichardKnop/go-mailchimp](https://codeship.com/projects/fdac3010-3acd-0134-e9c5-06456b66cf53/status?branch=master)](https://codeship.com/projects/166426)

# Mailchimp

A Golang SDK for Mailchimp API v3.

## Usage

```go
package main

import (
	"log"

	mailchimp "github.com/nhocki/go-mailchimp"
)

func main() {
	client, err := mailchimp.NewClient("the_api_key-us13", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the email is already subscribed
	memberResponse, err := client.CheckSubscription(
		"listID",
		"john@reese.com",
	)

	// User is already subscribed, update the subscription
	if err == nil {
		memberResponse, err = client.UpdateSubscription(
			"listID",
			"john@reese.com",
			"subscribed",
			map[string]interface{}{},
		)

		if err != nil {
			// Check the error response
			errResponse, ok := err.(*mailchimp.ErrorResponse)

			// Could not type assert error response
			if !ok {
				log.Fatal(err)
			}

			log.Fatal(errResponse)
		}

		log.Printf(
			"%s's subscription has been updated. Status: %s",
			memberResponse.EmailAddress,
			memberResponse.Status,
		)
		return
	}

	if err != nil {
		// Check the error response
		errResponse, ok := err.(*mailchimp.ErrorResponse)

		// Could not type assert error response
		if !ok {
			log.Fatal(err)
		}

		// 404 means we can process and subscribe user,
		// error other than 404 means we return error
		if errResponse.Status != http.StatusNotFound {
			log.Fatal(errResponse)
		}
	}

	// Subscribe the email
	memberResponse, err = client.Subscribe(
		"listID",
		"john@reese.com",
		map[string]interface{}{},
	)

	if err != nil {
		// Check the error response
		errResponse, ok := err.(*mailchimp.ErrorResponse)

		// Could not type assert error response
		if !ok {
			log.Fatal(err)
		}

		log.Fatal(errResponse)
	}

	log.Printf(
		"%s has been subscribed successfully. Status: %s",
		memberResponse.EmailAddress,
		memberResponse.Status,
	)
}
```
