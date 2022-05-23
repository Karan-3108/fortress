<!--
order: 8
-->

# Clients

## CLI

Find below a list of  `fortressd` commands added with the  `x/erc20` module. You can obtain the full list by using the `fortressd -h` command. A CLI command can look like this:

```bash
fortressd query erc20 params
```

### Queries

| Command         | Subcommand    | Description                    |
| --------------- | ------------- | ------------------------------ |
| `query` `erc20` | `params`      | Get erc20 params               |
| `query` `erc20` | `token-pair`  | Get registered token pair      |
| `query` `erc20` | `token-pairs` | Get all registered token pairs |

### Transactions

| Command      | Subcommand      | Description                    |
| ------------ | --------------- | ------------------------------ |
| `tx` `erc20` | `convert-coin`  | Convert a Cosmos Coin to ERC20 |
| `tx` `erc20` | `convert-erc20` | Convert a ERC20 to Cosmos Coin |

### Proposals

The `tx gov submit-proposal` commands allow users to query create a proposal using the governance module CLI:

**`register-coin`**

Allows users to submit a `RegisterCoinProposal`.

```bash
fortressd tx gov submit-proposal register-coin [metadata] [flags]
```

**`register-erc20`**

Allows users to submit a `RegisterERC20Proposal`.

```bash
fortressd tx gov submit-proposal register-erc20 [erc20-address] [flags]
```

**`toggle-token-conversion`**

Allows users to submit a `ToggleTokenConversionProposal`.

```bash
fortressd tx gov submit-proposal toggle-token-conversion [token] [flags]
```

**`param-change`**

Allows users to submit a `ParameterChangeProposal``.

```bash
fortressd tx gov submit-proposal param-change [proposal-file] [flags]
```

## gRPC

### Queries

| Verb   | Method                            | Description                    |
| ------ | --------------------------------- | ------------------------------ |
| `gRPC` | `fortress.erc20.v1.Query/Params`     | Get erc20 params               |
| `gRPC` | `fortress.erc20.v1.Query/TokenPair`  | Get registered token pair      |
| `gRPC` | `fortress.erc20.v1.Query/TokenPairs` | Get all registered token pairs |
| `GET`  | `/fortress/erc20/v1/params`          | Get erc20 params               |
| `GET`  | `/fortress/erc20/v1/token_pair`      | Get registered token pair      |
| `GET`  | `/fortress/erc20/v1/token_pairs`     | Get all registered token pairs |

### Transactions

| Verb   | Method                             | Description                    |
| ------ | ---------------------------------- | ------------------------------ |
| `gRPC` | `fortress.erc20.v1.Msg/ConvertCoin`   | Convert a Cosmos Coin to ERC20 |
| `gRPC` | `fortress.erc20.v1.Msg/ConvertERC20`  | Convert a ERC20 to Cosmos Coin |
| `GET`  | `/fortress/erc20/v1/tx/convert_coin`  | Convert a Cosmos Coin to ERC20 |
| `GET`  | `/fortress/erc20/v1/tx/convert_erc20` | Convert a ERC20 to Cosmos Coin |
