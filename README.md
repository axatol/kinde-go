# kinde-go

Kinde golang client

## getting started

### prerequisites

1. Create a Machine to machine (M2M) application and save the domain, client ID, and client secret
2. Navigate to the Kinde Management API and save the audience
3. Authorise the M2M application with the Kinde Management API and allow the relevant scopes

### quickstart

```go
package main

import "github.com/axatol/kinde-go"

func main() {
  // load config from environment variables
  // - KINDE_DOMAIN
  // - KINDE_AUDIENCE
  // - KINDE_CLIENT_ID
  // - KINDE_CLIENT_SECRET
  client, _ := kinde.New(context.TODO(), nil)

  // or load from another source or directly
  client, _ = kinde.New(context.TODO(), &kinde.ClientOptions{
    Domain:       "https://example.kinde.com",
    Audience:     "https://example.kinde.com/api",
    ClientID:     "clientidclientidclientidclientid",
    ClientSecret: "clientsecretclientsecretclientsecretclientsecret",
  })

  apis, _ := client.GetAPIs(context.TODO(), nil)
}
```

## todo

- pagination
- rate-limiting
