# StreamPay Application
StreamPay is a dApp that's built on the Stream Pay Testnet (which utilizes the `streampay` module. This sovereign chain is powereed by the token `uspay` or `SPAY`.

While users need to pay the fee to start a payment stream in SPAY tokens, any token (native or IBC tokens) can be streamed between accounts.

Try the app on - https://sp.OmniFlix.network

# streampay module
**streampay** is a module built using Cosmos SDK, Tendermint and [Starport](https://github.com/tendermint/starport) to stream payments from on address to the other, using `delayed` and `continuous` payments types inspired by the vesting model in the `auth` module of the Cosmos SDK.

## Installation

### Requirements

| Requirement | Notes                                |
|-------------|--------------------------------------|
| Go version  | [Go1.19](https://go.dev/doc/install) |
| Cosmos SDK  | v0.45.12                             |

### Get source code & Install

```bash=
git clone https://github.com/OmniFlix/streampay.git
cd streampay
go mod tidy
make install
```
check installation
```bash=
streampayd version
```


## Launch chain manually

```bash=
# Delete previous data
rm -rf ~/.streampay/config/*

#Init node	
streampayd unsafe-reset-all
streampayd init "sp-node"  --chain-id "sp-test-1"

# Add keys
streampayd keys add validator --keyring-backend test
streampayd keys add user1 --keyring-backend test
streampayd keys add user2 --keyring-backend test

# Add genesis accounts
streampayd add-genesis-account $(streampayd keys show validator -a --keyring-backend test) 1000000000stake
streampayd add-genesis-account $(streampayd keys show user1 -a --keyring-backend test) 1000000000000stake

# Create gentx
streampayd gentx validator 10000000stake --moniker "validator-1" --chain-id "sp-test-1" --keyring-backend test

# Collect Gentxs
streampayd collect-gentxs
streampayd validate-genesis

# Start Chain with default config
streampayd start
```
### CLI Commands

### Transactions
**stream-send** - to start a payment stream
```bash=
$ streampayd tx streampay stream-send -h
```
```bash=
creates a stream payment

Usage:
  streampayd tx streampay stream-send [flags]

Examples:
$ streampayd tx streampay stream-send [recipient] [amount] --end-time <end-timestamp> 
```
**example**:
```bash=
streampayd tx streampay stream-send streampay16qg7gpgt6hv9hqwrrk82r0f4kutqpy5zf03yx7 10000stake --end-time 1638786850 --chain-id sp-test-1 --from user1 --keyring-backend test
```

`Note: Use  --delayed flag to create delayed stream payment`

### Queries
```bash=
streampayd q streampay stream-payments -h
```
```bash=
Query stream payments.

Usage:
  streampayd query streampay stream-payments  [flags]

Examples:
$ streampayd query streampay stream-payments <id>
```

---

## Launch chain using starport
```
cd streampay
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Usage of CLI Commands

To Start a stream payment

cmd :

 `streampayd tx streampay stream-send [recipient] [amount] --end-time <unix-timestamp> --delayed --chain-id <chain-id> --from <key>`

To start a continuous payment stream
```bash=
streampayd tx streampay stream-send streampay1vnlgxmzh8mr5e43ku38f9470p2q0jfscksa98g 10000stake --end-time 1638786850  --chain-id streampay  --from bob
```
Use --delayed flag for delayed payments.

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
