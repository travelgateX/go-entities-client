package entities

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
