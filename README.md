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
