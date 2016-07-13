package ndnetapi

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"errors"
	"net/http"
	"strings"
	"time"
)

var (
	AN = "ndnetapi "
)

type Client struct {
	Name		string
	Endpoint	string
	Config          *Config // for debug only
}

type Config struct {
	Name		string
	BridgePrefix	string
	NedgeHost	string
	NedgePort	int
}

func ReadParseConfig(fname string) (Config, error) {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal("Error reading config file: ", fname, " error: ", err)
	}
	var conf Config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		log.Fatal("Error parsing config file: ", fname, " error: ", err)
	}
	return conf, nil
}

func ClientAlloc(configFile string) (c *Client, err error) {
	conf, err := ReadParseConfig(configFile)
	if err != nil {
		log.Fatal("Error initializing client from Config file: ", configFile, " error: ", err)
	}

	NedgeClient := &Client{
		Name: conf.Name,
		Endpoint: "http://" + conf.NedgeHost + ":" + string(conf.NedgePort) + "/",
		Config:	&conf,
	}

	return NedgeClient, nil
}

func (c *Client) Request(method, endpoint string, data map[string]interface{}) (body []byte, err error) {
	log.Debug("REST request to Nexenta, endpoint: ", endpoint, " data: ", data, " method: ", method)
	if c.Endpoint == "" {
		log.Error("Endpoint is not set, unable to issue requests")
		err = errors.New("Unable to issue json-rpc requests without specifying Endpoint")
		return nil, err
	}
	datajson, err := json.Marshal(data)
	if (err != nil) {
		log.Error(err)
	}

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	url := c.Endpoint + endpoint
	req, err := http.NewRequest(method, url, nil)
	if len(data) != 0 {
		req, err = http.NewRequest(method, url, strings.NewReader(string(datajson)))
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error while handling request", err)
		return nil, err
	}
	c.checkError(resp)
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if (err != nil) {
		log.Error(err)
	}
	if (resp.StatusCode == 202) {
		body, err = c.resend202(body)
	}
	return body, err
}

func (c *Client) resend202(body []byte) ([]byte, error) {
	time.Sleep(1000 * time.Millisecond)
	r := make(map[string][]map[string]string)
	err := json.Unmarshal(body, &r)
	if (err != nil) {
		log.Error(err)
	}

	url := c.Endpoint + r["links"][0]["href"]
	resp, err := http.Get(url)
	if err != nil {
		log.Error("Error while handling request", err)
		return nil, err
	}
	defer resp.Body.Close()
	c.checkError(resp)

	if resp.StatusCode == 202 {
		body, err = c.resend202(body)
	}
	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}

func (c *Client) checkError(resp *http.Response) (bool) {
	log.Debug("REST returned error code: ", resp.StatusCode)
	if resp.StatusCode > 399 {
		body, err := ioutil.ReadAll(resp.Body)
		log.Error(resp.StatusCode, string(body), err)
		return true
	}
	return false
}


func (c *Client) CreateNetwork(name string) (err error) {
	log.Debug(AN, "Creating network %s", name)
	err = nil
	return err
}

func (c *Client) DeleteNetwork(name string) (err error) {
	log.Debug(AN, "Deleting Volume ", name)
	err = nil
	return err
}

func (c *Client) ListNetworks() (nlist []string, err error) {
	log.Debug(AN, "ListNetworks ")
	err = nil
	return nlist, err
}
