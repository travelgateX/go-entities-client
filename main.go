package main

import (
	"log"

	"github.com/travelgateX/go-entities-client/entities"
)

func main() {
	log.Println("Hello entities!")

	c := entities.NewDefaultClient("Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ik5EQTFOa00zUVVSRk1EQXdOekF4TVRKRk5UQkdSRGMwUWpreVEwVTJOVVV6UkRrNU5rUkNSZyJ9.eyJodHRwczovL3RyYXZlbGdhdGV4LmNvbS9pYW0iOlt7ImEiOm51bGwsImMiOiJ4dGciLCJnIjpbeyJhIjp7ImFsbCI6eyJlbnRpdHkiOnsiYWNjIjpbImNydWQxYWYiXSwiY2xpIjpbImNydWQxYWYiXSwic3VwIjpbImNydWQxYWYiXX0sImhvdGxzdCI6eyJob3QiOlsiMXgiXX0sImh1YmdyYSI6eyJib2siOlsiMXhmIl0sImJvbyI6WyIxeGYiXSwiYnJkIjpbIjF4ZiJdLCJjYXQiOlsiMXhmIl0sImNmZyI6WyIxeGYiXSwiY25sIjpbIjF4ZiJdLCJxdGUiOlsiMXhmIl0sInJvbSI6WyIxeGYiXSwic3JjIjpbIjF4ZiJdfSwiaWFtIjp7ImFwaSI6WyJjcnVkMWFmIl0sImdycCI6WyJjcnVkMWFmIl0sIm1iciI6WyJjcnVkMWFmIl0sInByZCI6WyJjcnVkMWFmIl0sInJvbCI6WyJjcnVkMWFmIl0sInJzYyI6WyJjcnVkMWFmIl0sIm9wdCI6WyJjcnVkMWFmIl19fX0sImMiOiJkZXZvcHMiLCJnIjpudWxsLCJwIjp7ImlhbSI6eyJncnAiOlsiY3J1ZDEiXX19LCJ0IjoidGVhbSJ9XSwicCI6eyJpYW0iOnsiZ3JwIjpbImNydWQxIl19fSwidCI6InJvb3QifV0sImh0dHBzOi8vdHJhdmVsZ2F0ZXguY29tL21lbWJlcl9pZCI6ImNndXptYW5AeG1sdHJhdmVsZ2F0ZS5jb20iLCJpc3MiOiJodHRwczovL3h0Zy1kZXYuZXUuYXV0aDAuY29tLyIsInN1YiI6IkNzM1prTmpwTUN3MmZWVDlnNnBuMGxzU0p3ZDYxUDlYQGNsaWVudHMiLCJhdWQiOiJodHRwczovL2FwaS50cmF2ZWxnYXRleC5jb20iLCJpYXQiOjE1Mzg5ODM0MTQsImV4cCI6MTU0Njc1OTQxNCwiYXpwIjoiQ3MzWmtOanBNQ3cyZlZUOWc2cG4wbHNTSndkNjFQOVgiLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMifQ.NrrLo8-qtlZdyB8Z2yxy3Kal29RbCW25dpdQtTOD0xyPVys-OaVvJcz8mGgxKSwCtwOS7VY9Glvupo5lqkvIwLivPPB3_kfMAkcj_aAd1jM_ZCjyDI2aFpWh1usJnKOwzkNLelh938-wFt4wdv9US_awynOWEInCigX2DfIZrGxrkmN4g_pqnbsLL0jhK0EDceBtvc_kVtZ9jlOsV1L9qfRIf98V3HqhN3ZKCJRll5ewmC4R1UMnavCHjJ9xjr6PsegideWKGIFHM-BdChPiN5yJUvtdLk84t5OAvYjaRXnGnQ_5qYpsANt5Ll4UvejIUFdkEX2fJz8AeqD3O8zodw")
	res := c.NewRequest(`
		query{
			admin{
			accesses(filter:{
				accessID:3
			}){
				edges{
					node{
						accessData{
							code
							name
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
	`)

	//log.Printf("res=%v", res)
	log.Printf("res=%v", res.Admin.Accesses.Edges[0].Node.AccessData.Name)
}
