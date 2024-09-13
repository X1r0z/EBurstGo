package common

import (
	"crypto/tls"
	"github.com/X1r0z/go-ntlmssp"
	"net/http"
	"net/url"
)

func NewHttpClient(info *TaskInfo) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Renegotiation:      tls.RenegotiateOnceAsClient,
				MinVersion:         tls.VersionTLS10,
			},
			Proxy: func(_ *http.Request) (*url.URL, error) {
				if info.Proxy != "" {
					return url.Parse(info.Proxy)
				} else {
					return nil, nil
				}
			},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func NewNtlmClient(info *TaskInfo) *http.Client {
	return &http.Client{
		Transport: ntlmssp.Negotiator{
			RoundTripper: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					Renegotiation:      tls.RenegotiateOnceAsClient,
					MinVersion:         tls.VersionTLS10,
				},
				Proxy: func(_ *http.Request) (*url.URL, error) {
					if info.Proxy != "" {
						return url.Parse(info.Proxy)
					} else {
						return nil, nil
					}
				},
			},
			Pth: info.PthMode,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}
