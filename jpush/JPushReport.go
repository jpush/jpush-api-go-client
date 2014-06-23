package jpush

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type JPushReportClient struct {
	appKey       string
	masterSecret string
}

type JPushReportObject struct {
	Android_received int
	Ios_apns_sent    int
	Msg_id           int
}

func NewJPushReportClient(appkey, master_secret string) *JPushReportClient {
	return &JPushReportClient{appkey, master_secret}
}

func (vjprc *JPushReportClient) GetReport(msg_id ...int) (string, error) {
	if len(vjprc.appKey) != KEY_LENGTH || len(vjprc.masterSecret) != KEY_LENGTH {
		return "", fmt.Errorf("invalidate appkey/masterSecret")
	}

	b, e := json.Marshal(msg_id)
	if e != nil {
		return "", e
	}

	reqEndpoint := "https://report.jpush.cn"
	reqUrl := fmt.Sprintf("%s/v2/received?msg_ids=%s", reqEndpoint, string(b[1:len(b)-1]))
	req, e := http.NewRequest("GET", reqUrl, nil)
	if e != nil {
		return "NewRequestReport", e
	}

	req.Header.Set("User-Agent", "JPush-API-GO-Client")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept-Charset", "UTF-8")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(vjprc.appKey, vjprc.masterSecret)

	c := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*CONN_TIMEOUT)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(RW_TIMEOUT * time.Second))
				return c, nil
			},
		},
	}

	resp, e := c.Do(req)
	defer resp.Body.Close()

	if e != nil {
		return resp.Status, e
	}

	buf := make([]byte, 4096)
	nr, _ := resp.Body.Read(buf)

	if resp.StatusCode == http.StatusOK {
		return string(buf[:nr]), nil
	}

	return resp.Status, fmt.Errorf(string(buf[:nr]))
}

func (vjprc *JPushReportClient) GetReportObject(msg_ids ...int) ([]JPushReportObject, error) {
	report, err := vjprc.GetReport(msg_ids...)
	if err != nil {
		return nil, err
	}

	var res []JPushReportObject
	if e := json.Unmarshal([]byte(report), &res); e != nil {
		return nil, e
	}

	return res, nil
}
