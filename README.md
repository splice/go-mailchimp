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
  client := mailchimp.NewClient("apixyc-us11", nil)
  _, err := client.Subscribe("john@doe.com", "listidxyz")
  if err != nil {
    log.Fatal(err)
  }
}
```
