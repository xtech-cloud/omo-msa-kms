# 简介

密钥管理服务

支持http和grpc两种调用方式

# 编译

进入omo-msa-kms目录，执行compile.sh脚本


# 部署

设置环境变量

```
export GIN_MODE=release
export KMS_HTTP_ADDR=:80
export KMS_GRPC_ADDR=:10080
```

- SQLite

设置以下环境变量

```
export KMS_DATABASE_DRIVER=sqlite
export KMS_SQLITE_FILEPATH=<db文件所在路径>
```

- Mysql

设置以下环境变量

```
export KMS_DATABASE_DRIVER=mysql
export KMS_MYSQL_ADDR=127.0.0.1:3306
export KMS_MYSQL_USER=<mysyql用户>
export KMS_MYSQL_PASSWORD=<mysql密码>
export KMS_MYSQL_DATABASE=kms
```

设置完环境变量后启动kms


# HTTP API

## app

### `/kms/app/create`

**简要描述:**

- 创建应用
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname|是  |string |应用名|


**返回示例**

```json
{
    "code":0, 
    "message":"",
    "data":{
        "appkey":"ee488cd5bf1363902a1bcb0359464e48",
        "appsecret":"42f7f2011c1a7a3c7d9ab26332ef99bb",
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| appkey | string | 应用识别码|
| appsecret | string | 应用密钥 |

### `/kms/app/query`

**简要描述:**

- 查询应用
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname|是  |string |应用名|


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
        "AppName":"test2",
        "AppKey":"bf34b04d9698ca9c95f5def56d960827",
        "AppSecret":"ebe8a84a8708b6cef93d2449bd650910",
        "PrivateKey":"-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAlL7SgYxOYxREy40esuu0jgd2CoDyMRvFRBWpR5he2VQcy8o8\nMSKAfDzH26FVZc9z3VafhD5GVlIO13gBK/336vnN/VBHPQ36F8HGWM7YekQ9+N5Q\ncT4mWqUQsWUatbQFi2C86OpmLGCL3jq2i7S6mxNLJbe7E+dg+cy0SagIlXlbnNXr\ncHnuJoVXBiwhmig1PzHgTRvSjXKXAIjQGWkiCdiyzYx5bif7gNW/ib7PJruXfFLI\n4lF2veaylYcbC/+BY42YNoAWMsGBWnCBLCCLVyKCt698bVftOELUtZOy9+TcZU71\n+pOwPf+eeqhK4pAXDGG1pLz1hBhmVbVlg7Rk0wIDAQABAoIBABYD+Fw8TA3WHiiS\nhoys3lh3Oj1rwG0MUzI0ko2KO9+m12xCTo5nMOUyidI0GtOq1NdZztpf7UExfAjg\nNiwwttUMjDSGAUVEIFQL1jOmyduu5g1DulxIepzH+aSH9mAWeQucEdnXd6/xykHm\nJsaexU/WlzTJ8OKNSIkwhy6vtDWLOwQm+yldYu0TNZJ0jIs1hVjjLsbf/4h6Hk3y\nYIUa1T5NJpOmXULSfgxADNLFFprLNBqFeHarq8u+g38TOLzYC1eF8bsLPWe1Wzoa\nPWCtKcNOvMPPSEjv85KOjvF7Q8OnmEzKhwtgSu8AlxMgVG9lAc8XrWA9S2xSs20l\nJMM3IyECgYEAwoA/moXrevn/SsrQURjdvSjpI1e1C+6Dhu9ivwfMQqdpX5WD23TY\n4f3EWXN8u0Nz5aVHdd7O3bLllZjcUnJay4jFKqNZyjWaGxNmJZDHNth8unlr4fki\nEc3Z9H33FlCQG2/cuq+P7ujV7g1wTAM1yy/IGhUrjyCnme0CE4ZkAK8CgYEAw8br\na9OiSfAVO3thwK0m++BuVC6LqSe0gVazCgJtn0aAgmLeeMhwHoG+7DAtWau0PyYS\nh5UVYF5yaBvX3Zs0ze9wqzMp38954B25nOA3lv1QlnQ1I15aDZlvsreALnhe464d\nJiEfP0rvd4iTiKUKqsOKmtL0bkNl03xtLGXu/x0CgYBcHV0CE7aocUnE5DSwk7RA\nZ+WyRVGLKxTDjRAZJNpKHvs6t5bREo+8x/B75MQH9DQpaJNlcXZLbPRqWxDNQzdY\n+ZdXUDGwIJ6xgAh6dgzDHthDgEnlpZXLFNDKh/XDbbgyJlJFX+ws27yll1u9xC9v\n4VtFbw1IJdD6h1LaaGVoJQKBgQCAbEYV7zev5Kso36CZ8Xt3EhuNYRMAHSmNBkBf\nuoQKTQcTgKOK+4CAon+JE3lMLxQHsIPLKIJjOtE1db4+ggc7Z2uzAdbgF4tM9nLB\nc1tD0ltAtm39C3FrJlFdHH4a/Z7RH2/DiUqkDBXVhWOx6QF8TtTnBqaMhe2Pszky\nPJNwCQKBgDTUcA7sBouf3f23CVY0q073NnT/eDgPrhIe1oc9/SWLnteVRYqc6QCL\nOsXFUzdTLbVskyhuDSaebgIjdVlSY+voxPAMpVyQAXZGXoemKH+oIQpMf8/QBBdi\nlK/ypitZJ9jRXgrXWhie4etiXp+f6UAbTFNQU3W3ZANUtjz9WjT0\n-----END RSA PRIVATE KEY-----\n",
        "PublicKey":"-----BEGIN RSA PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlL7SgYxOYxREy40esuu0\njgd2CoDyMRvFRBWpR5he2VQcy8o8MSKAfDzH26FVZc9z3VafhD5GVlIO13gBK/33\n6vnN/VBHPQ36F8HGWM7YekQ9+N5QcT4mWqUQsWUatbQFi2C86OpmLGCL3jq2i7S6\nmxNLJbe7E+dg+cy0SagIlXlbnNXrcHnuJoVXBiwhmig1PzHgTRvSjXKXAIjQGWki\nCdiyzYx5bif7gNW/ib7PJruXfFLI4lF2veaylYcbC/+BY42YNoAWMsGBWnCBLCCL\nVyKCt698bVftOELUtZOy9+TcZU71+pOwPf+eeqhK4pAXDGG1pLz1hBhmVbVlg7Rk\n0wIDAQAB\n-----END RSA PUBLIC KEY-----\n",
        "Profile":"",
        "CreatedAt": "2019-01-01 13:13:13.22333 +0000 UTC"
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| AppName| string | 应用名|
| AppKey| string | 应用识别码|
| AppSecret| string | 应用密钥 |
| PublicKey| string | 公钥|
| PrivateKey| string | 私钥|
| Proafile| string | 简介|
| CreatedAt| string | 应用创建时间|

### `/kms/app/list`

**简要描述:**

- 列出应用
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
| appname | 是 | string |  应用名 |


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
        "Applications":["test", "test2"]
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| Applications| string[] | 应用名列表|

### `/kms/app/modify/profile`

**简要描述:**

- 更改简介
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname | 是| string | 应用名|
|profile| | string | 简介|


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
    }
}
```

