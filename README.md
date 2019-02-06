# go-github-webhook

A GitHub webhook handler written in Go. It simplifies payload validation, etc ...

## Sample app

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/GitbookIO/go-github-webhook"
)

func main() {
    // Your GitHub secret (this could also be dynamic), if you're secret is empty it will skip validation
    secret := ""

    if err := http.ListenAndServe(":8000", WebhookLog(secret)); err != nil {
        fmt.Errorf("Error: %s", err)
    }
}

func WebhookLog(secret string) http.Handler {
    return github.Handler(secret, func(event string, payload *github.GitHubPayload, req *http.Request) error {
        // Log webhook
        fmt.Println("Received", event, "for ", payload.Repository.Name)

        // You'll probably want to do some real processing
        fmt.Println("Can clone repo at:", payload.Repository.CloneURL)

        // All is good (return an error to fail)
        return nil
    })
}
```
