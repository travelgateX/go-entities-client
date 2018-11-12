package entities

import (
	"strconv"
	"strings"
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
