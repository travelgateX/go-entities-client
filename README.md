# go-entities-client
GraphQL Go client to connect with the TravelgateX Entities API

* Entities API strong Go types for response data
* Build and execute any Entities API request
* Help package with some basic requests
* Options to log and debug transactions

## Installation
Make sure you have a working Go environment. To install go-entities-client, simply run:

```
$ go get github.com/travelgateX/go-entities-client
```

## Initialization
The service endpoint can be provide or you can let the library choose the endpoint depending on the environment variable DEPLOY_MODE. Also you must provide a valid TravelgateX bearer with permission to manage the EntitiesAPI. There are two constructors:
```go
// Entities client, endpoint provided
ent := entities.NewClient("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImt...", "https://api...")

// Entities default client, endpoint chosen by DEPLOY_MODE environment variable
ent := entities.NewDefaultClient("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImt...")
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
ent.DebugMode(false)
res, err := ent.Accesses(5)

```

## Road Map
Interesting features to develop:

* Add unit tests
* Synchronization of the data model from the endpoint
* More request templates of the most used functions
* Possibility of obtaining raw response without fit the result with the data model

## Contribution
1. Fork ( https://github.com/[my-github-username]/go-entities-client/fork )
2. Create new feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push branch (`git push origin my-new-feature`)
5. Create new Pull Request