<!--
parent:
  order: false
-->

<div align="center">
  <h1> Savour RPC Repo </h1>
</div>

<div align="center">
  <a href="https://github.com/SavourDao/snow/releases/tag/v1.0.0-alpha">
    <img alt="Version" src="https://img.shields.io/github/tag/snow/savour-core.svg" />
  </a>
  <a href="https://github.com/SavourDao/savour-core/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/SavourDao/savour-core.svg" />
  </a>
</div>

#### 1. Introdcution

This document describes the protocols and data formats for communication between the various Savour systems.

#### 2. protocol

- GRPC
- Proto 3.0
- Do not enable TLS

#### 3. repo link

https://github.com/SavourDao/savour-proto.git

#### 4. ReturnCode

```
enum ReturnCode{
    SUCCESS = 0;
    ERROR = 1;
}
```

Two kinds of status codes are currently defined

- 0: SUCCESS;
- 1: ERROR;


Example: Query supported currencies and broadcast transactions:

```
service WalletService {
  rpc getSupportCoins(SupportCoinsRequest) returns (SupportCoinsResponse) {}
  rpc SendTx(SendTxRequest) returns (SendTxResponse) {}
}
```

#### 5.The port on which the business is listening (please update it yourself)
[finish]
