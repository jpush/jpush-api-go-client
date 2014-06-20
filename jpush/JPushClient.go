package jpush

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	KEY_LENGTH   = 24
	CONN_TIMEOUT = 5
	RW_TIMEOUT   = 30
)

type JPushClient struct {
	appKey       string
	masterSecret string
}

func NewJPushClient(appkey, master_secret string) *JPushClient {
	return &JPushClient{appkey, master_secret}
}

func (vjpc *JPushClient) Push(po *PushPayload) (string, error) {
	if len(vjpc.appKey) != KEY_LENGTH || len(vjpc.masterSecret) != KEY_LENGTH {
		return "", fmt.Errorf("invalidate appkey/masterSecret")
	}

	postr, e := po.ToJsonString()
	if e != nil {
		return postr, e
	}

	body := strings.NewReader(postr)
	req, e := http.NewRequest("POST", "https://api.jpush.cn/v3/push", body)
	if e != nil {
		return "NewRequest", e
	}

	req.Header.Set("User-Agent", "JPush-API-GO-Client")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept-Charset", "UTF-8")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(vjpc.appKey, vjpc.masterSecret)

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
