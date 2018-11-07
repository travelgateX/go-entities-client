package entities

import (
	"fmt"
	"strconv"
	"strings"
)

// GrantAccessToGroupRQ returns graphql request
func GrantAccessToGroupRQ(id int, groups []string) string {
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

	var gr []string
	for _, g := range groups {
		gr = append(gr, strconv.Quote(g))
	}

	rq = strings.Replace(rq, "$ID$", strconv.Itoa(id), 1)
	rq = strings.Replace(rq, "$GROUPS$", fmt.Sprintf("%v", gr), 1)

	return rq
}
