package entities

import (
	"strconv"
	"strings"

	"github.com/travelgateX/go-entities-client/model"
)

// grantAccessToGroupRQ returns graphql request mutation
func grantAccessToGroupRQ(id int, groups []string) string {
	rq := `
		mutation{
			admin{
				grantAccessToGroup(input:{
					id:$ID$
					groups:$GROUPS$
				}){
					code
					error{
						code
						type
						description
					}
				}
			}
	  	}    
	`
	rq = strings.Replace(rq, "$ID$", strconv.Itoa(id), 1)
	rq = strings.Replace(rq, "$GROUPS$", sliceToQuotedStringFormat(groups), 1)
	return rq
}

// grantSupplierToGroupRQ returns graphql request mutation
func grantSupplierToGroupRQ(id string, groups []string) string {
	rq := `
	  	mutation{
			admin{
		  		grantSupplierToGroup(input:{
					id:"$ID$"
					groups:$GROUPS$
		  		}){
					code
					error{
						code
						type
						description
					}
		  		}
			}
	  	}	    
	`
	rq = strings.Replace(rq, "$ID$", id, 1)
	rq = strings.Replace(rq, "$GROUPS$", sliceToQuotedStringFormat(groups), 1)
	return rq
}

// deleteAccessFromGroupRQ returns graphql request mutation
func deleteAccessFromGroupRQ(id int, groups []string) string {
	rq := `
		mutation{
			admin{
				deleteAccessFromGroup(input:{
					id:$ID$
					groups:$GROUPS$
				}){
					code
					error{
						code
						type
						description
					}
				}
			}
	  	}    
	`
	rq = strings.Replace(rq, "$ID$", strconv.Itoa(id), 1)
	rq = strings.Replace(rq, "$GROUPS$", sliceToQuotedStringFormat(groups), 1)
	return rq
}

// deleteSupplierFromGroupRQ returns graphql request mutation
func deleteSupplierFromGroupRQ(id string, groups []string) string {
	rq := `
	  	mutation{
			admin{
		  		deleteSupplierFromGroup(input:{
					id:"$ID$"
					groups:$GROUPS$
		  		}){
					code
					error{
						code
						type
						description
					}
		  		}
			}
	  	}	    
	`
	rq = strings.Replace(rq, "$ID$", id, 1)
	rq = strings.Replace(rq, "$GROUPS$", sliceToQuotedStringFormat(groups), 1)
	return rq
}

// createClient returns graphql request mutation
func createDefaultClient(input model.CreateClientInput) string {
	rq := `
		mutation{
			admin{
				createClient(input:{
					name: "$NAME$"
					isActive: "$ACTIVE$"
					group: "$GROUP$"
				}){
					code
					error{
						code
						type
						description
					}
				
				}
			}
		}
	`
	rq = strings.Replace(rq, "$NAME$", input.Name, 1)
	rq = strings.Replace(rq, "$ACTIVE$", strconv.FormatBool(input.IsActive), 1)
	rq = strings.Replace(rq, "$GROUP$", input.Group, 1)
	return rq
}
