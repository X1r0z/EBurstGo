package lib

import (
	"crypto/tls"
	"github.com/Azure/go-ntlmssp"
	"net/http"
)

var ExchangeUrls = map[string]string{
	"autodiscover": "/autodiscover",
	"ews":          "/ews",
	"mapi":         "/mapi",
	"activesync":   "/Microsoft-Server-ActiveSync",
	"oab":          "/oab", // TODO
	"rpc":          "/rpc",
	"owa":          "/owa/auth.owa",
	"powershell":   "/powershell",
	"ecp":          "/owa/auth.owa",
}

var Client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			Renegotiation:      tls.RenegotiateOnceAsClient,
		},
	},
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

var NtlmClient = &http.Client{
	Transport: ntlmssp.Negotiator{
		RoundTripper: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Renegotiation:      tls.RenegotiateOnceAsClient,
			},
		},
	},
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

var Log *Logging
