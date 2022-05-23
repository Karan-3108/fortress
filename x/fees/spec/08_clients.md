<!--
order: 8
-->

# Clients

## CLI

Find below a list of  `fortressd` commands added with the  `x/fees` module. You can obtain the full list by using the `fortressd -h` command. A CLI command can look like this:

```bash
fortressd query fees params
```

### Queries

| Command        | Subcommand    | Description                    |
| :------------- | :------------ | :----------------------------- |
| `query` `fees` | `params`      | Get fees params                |
| `query` `fees` | `fee-info`    | Get registered fee info        |
| `query` `fees` | `fee-infos`   | Get all registered fee infos   |

### Transactions

| Command     | Subcommand      | Description                                |
| :---------- | :-------------- | :----------------------------------------- |
| `tx` `fees` | `register-fee`  | Register a contract for receiving fees     |
| `tx` `fees` | `update-fee`    | Update the withdraw address for a contract |
| `tx` `fees` | `cancel-fee`    | Remove the fee info for a contract         |

## gRPC

### Queries

| Verb   | Method                                       | Description                                 |
| :----- | :------------------------------------------- | :------------------------------------------ |
| `gRPC` | `fortress.fees.v1.Query/Params`                 | Get fees params                             |
| `gRPC` | `fortress.fees.v1.Query/DevFeeInfo`             | Get registered fee info                     |
| `gRPC` | `fortress.fees.v1.Query/DevFeeInfos`            | Get all registered fee infos                |
| `gRPC` | `fortress.fees.v1.Query/DevFeeInfosPerDeployer` | Get all registered fee infos for a deployer |
| `GET`  | `/fortress/fees/v1/params`                      | Get fees params                             |
| `GET`  | `/fortress/fees/v1/fees/{contract_address}`     | Get registered fee info                     |
| `GET`  | `/fortress/fees/v1/fees`                        | Get all registered fee infos                |
| `GET`  | `/fortress/fees/v1/fees/{deployer_address}`     | Get all registered fee infos for a deployer |

### Transactions

| Verb   | Method                                    | Description                                |
| :----- | :---------------------------------------- | :----------------------------------------- |
| `gRPC` | `fortress.fees.v1.Msg/RegisterDevFeeInfo`    | Register a contract for receiving fees     |
| `gRPC` | `fortress.fees.v1.Msg/UpdateDevFeeInfo`      | Update the withdraw address for a contract |
| `gRPC` | `fortress.fees.v1.Msg/CancelDevFeeInfo`      | Remove the fee info for a contract         |
| `POST` | `/fortress/fees/v1/tx/register_dev_fee_info` | Register a contract for receiving fees     |
| `POST` | `/fortress/fees/v1/tx/update_dev_fee_info`   | Update the withdraw address for a contract |
| `POST` | `/fortress/fees/v1/tx/cancel_dev_fee_info`   | Remove the fee info for a contract         |