### `/kms/app/modify/security`

**简要描述:**

- 更改简介
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appname | 是| string | 应用名|
|appkey| | string |应用识别码 |
|appsecret| | string | 应用密钥 |
|publickey| | string | 公钥|
|privatekey| | string | 私钥|


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
    }
}
```



## key 

### `/kms/key/generate`

**简要描述:**

- 生成激活码
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appkey|是  |string |应用识别码|
|appsecret|是  |string |应用密钥|
|count|  |int|生成的激活码的数量|
|capacity|  |int|激活码可激活的设备数量|
|expiry||int|有效期，以天为单位，默认为0（永久）|
|storage|  |string |自定义数据|
|profile|  |string |激活码简介|

**返回示例**

```json
{
    "code": 0,
    "message": "",
    "data":{
        "keys":["ccc292740d9772a7ed78a6c76f880de8"]
    }
} 
```

**返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|:-----|
| keys | string[] |生成的激活码列表|


### `/kms/key/query`

**简要描述:**

- 查询激活码
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|number|是  |string |激活码|

**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
        "AppName":"test",
        "Capacity":1,
        "CreatedAt":"2019-05-09 09:32:58.1771754 +0000 UTC",
        "Expiry":0,
        "Number":"ccc292740d9772a7ed78a6c76f880de8",
        "Status":0,
        "Storage":"",
        "Active":["PAJ2392323"]
    }
}
```

