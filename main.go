package main
import (
	"github.com/roistaff/groq-go-client/client"
	"os"
	"context"
)
func main(){
	groq,err := client.New("https://api.groq.com/openai/v1/chat/completions",os.Getenv("GROQ_API_KEY"), nil)
	if err != nil{
		panic(err)
	}
}
items,err := Client.Chat(context.Background(),)
