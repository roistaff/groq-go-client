package client
import(
	"net/http"
	"net/url"
	"time"
	"log"
	"context"
)
type Chat struct{

}

type Client struct {
	BaseURL	*url.URL
	HTTPClient *http.Client
	Token	string
	Logger	*log.Logger
}
func New(rawBaseURL,token string,logger *log.Logger)(*Client,error){
	baseURL,err := url.Parse(rawBaseURL)
	if err != nil{
		return nil,err
	}
	if logger == nil{
		logger = log.New(os.Stderr,"[LOG]",log.LstdFlags)
	}
	return &Client{
		BaseURL:	baseURL,
		HTTPClient:	http.DefaultClient,
		Token:	token,
		Logger:	logger,
	},nil
}
func(c *Client)GetChat(ctx context.Context,roleSystem,roleUser,string){

	return nil,nil
}
