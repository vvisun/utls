package algoutil

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/vvisun/utls/leaflog"
)

// http跨域
func AccessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Token")
		h.ServeHTTP(w, r)
	})
}

func OptionControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			json.NewEncoder(w).Encode(`{ "code": 0, "data": "success"}`)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// http's get method
func HTTPGet(apiUrl string, data url.Values) ([]byte, error) {
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		//leaflog.Debug("get failed, err:%v\n", err)
		return nil, err
	}
	u.RawQuery = data.Encode() // URL encode

	var body []byte
	rspn, err := http.Get(u.String())
	//leaflog.Debug("get url:%s\n", u.String())
	if err != nil {
		//leaflog.Debug("get failed, err:%v\n", err)
		return nil, err
	}
	defer rspn.Body.Close()
	body, err = io.ReadAll(rspn.Body)
	if err != nil {
		//leaflog.Debug("get failed, err:%v\n", err)
		return nil, err
	}
	return body, err
}

func HTTPPost(apiUrl string, data url.Values, param interface{}) ([]byte, error) {
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		leaflog.Debug("post failed, err:%v\n", err)
		return nil, err
	}
	u.RawQuery = data.Encode() // URL encode

	info, err := json.Marshal(param)
	if err != nil {
		leaflog.Debug("post failed, err:%v\n", err)
		return nil, err
	}

	var body []byte
	rspn, err := http.Post(u.String(), "text/plain", strings.NewReader(string(info)))
	leaflog.Debug("post url:%s param:%s\n", u.String(), string(info))
	if err != nil {
		leaflog.Debug("post failed, err:%v\n", err)
		return nil, err
	}
	defer rspn.Body.Close()
	body, err = io.ReadAll(rspn.Body)
	if err != nil {
		leaflog.Debug("post failed, err:%v\n", err)
		return nil, err
	}
	return body, err
}
