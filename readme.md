# paymentstream
**paymentstream** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport) to stream payments from on addrs to the other, using delayed and continuous payments types inspired by Vesting types from Cosmos SDK. 

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## CLI Example

To Start a payment stream
cmd : `spd tx paymentstream stream-send [recipient] [amount] --end-time <unix-timestamp> --delayed <bool> --chain-id <chain-id> --from <key>`
```shell
spd tx paymentstream stream-send streampay1vnlgxmzh8mr5e43ku38f9470p2q0jfscksa98g 10000stake --end-time 1638786850 --delayed false --chain-id paymentstream --from bob
```
Use --delayed false for continuous payments.

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/OmniFlix/payment-stream@latest! | sudo bash
```
`OmniFlix/payment-stream` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
