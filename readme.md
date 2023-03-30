# StreamPay Application
StreamPay is a dApp that's built on the Stream Pay Testnet (which utilizes the `paymentstream` module. This sovereign chain is powereed by the token `uspay` or `SPAY`.

While users need to pay the fee to start a payment stream in SPAY tokens, any token (native or IBC tokens) can be streamed between accounts.

Try the app on - https://sp.OmniFlix.network

# paymentstream module
**paymentstream** is a module built using Cosmos SDK, Tendermint and [Starport](https://github.com/tendermint/starport) to stream payments from on address to the other, using `delayed` and `continuous` payments types inspired by the vesting model in the `auth` module of the Cosmos SDK.

## Installation

### Requirements

| Requirement | Notes                                |
|-------------|--------------------------------------|
| Go version  | [Go1.17](https://go.dev/doc/install) |
| Cosmos SDK  | v0.44.4                              |

### Get source code & Install

```bash=
git clone https://github.com/OmniFlix/payment-stream.git
cd payment-stream
go mod tidy
make install
```
check installation
```bash=
payment-streamd version
```


## Launch chain manually

```bash=
# Delete previous data
rm -rf ~/.payment-stream/config/*

#Init node	
payment-streamd unsafe-reset-all
payment-streamd init "sp-node"  --chain-id "sp-test-1"

# Add keys
payment-streamd keys add validator --keyring-backend test
payment-streamd keys add user1 --keyring-backend test
payment-streamd keys add user2 --keyring-backend test

# Add genesis accounts
payment-streamd add-genesis-account $(payment-streamd keys show validator -a --keyring-backend test) 1000000000stake
payment-streamd add-genesis-account $(payment-streamd keys show user1 -a --keyring-backend test) 1000000000000stake

# Create gentx
payment-streamd gentx validator 10000000stake --moniker "validator-1" --chain-id "sp-test-1" --keyring-backend test

# Collect Gentxs
payment-streamd collect-gentxs
payment-streamd validate-genesis

# Start Chain with default config
payment-streamd start
```
### CLI Commands

### Transactions
**stream-send** - to start a payment stream
```bash=
$ payment-streamd tx paymentstream stream-send -h
```
```bash=
creates a payment stream

Usage:
  payment-streamd tx paymentstream stream-send [flags]

Examples:
$ payment-streamd tx paymnetstream stream-send [recipient] [amount] --end-time <end-timestamp> 
```
**example**:
```bash=
payment-streamd tx paymentstream stream-send streampay16qg7gpgt6hv9hqwrrk82r0f4kutqpy5zf03yx7 10000stake --end-time 1638786850 --chain-id sp-test-1 --from user1 --keyring-backend test
```

`Note: Use  --delayed flag to create delayed payment stream`

### Queries
```bash=
payment-streamd q paymentstream payment-streams -h
```
```bash=
Query payment streams.

Usage:
  payment-streamd query paymentstream payment-streams  [flags]

Examples:
$ payment-streamd query payment-stream payment-streams <id>
```

---

## Launch chain using starport
```
cd payment-stream
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Usage of CLI Commands

To Start a payment stream

cmd :

 `spd tx paymentstream stream-send [recipient] [amount] --end-time <unix-timestamp> --delayed --chain-id <chain-id> --from <key>`

To start a continuous payment stream
```bash=
spd tx paymentstream stream-send streampay1vnlgxmzh8mr5e43ku38f9470p2q0jfscksa98g 10000stake --end-time 1638786850  --chain-id paymentstream  --from bob
```
Use --delayed flag for delayed payments.

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
