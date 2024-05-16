package client
import(
	"net/http"
	"net/url"
	"time"
	"log"
	"context"
	"io/ioutil"
)
type Item struct{
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
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
	reqURL := *c.BaseURL
	req,err := http.NewRequest(http.MethodGet,reqURL.String(),nil)
	if err != nil{
		return nil,err
	}
	req.Header.Set("User-Agent","groq-go-client")
	req = req.WithContext(ctx)
	resp,err := c.HTTPClient.Do(req)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	switch resp.StatusCode{
	case http.StatusOK:
		bodyBytes,err := ioutil.ReadAll(resp.Body)
		if err != nil{
			return nil,err
		}
		var items []*Item
		if err := json.Unmarshal(bodyBytes,&items); err != nil{
			return nil,err
		}
		return items,nil
	case http.StatusBadRequest:
		return nil,errors.New("bad request")
	case http.StatusNotFound:
		return nil,errors.New("where are you?")
	default:
		return nil,errors.New("highly illogical")
	}
}
