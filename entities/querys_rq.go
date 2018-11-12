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

// suppliersByGroupCodeRQ returns graphql request query
func suppliersRQ() string {
	return `
        query{
            admin{
                suppliers{
                    edges{
                        node{
                            supplierData{
                                code
                                name
                                isActive
                                owner{
                                    code
                                }
                            }
                        }
                    }
                }
            }
        }
	`
}

// suppliersByGroupCodeRQ returns graphql request query
func suppliersByGroupCodeRQ(grCode string) string {
	rq := `
        query{
            admin{
                suppliers(filter:{
                    groupID:"$CODE$"
                }){
                    edges{
                        node{
                            supplierData{
                                code
                                name
                                isActive
                                owner{
                                    code
                                }
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
