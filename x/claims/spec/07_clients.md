<!--
order: 7
-->

# Clients

A user can query the `x/claims` module using the CLI, gRPC or REST.

## CLI

Find below a list of `fortressd` commands added with the `x/claims` module. You can obtain the full list by using the `fortressd -h` command.

### Queries

The `query` commands allow users to query `claims` state.

**`total-unclaimed`**

Allows users to query total amount of unclaimed tokens from the airdrop.

```bash
fortressd query claims total-unclaimed [flags]
```

**`claims-records`**

Allows users to query all the claims records available.

```bash
fortressd query claims claims-records [flags]
```

**`claims-record`**

Allows users to query a claims record for a given user.

```bash
fortressd query claims claims-record [address] [flags]
```

**`params`**

Allows users to query claims params.

```bash
fortressd query claims params [flags]
```

## gRPC

### Queries

| Verb   | Method                                     | Description                                      |
|--------|--------------------------------------------|--------------------------------------------------|
| `gRPC` | `fortress.claims.v1.Query/TotalUnclaimed`     | Gets the total unclaimed tokens from the airdrop |
| `gRPC` | `fortress.claims.v1.Query/ClaimsRecords`      | Gets all registered claims records               |
| `gRPC` | `fortress.claims.v1.Query/ClaimsRecord`       | Get the claims record for a given user            |
| `gRPC` | `fortress.claims.v1.Query/Params`             | Gets claims params                               |
| `GET`  | `/fortress/claims/v1/total_unclaimed`         | Gets the total unclaimed tokens from the airdrop |
| `GET`  | `/fortress/claims/v1/claims_records`          | Gets all registered claims records               |
| `GET`  | `/fortress/claims/v1/claims_record/{address}` | Gets a claims record for a given user            |
| `GET`  | `/fortress/claims/v1/params`                  | Gets claims params                               |
