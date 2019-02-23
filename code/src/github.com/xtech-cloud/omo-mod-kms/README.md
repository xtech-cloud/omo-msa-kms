# 安装

```
go get github.com/xtech-cloud/omo-mod-kms
```

# 使用

```go

package main

import (
    kms "github.com/xtech-cloud/omo-mod-kms"
)

func main() {
    //创建一个应用
    appkey,appsecret, pubkey, prikey, err := kms.CreateApp("omo")

    devicecode := "9AFE1346C68974C634960F7F4B876271"

    //生成一个无自定义数据的永久授权文件
    license0, _ := kms.MakeLicense(appkey, appsecrect, devicecode, "", 0, pubkey, prikey)

    //生成一个含自定义数据的90天授权文件
    license90, _ := kms.MakeLicense(appkey, appsecrect, devicecode, "{\"app\":\"omo\"}", 90, pubkey, prikey)

    //验证授权文件
    code, _ := kms.VerifyLicense(license90, appkey, appsecret, devicecode)
}
```

# 授权文件格式说明

```
key:
673b7d576ea8e6ae577f162274f084ec
code:
9AFE1346C68974C634960F7F4B876271
timestamp:
1544668747
expiry:
90
storage:
{"company":"omo"}
cer:
HbkuFbfAvYQaMeQlFEkFW9-aC3v_U-VSBkt5yjg0KXLCV1q1OIMuvzwIU5M_v00NVtaOHDOa5mIs0HDPvnsrXCaTovwQKVmXAns0IZqSw8meul295xyWiS-XqFBKJkY4aRz0M8m2GNvUVGy2J03lg5QM7G9G0_IfOirkK8bWw6m0nZ779U1Jw1oypNliRQlJ1DzKIGE6raoCXPosJ7S8EV-NQaG0Tc26M_GZR-ikT8JkqvMp3kpLDBm2gO3zycgyiIfHBhFlYb54XEkY-7onP1vsA88OzQF4g2aV_twQ_zpGwkuLjTUuxLl9LnKkUAYlhekH17Ihsc_Q7TVa0zd-PL_kgjhIDmJf0aPU-hyvEDDqdodPCJEHh0ZhvQMvWhGOYEF6lnIFwQP5H52_73zVWLd303uO6q1QQzDtoQZmFC7arP9mxyCo4_7SezeDqYYnNXNIOu5PcFf70GdKN0E3U1xnI34RkMr-iMqE1WZiYGmPVnHWb93NaA_D1PvIODLGoyoniq6ev2R3avLe5JY0_YRddMb0Q_xYyyi6iB-Dqi41v6qeU_DD6H_dx3T6styO1C6KBAD30OMTOEikUYsRMT_1IHzHon9QnixW12IUvyI=
sig:
pB4_s8FggA0M-9CxD7mYfQ8oQC-oLARjpZreWvgO5kJGskd-huAQxPMbArZdZ6xQ58DjwWtIeyAgrBdpTXYI-H9gSVJrW94cDPWV4ND-i4B0kCFsavqzWIbOAXGmhWralQZL3ozDp7et4QFYs327upuQf-reNNj21a8_1ZQcvZdF-hTfr3hY7YQ5D3QHDnBHpWMvxAYcvBwuDbwIxIkXN0wkfNBRHYdXwQnOi1LK548YvUT4CX4liqgpjWf1HmEVEgGfODvxzO9KG5SKMeOtEducyKPvmrqj4rzMp3fpZeWonfLm0TarWDyouNnO967XYnB_195UmVZ2EFZMa5kl3Q==
```

| 字段 | 说明 |
|:--|:--|
|key|AppID|
|code|授权码|
|timestamp|证书生成时间|
|expiry|有效期（天）,0表示永久有效|
|storage|数据存储，可存放自定义的数据|
|cer|证书|
|sig|授权文件的签名|

# 授权文件制作流程

> pubkey >> {AES->BASE64} >> cer

> [key, code, timestamp, expiry, storage, cer] >> {merge} >> payload

> payload >> {AES->MD5} >> identity

> identity >> {RSA->BASE64} >> sig


- 将pubkey进行AES加密和BASE64编码，得到cer。
- 将API传入的 key、code、timestamp、expiry、storage和生成的cer合并为payload。
- 将payload进行AES加密和MD5取值，得到identity。
- 对identity进行RSA签名和BASE64编码，得到sig。
- 将payload和sig写入授权文件。


# 授权文件验证流程

> [key, code, timestamp, expiry, storage, cer] >> {merge} >> payload

> payload >> {AES->MD5} >> identity

> cer >> {BASE64->AES} >> pubkey

> [pubkey, identity, sig] >> {RSA}

- 读取授权文件中的payload部分(key、code、timestamp、expiry、storage、cer)。
- 使用payload进行AES加密和MD5取值后，得到identity。
- 读取授权文件中的cer，BASE64解码后AES解密，得到pubkey。
- 使用pubkey验证identity和sig是否匹配。

# 授权文件验证错误值

| 错误码 | 说明 |
|:--|:--|
|0|无错误|
|1|无效的授权文件|
|2|缺少字段|
|3|证书解码错误|
|4|证书解密错误|
|5|签名解码错误|
|6|签名验证错误|
|7|时间戳解析错误|
|8|有效期解析错误|
|14|授权文件过期|

