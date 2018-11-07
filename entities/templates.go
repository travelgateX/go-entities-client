package entities

import "strings"

// GrantAccessToGroupRequest :
func GrantAccessToGroupRequest(ID, gr string) string {
	res := strings.Replace(grantAccessToGroup01, "#ID#", ID, 1)
	return strings.Replace(res, "#GROUP#", gr, 1)
}

// DeleteAccessFromGroupRequest :
func DeleteAccessFromGroupRequest(ID, gr string) string {
	res := strings.Replace(deleteAccessFromGroup01, "#ID#", ID, 1)
	return strings.Replace(res, "#GROUP#", gr, 1)
}

// grantAccessToGroup01
const grantAccessToGroup01 = `
	mutation{
		admin{
			grantAccessToGroup(input:{
				id:#ID#
				groups:["#GROUP#"]
			}){
				accessData{
					name
					code
					isActive
					isTest
					isSchedulerActive
					owner{
						code
					}
					supplier{
						supplierData{
							code
							name
						}
					}
				}
				error{
					code
					type
					description
				}
			}
		}
	}
`

const deleteAccessFromGroup01 = `
	mutation {
		admin {
			deleteAccessFromGroup(input: {id: #ID#, groups: ["#GROUP#"]}) {
				accessData {
					name
					code
					isActive
					isTest
					isSchedulerActive
					owner {
						code
					}
					supplier {
						supplierData {
							code
							name
						}
					}
				}
				error {
					code
					type
					description
				}
			}
		}
	}
`
