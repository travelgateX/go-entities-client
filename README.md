# go-entities-client
GraphQL client to connect with the TravelgateX Entities API

* Entities API strong Go types for response data
* Build and execute any Entities API request
* Help package with some basic requests
* Options to log and debug transactions

## Installation
Make sure you have a working Go environment. To install go-entities-client, simply run:

```
$ go github.com/travelgateX/go-entities-client
```

## Initialization
The service endpoint can be provide or you can let the library choose the endpoint depending on the environment variable DEPLOY_MODE. Also you must provide a valid TravelgateX bearer with permission to manage the EntitiesAPI. There are two constructors:
```go
// Default client, endpoint chosen by DEPLOY_MODE environment variable
ent := entities.NewDefaultClient("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImt...")

// Entities client, endpoint provided
ent := entities.NewClient("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImt...", "https://api...")
```

## Usage
```go
// Execute Find Acceses by ID template
res, err := ent.Accesses(5)

// Execute new customized query (NewMutation for mutations)
res, err := ent.NewQuery(`
    query{
        admin{
            accesses(filter:{accessID:5}){
                edges{
                    node{
                        accessData{
                            name
                            code
                        }
                    }
                }
            }
        }
    }
`)

// Log errors
if err != nil {
    log.Fatal(err)
}

// Debug Mode to log transactions
// set 'true' before run a transaction
ent.DebugMode(true)
res, err := ent.Accesses(5)

```