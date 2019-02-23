
---

[TOC]

---

# API列表

## app

### `/kms/app/create`

** 简要描述: **

- 创建应用
  
** 请求方式： **

- POST 

** 参数： ** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname|是  |string |应用名|


** 返回示例 **

```json
{
    "code":0, 
    "message":"",
    "data":{
        "appkey":"ee488cd5bf1363902a1bcb0359464e48",
        "appsecret":"42f7f2011c1a7a3c7d9ab26332ef99bb",
        "privatekey":"-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAuUuenUO6C+qto/EIEIsriyu5WZh02+cd8blHlx3hPbBH2bq1\nf9yarm94mjGXBpoMDHyPqyGhxW8iGQhadQg2pUE0EX8GASRYKV529sJPT/GxFZF+\nqvJ3Cx9/PYPq25yOMLGdjldJCzHPDXL9NL+NQTAQJyl96siafQrh5MGY6FZsYQ4U\nY254NmNvQBIcHsJzMgUQRgHiD8P7zMoiUwNs/nUKXsPcL/Gp1hSQNCS3xkFdHW1Y\nxPAuXHca4aLtAgSgHXbCNV7d9WVD+Bxa4HCoj4kBgyvSGxm3UuaSuG1U4HL0JA7o\nvOo5WxQtk+Yh0RZMpGEMnZ72JfuzufjJWIAylQIDAQABAoIBAQCpdOWvD0QzJ8D7\nIDBu6Me+tBMDRDEC9t9ktDwrkHDwOKnNDjp0n6x6gIk8AsQKjpEcTkIZkF1gsRzu\n5bvMMeG/ydyzb4ZvPuy9kJ2dV+CuzOtVPUUpUyrZakm0WVNM2mUTVSwxA9RSP5AK\nzbGRLukgx7LODd0Q/bqDRpOF4CNVvjpRWMGxSE8TR84ll0mtMM16jEDatoXZ+ru7\n1b364W+7nqBHc+iFAS75dTPi/Aq8berS8ejcWSY1sLwclBwNgKnrPeMtwRTLeAQb\nLOLKAWLmL7jDE+N9M3QllqD2V/162qFtN4e/iV6QXRMzFtZcYAGaxBXo+mU97a9t\n43QCY6mBAoGBAPa7X+GGb8DDA7i8bZaJgk1aF2zT3a3sSu+z0vHP/NkC/yf66Vm6\nDY3xCZ28Hh0Y7x3b3oqThM/kfIAg5UcCoxZXqug9z3xnLnT6sl1r9bo22k/BScKN\nN0Fc9pyLxsfsF/o10DStfLOoRmpRn7CTST0T+oTACS/TgY5UdZf+AwIhAoGBAMBB\ndWCjeB0hbftEfnaj1cKFhfoJtxhnJQyZXxgCTvg3a1w8ppv2+896f7Ss3KySXdr6\nwPSd0yaJaQomsy5aDImh0NR4riC17ezNfpG98hm9hU5Fb6W/5+5KmYEGl/O7JkNX\njuJddaeFCuvFewmCAxNe6iS7mwk2qavWWzWrDAn1AoGBAIfYJHM0JIVaGct1m1S9\nwws5phoaEDx5E80kEelnXUxSVQ+7Gw123Y2f/25kK6RTnRuwebeMfxxDLbUR54qQ\nTiC4BfY2Se2xlad8fAXpsfJjzxEIV8cRHzISsAkDAGAEgjMVu9u1BSqZZKOW6zg5\n76H8RPsgDC7T4cnugIw9o4RhAoGACV9nlIWDpo/pp7VIIQqNskUGKhFtLrBgmwOj\nIi2CohE5l66RRMs+rXfAYBOJUpR0JOfSnlpPX/KU/1yxoZHcSJ1t693SR1/3MHM4\nN2y2L9EQfade8QqxCOn0H9ktcUFvdRsAqUJ7KOaQiLwA6o6/DaOz8ISA42sZzHnP\nyocDOD0CgYBuMJ8+GSq5GFFbSEmbhJgpCoXamGtXGR43EneqphfewcCvstueFYQ4\n9ojC4/Qd4x+q40PoZNQ6AheFfkqfknChVoEAwSxAQOmz1RJe/IVruKygxahUmI6h\n870wGuWFeDyNqeL6nbACpXWkYA2JJ/CId7Tmf3HN5nVT+sjKwqCRrQ==\n-----END RSA PRIVATE KEY-----\n",
        "publickey":"-----BEGIN RSA PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuUuenUO6C+qto/EIEIsr\niyu5WZh02+cd8blHlx3hPbBH2bq1f9yarm94mjGXBpoMDHyPqyGhxW8iGQhadQg2\npUE0EX8GASRYKV529sJPT/GxFZF+qvJ3Cx9/PYPq25yOMLGdjldJCzHPDXL9NL+N\nQTAQJyl96siafQrh5MGY6FZsYQ4UY254NmNvQBIcHsJzMgUQRgHiD8P7zMoiUwNs\n/nUKXsPcL/Gp1hSQNCS3xkFdHW1YxPAuXHca4aLtAgSgHXbCNV7d9WVD+Bxa4HCo\nj4kBgyvSGxm3UuaSuG1U4HL0JA7ovOo5WxQtk+Yh0RZMpGEMnZ72JfuzufjJWIAy\nlQIDAQAB\n-----END RSA PUBLIC KEY-----\n"
    }
}
```

