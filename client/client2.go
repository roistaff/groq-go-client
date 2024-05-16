package client
import(
	"fmt"
	"net/http"
	"net/url"
	"errors"
)
type Client struct{
	URL	*url.URL
	HTTPClient	*http.Client
	APIToken	string
}
func NewClient(urlStr,token string)(*Client,error){
	if len(token) == 0 {
		return nil,errors.New("misssing token")
	}
	parsedURL,err := url.ParseRequestURL(urlStr)
	if err != nil {
		return nil, erros.Wrapf(err, "fsiled to parse url: %s",urlStr)
	}

}
func Print(){
	fmt.Println("HELLO")
}
