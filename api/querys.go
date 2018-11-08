package api

import (
	"strconv"
	"strings"
)

// accessesRQ returns graphql request query
func accessesRQ(id int) string {
	rq := `
		query{
			admin{
				accesses(filter:{
					accessID:$ID$
				}){
					edges{
						node{
							accessData{
								name
								code
								supplier{
									supplierData{
										code
										name
									}
								}
							}
						}
					}
				}
			}
		}
	`
	return strings.Replace(rq, "$ID$", strconv.Itoa(id), 1)
}