** 返回参数说明 ** 

|参数名|类型|说明|
|:-----  |:-----|:-----|
| appkey | string | 应用识别码|
| appsecret | string | 应用密钥 |
| privatekey | string | 私钥 |
| publickey | string | 公钥 |

** 备注 ** 

## license

### `/kms/license/make`

** 简要描述: **

- 生成授权证书
  
** 请求方式： **

- POST 

** 参数： ** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appkey|是  |string |应用识别码|
|appsecret|是  |string |应用密钥|
|devicecode|是  |string |设备码|
|publickey|是  |string |公钥|
|privatekey|是  |string |私钥|
|expiry|否|int|有效期，以天为单位，默认为0（永久）|
|storage|否  |string |自定义数据|

** 返回示例 **

```json
{
    "code": 0,
    "message": "",
    "data":{
        "license": "key:\n1c89cd9a1f1f78f9a9085b6809441dca\ncode:\n00000000\ntimestamp:\n1550909349\nexpiry:\n0\nstorage:\n{}\ncer:\nQT5gCZAZPMWgk9Gati6azD-DqbI9OAp7kQNj0hYL0TI3DanDnSf0tXjcXxRJARgDUH8D5HlFM2VleTbGCioegWUiLqp7pBzawX3qzrU8uP5ci_Vn6wG_CXvApY5HoIT2AjvLGxMeQra-uxIZz1zQsOK1xOUJzbIJsKVWnlGX_yDKte3-YtYPfYBKVyp3JlPwUYLk0t4Rzb9fztcKZKJEGcNAqusLFYDOWad8czSjuZ9s6kvyFWzgajY5gB_QoOKPx64-B1RfG5JwPzO1xyBLu0KC1lzULjWQRdngGd4unwWP-auDDt2uPuvMjCnPnl4cHD_cQa6PQMaXnNCKaAEdf7P1wM0EDXOWvq3pPDKwREgyzDGuk1DcxGH58Vq4PzaXmnLeyXqqcMk7aXnmM8hi-SOOwKYNWSa6lH19DzQfKa-kg8MtGYimlp76sEXW-AsAOtfM72Tr0ll2xWnLE13oe10bqdlcOrYEb0SZgANwqv01KEgc3S6AnUglR5_E5KY0GcyGLcDNGQy5LMvu_eYT9JAw1a8vvJTNoLQFXDCza7_zvQIsXa8MZJESV7oouIikPCMDLDi7IoEpj_lxcNTJclPMjLynh5uFkTju4KnlNJA=\nsig:\neU8kwP6QOCsWfhSPP9NrNp5f8V8ZQQD78gLkUwZpycRJXUGpt_DbpRxNHwiKSQVcdIdOah9Exf-sXAP_z2WZA176JzNsQ3JP7oCLUZ-NN7M5b1eeuRYNSjGk_SIz_g0hrqlROjdfHoyEusbCnsE6hNPJ0WX6DzBn4mDbIStQ_nLTMTd8REFW6lWBQt-o4gkitcL5dmdPT2oMcQiBKMd8g8Unyyc4142UlO-_XufJxGmkU5ssf6Jh4dtR2q64mgQSmD_qYuT9Yu7qrVzbeUZOlhFjKjAdr6_r0meYVF4jj_bgA7bk7BZqyZZuM_6IE4ZjCcQaje4EH8sX-Vlwqpmskw=="
    }
} 
```

** 返回参数说明 ** 

|参数名|类型|说明|
|:-----  |:-----|:-----|
| license | string | 授权证书 |

** 备注 ** 


### `/kms/license/verify`

** 简要描述: **

- 验证授权证书
  
** 请求方式： **

- POST 

** 参数： ** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appkey|是  |string |应用识别码|
|appsecret|是  |string |应用密钥|
|devicecode|是  |string |设备码|
|license|是  |string |授权证书|

** 返回示例 **

```json
{
    "code": 0,
    "message": "",
    "data": {
        "result": 0
    }
} 
```

** 返回参数说明 ** 

|参数名|类型|说明|
|:-----  |:-----|:-----|
| result | int | 验证结果|

** 备注 ** 

| 验证结果 | 说明 |
| :--- | :--- |
| 0 | 无错误 |
| 1 | 无效证书 |
| 2 | 缺少字段 |
| 3 | 证书解码错误 |
| 4 | 证书解密错误 |
| 5 | 签名解码错误 |
| 6 | 签名解密错误 |
| 7 | 时间戳解析错误 |
| 8 | 有效期解析错误 |
| 14 | 授权文件过期 |
