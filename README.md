# Oracle Consumer Chain

The oracle-consumer is an application of the Cosmos SDK that demonstrates the use of the [pricefeed-module](https://) implemented by BandProtocol. This module allows other Cosmos SDK applications to easily obtain data from BandChain through IBC.

- `oracle-consumerd`: The Oracle Consumer Daemon and command-line interface (CLI). Runs a full-node of the consumer application. 

oracle-consumer is built on the Cosmos SDK using the following modules:

- `x/consumer`: Consume data from pricefeed module.
- `x/pricefeed`: Logic of requesting data from BandChain.

### Prerequisites
Be sure you have met the prerequisites before you install and use Consumer Chain

#### Operating systems
- Ubuntu 22.04

#### Go (for ignite and consumer chain)
- 1.19.5 or higher

#### Rust (for Hermes Relayer)
- 1.65.0 or higher



### Running Consumer Chain
This guide will provide instructions on how to install the `consumerd` binary and run the cli on a server. By following these steps, you will have the binary properly set up and ready to use.


##### Install ignite CLI
Ignite CLI is an easy-to-use CLI tool for creating and maintaining sovereign application-specific blockchains. Blockchains created with Ignite CLI use Cosmos SDK and Tendermint. 

To install the `ignite` binary in `/usr/local/bin` run the following command:
```
curl https://get.ignite.com/cli! | bash
```

##### Installing build-essential package in Ubuntu 

```
sudo apt update && sudo apt install build-essential
```

##### Install the binaries and run consumer chain

```
git clone -b <latest-release-tag> https://github.com/bandprotocol/oracle-consumer.git
cd oracle-consumer & ignite chain serve -v
```

You can locate the `<latest-release-tag>` [here](https://github.com/bandprotocol/oracle-consumer/releases).

> When you initiate the chain, the Alice and Bob mnemonic will be displayed in your logging. Kindly preserve this mnemonic for use in the section below.

```
[IGNITE] ✔ Added account alice with address cosmos1d2dymcelkm6qgfaje3t4xt8vlpqfn227nvu9ja and mnemonic:
[IGNITE]   [your alice mnemonic]
[IGNITE] ✔ Added account bob with address cosmos1cn9yzz5p7scdc7s0he7vujyfcvtldlnv86tnys and mnemonic:
[IGNITE]   [your bob mnemonic]
```

By executing the command `ignite chain serve -v`, the consumerd binary will be installed and your consumer chain will start running.

Verify that everything installed successfully by running:

```
oracle-consumerd version --long
```

You should see something similar to the following:

```
...
build_tags: ""
commit: ...
cosmos_sdk_version: v0.46.11
go: go version go1.19.5 darwin/amd64
name: Oracle-Consumer
server_name: oracle-consumerd
version: ""
```

### Connecting Hermes Relayer to BandChain
The last step is to set up a relayer to listen and relay IBC packets between a oracle-consumer chain and BandChain.

Here are the simple guides for setting up a relayer.
* [Setup hermes relayer](docs/setup_hermes_relayer.md)


#### Add relayer key

##### create mnemonic file on consumer chain and BandChain
> Note: Upon logging for the first time that you run consumer chain, you will be given either Alice's or Bob's mnemonic, which can be used as the consumer mnemonic. It's important to note that you need to have funds in your key in order to send transactions on each chain.


```
# hermes/
touch mem-oracleconsumer.txt
```
- add your oracle-consumer chain mnemonic in mem-consumer.txt

```
# hermes/
touch mem-band.txt
```
- add your BandChain mnemonic in mem-band.txt

##### add keys to hermes by following command

###### consumer key

```
target/release/hermes --config config_relayer.toml keys add --chain oracleconsumer --mnemonic-file "mem-oracleconsumer.txt"
```

###### and BandChian key

```
target/release/hermes --config config_relayer.toml keys add --chain band-laozi-testnet6 --mnemonic-file "mem-band.txt"  --hd-path "m/44'/494'/0'/0/0"
```

#### Create client connection

```
target/release/hermes --config config_relayer.toml create channel --a-chain band-laozi-testnet6 --b-chain oracleconsumer --a-port oracle --b-port pricefeed --order unordered --channel-version bandchain-1 --new-client-connection
```

#### Start hermes relayer

```
target/release/hermes --config config_relayer.toml start
```


### Open and Vote the update symbol requests proposal
The purpose of this proposal is to request price data from BandChain at `block_interval` specified in the proposal. If the proposal is approved, the pricefeed module will retrieve the data and store the response on the consumer chain.

#### create proposal.json

```json
{
    "title": "Update Symbol requests",
    "description": "Update symbol that request price from BandChain",
    "symbol_requests": [
        {
            "symbol": "BTC",
            "oracle_script_id": "396",
            "block_interval": "40"
        },
        {
            "symbol": "ETH",
            "oracle_script_id": "396",
            "block_interval": "40"
        }
    ],
    "deposit": "10000000stake"
}
```

#### Submit proposal

```
oracle-consumerd tx gov submit-legacy-proposal update-symbol-request simple_proposals/update-symbol-requests.json --from alice
```

#### Vote the proposal

```
oracle-consumerd tx gov vote 1 yes --from alice
```

```
oracle-consumerd tx gov vote 1 yes --from bob
```

#### Check proposal status

```
oracle-consumerd query gov proposals
```


### Query latest price that got from BandChain

Query latest price that got from BandChin via pricefeed module

```
oracle-consumerd query pricefeed price BTC
```

Query latest price that got from BandChin via consumer module

```
oracle-consumerd query consumer price BTC
```
