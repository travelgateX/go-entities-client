package entities

import "strings"

// accessesRQ returns graphql request query
func accessesRQ() string {
	return `
	  {
		admin {
		  accesses {
			edges {
			  node {
				accessData {
				  name
				  code
				  isActive
				  isTest
				  markets
				}
			  }
			}
		  }
		}
	  }	  
	`
}

// accessesRQ returns graphql request query
func accessesByGroupCodeRQ(grCode string) string {
	rq := `
	  {
		admin {
		  accesses(filter:{
			group:"$CODE$"
		  }) {
			edges {
			  node {
				accessData {
				  name
				  code
				  isActive
				  isTest
				  markets
				}
			  }
			}
		  }
		}
	  }	  
	`
	rq = strings.Replace(rq, "$CODE$", grCode, 1)
	return rq
}