**返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|:-----|
| AppName | string |应用名|
| Capacity| string |可激活的最大设备数目|
| CreatedAt| string |激活码生成的时间|
| Expiry| string |有效期（天），0代表永久有效|
| Number | string |激活码|
| Status| string |自定义的激活码状态，非0代表激活码无法激活|
| Storage| string |自定义数据字段|
| Active| string[] |已经激活的设备|

### `/kms/key/list`

**简要描述:**

- 列出激活码
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|appkey|是|string|应用识别码|
|appsecret|是|string|应用密钥|

**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
        "Keys":["ccc292740d9772a7ed78a6c76f880de8","b51fae6a1fedc4d55d4034e608858c85"]
    }
}
```

**返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|:-----|
| Keys | string[] |激活码列表|

### `/kms/key/modify/profile`

**简要描述:**

- 更改简介
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|number| 是| string | 激活码|
|profile| | string | 简介|


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
    }
}
```

### `/kms/key/modify/status`

**简要描述:**

- 更改激活码状态
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|number| 是| string | 激活码|
|status| | int| 状态，非0代表无法激活|


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
    }
}
```

### `/kms/key/activate`

**简要描述:**

- 激活
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|number| 是| string | 激活码|
|code| 是| string| 设备码|


**返回示例**

```json
{
    "code":0,
    "message":"",
    "data":{
        "uid":"923d2d4092a4e5a22ea39767c1407cec",
        "license":"key:\n13bb262e2c7a425008e641748eb4741a\ncode:\nPAJ2392323\ntimestamp:\n1557394406\nexpiry:\n0\nstorage:\n{\"company\":\"wawawa\"}\ncer:\namdLTExQCLqhRQS5wkQZIUvM4i4b_W35pgYStVDuzTjjDTutkARTFC9Hu1ApToaQWmrZ2bdp0Se6ogevOLFF7Ukd_1z8GSX4pu9VhLq5JQajqHftZFgnt6vnfXkZg0vGNkoJ23PqPGw__qH48vlCyaEtXK5ipIK5iPcSrYTjnbDn4krkhXPlMZ2yGDORc4LuPe8ZRZF46dilHZ_NMRUYJzddHnl69ttmk-vdo1QGDu9BxqUQeuMs7sLyxMDXxwmvOr2-n7a6pkx2Gv0OdmYkyt9GvAnW_dIR7G5_l1NXXYWRTmla4xL6RQTvWi-vr9aY_e0zRrlfrtvGqUjiJmdCnjCxG_gHBwYZT5wFia3jEgM2_rQg6gWjwEWrsaAlZ4NOBH52favXNbH1t331_QM_8DJB8XfPQF0sxpNIv7dHrZGF7i-lLsRuEVH-NvzgM7jaYjME-3mIVbNuwSWVgATPYPfvLMU6NGuooTy6XecmCAeFOdCrCdPGCWbNddwkJY8tBeuIgytP2j2wntSF7lnNiQ11efgcrf0mWsF7BtU6W8laQ2nP62-EGzh59NBGGXZLDUMFb4YDdu37Kf77VVl9q9bdc-Kz5vZxrZrqj1Wt-uc=\nsig:\nVIlE4Y0hu0NW8ekgioOpCXu0qAlXN8RHkILckLW1jS3LNUEH8jIrq5w7E4RiQLniaZX1wTHa3MjQCjuF4aiFp2v0pwTjMr8-iQWrHrL5vIhqbBmjSWynNdm93sQawSh_by7xFGpm0QW2sP5FBEdhyJjwcjQAAcTrVgJS9-07xb-i3rj1l8CALv2JULjXhMjg5gfTGf9vuR397IvlWjdLgm15cdtYbhf5x4jX431pESI7GvXlXW0MYxcWLmIwTWa9CvR8ygV8dcnH-__HC3H4pbu9PeI-P9TOlGAkPmN-odUvvnk2tAfi6t4cp70RpYLjV2oINEL4PR0tg7oJ1OXWUw=="
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:----|:---|:--- |
|uid|string | 授权文件ID|
|license|string| 授权文件内容|


**备注** 

|错误码|说明|
|:----|:---|
|1|无效激活码|
|2|激活码状态不允许|
|3|激活设备已达上限|



## license

### `/kms/license/verify`

**简要描述:**

- 验证授权证书
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|uid|是  |string |授权文件ID|
|code|是  |string |设备码|

**返回示例**

```json
{
    "code": 0,
    "message": "",
    "data": {
        "status": 0
    }
} 
```

**返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|:-----|
| status| int | 对应的激活码的状态，非0表示无法激活|

**备注** 

| 错误码| 说明 |
| :--- | :--- |
| 1 | 设备码不匹配|

### `/kms/license/fetch`

**简要描述:**

- 获取授权证书
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|code|是  |string |设备码|

**返回示例**

```json
{
    "code": 0,
    "message": "",
    "data": {
        "license":"key:\n13bb262e2c7a425008e641748eb4741a\ncode:\nPAJ2392323\ntimestamp:\n1557394406\nexpiry:\n0\nstorage:\n{\"company\":\"wawawa\"}\ncer:\namdLTExQCLqhRQS5wkQZIUvM4i4b_W35pgYStVDuzTjjDTutkARTFC9Hu1ApToaQWmrZ2bdp0Se6ogevOLFF7Ukd_1z8GSX4pu9VhLq5JQajqHftZFgnt6vnfXkZg0vGNkoJ23PqPGw__qH48vlCyaEtXK5ipIK5iPcSrYTjnbDn4krkhXPlMZ2yGDORc4LuPe8ZRZF46dilHZ_NMRUYJzddHnl69ttmk-vdo1QGDu9BxqUQeuMs7sLyxMDXxwmvOr2-n7a6pkx2Gv0OdmYkyt9GvAnW_dIR7G5_l1NXXYWRTmla4xL6RQTvWi-vr9aY_e0zRrlfrtvGqUjiJmdCnjCxG_gHBwYZT5wFia3jEgM2_rQg6gWjwEWrsaAlZ4NOBH52favXNbH1t331_QM_8DJB8XfPQF0sxpNIv7dHrZGF7i-lLsRuEVH-NvzgM7jaYjME-3mIVbNuwSWVgATPYPfvLMU6NGuooTy6XecmCAeFOdCrCdPGCWbNddwkJY8tBeuIgytP2j2wntSF7lnNiQ11efgcrf0mWsF7BtU6W8laQ2nP62-EGzh59NBGGXZLDUMFb4YDdu37Kf77VVl9q9bdc-Kz5vZxrZrqj1Wt-uc=\nsig:\nVIlE4Y0hu0NW8ekgioOpCXu0qAlXN8RHkILckLW1jS3LNUEH8jIrq5w7E4RiQLniaZX1wTHa3MjQCjuF4aiFp2v0pwTjMr8-iQWrHrL5vIhqbBmjSWynNdm93sQawSh_by7xFGpm0QW2sP5FBEdhyJjwcjQAAcTrVgJS9-07xb-i3rj1l8CALv2JULjXhMjg5gfTGf9vuR397IvlWjdLgm15cdtYbhf5x4jX431pESI7GvXlXW0MYxcWLmIwTWa9CvR8ygV8dcnH-__HC3H4pbu9PeI-P9TOlGAkPmN-odUvvnk2tAfi6t4cp70RpYLjV2oINEL4PR0tg7oJ1OXWUw=="
    }
} 
```

**返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|:-----|
|license|string| 授权文件内容|

**备注** 

| 错误码| 说明 |
| :--- | :--- |
| 1 | 设备码不匹配|
| 2 | 证书状态错误 |

