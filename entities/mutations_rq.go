package entities

import (
	"strings"
)

// grantAccessToGroupRQ returns graphql request mutation
func grantAccessToGroupRQ(code string, groups []string) string {
	rq := `
	  mutation{
			admin{
				grantAccessToGroup(input:{
					id:$CODE$
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
	rq = strings.Replace(rq, "$CODE$", code, 1)
	rq = strings.Replace(rq, "$GROUPS$", sliceToQuotedStringFormat(groups), 1)
	return rq
}
