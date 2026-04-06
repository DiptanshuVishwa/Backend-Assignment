package main
import("os"
	"log"
"backend-assignment/src")

func main(){
	app := src.SetupApp()
	port := os.Getenv("PORT")
	if port == ""{port = "3000"}
	log.Println("Server is running on port" +port)
	app.Listen(port)
}