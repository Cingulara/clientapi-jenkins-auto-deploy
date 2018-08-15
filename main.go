// main.go

package main

func main() {
	a := App{}
	a.Initialize("clientapi", "clientapi", "client-api-db-pipeline", "3306", "clientapi")

	a.Run(":8080")
}
