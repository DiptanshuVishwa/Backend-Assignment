package main
import("os"
	"log"
"github.com/DiptanshuVishwa/Backend-Assignment/src")

func main(){
	app := src.SetupApp()
	port := os.Getenv("PORT")
	if port == ""{port = "3000"}
	log.Println("Server is running on port" +port)
	app.Listen(":"+ port)
}