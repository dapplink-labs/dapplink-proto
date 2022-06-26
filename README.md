### Savour RPC 文档
版本: v1.0
时间: 2022/6/26

#### 1. 概要

本文档描述了 Savour 各个系统间通讯的协议和数据格式。

#### 2. 协议

- GRPC
- Proto 3.0
- 不开启 TLS

#### 3. 地址

https://github.com/SavourDao/savour-protocal.git

#### 4. 错误约定
数据结构:
```
message Error {
  int32  code      = 1; // 0: 成功; 404: NotFound; : 失败;
  string brief     = 2;
  string detail    = 3;
  bool   can_retry = 4; // true: 可以重试，false: 不能, 在写入失败的时候，该标识位提示客户端来决定是否可以重试。
}
```
目前定义了三种状态码，
- 0: 成功;
- 404: NotFound;
- 1: 其他失败;
-- 关于字段 can_trey:
-- 表示该 API 是否可以重试。
-- 在写入失败的时候，该标识位提示客户端来决定是否可以重试。
-- 在某些特殊的 API，重试可能会导致问题的，返回 false。

例子: 查询支持的币种和广播交易:
```
service WalletService {
  rpc getSupportCoins(SupportCoinsRequest) returns (SupportCoinsResponse) {}
  rpc SendTx(SendTxRequest) returns (SendTxResponse) {}
}
```
如果该地址不存在 响应值为(采用 JSON 格式描述)
```
{
  "code": 404,
  "brief": "Not found user!",
  "detail": "",
  "can_retry": false
}
```
其他公共的数据结构定义
返回空的时候采用自定义的 Empty。
更多请参考各个实现

#### 5. 业务监听的端口(请自行更新)
[完]
