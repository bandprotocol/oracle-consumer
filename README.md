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


### Open and Vote the source channel param change proposal
The current default value for the source channel is `[not_set]`. If you wish to obtain BandChain data through IBC, you will need to open the proposal to change the source channel param to your own source channel. An example of how to open parameter change proposal is provided below.

#### create proposal.json
> Note: this example has been provided in `example/proposals/source-channel-params-change.json`

```json
{
  "title": "Param change for SourceChannel",
  "description": "Proposal for change SourceChannel param in pricefeed module",
  "changes": [
    {
      "subspace": "pricefeed",
      "key": "SourceChannel",
      "value": "channel-0"
    }
  ],
  "deposit": "10000000stake"
}
```

#### Submit proposal

```
oracle-consumerd tx gov submit-legacy-proposal param-change example/proposals/source-channel-params-change.json --from alice
```

#### Vote the proposal

```
oracle-consumerd tx gov vote 1 yes --from alice
```

```
oracle-consumerd tx gov vote 1 yes --from bob
```

### Open and Vote the update symbol requests proposal
The purpose of this proposal is to request price data from BandChain at `block_interval` specified in the proposal. If the proposal is approved, the pricefeed module will retrieve the data and store the response on the consumer chain.

#### create proposal.json
> Note: this example has been provided in `example/proposals/update-symbol-requests.json`

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
oracle-consumerd tx gov submit-legacy-proposal update-symbol-request example/proposals/update-symbol-requests.json --from alice
```

#### Vote the proposal

```
oracle-consumerd tx gov vote 2 yes --from alice
```

```
oracle-consumerd tx gov vote 2 yes --from bob
```

### Check proposal status

```
oracle-consumerd query gov proposals
```

### Option 2: Another way to initiate the symbol requests

To utilize the Ignite feature to replace the genesis state, insert the code shown below into the `config.yml` file.

```yml
genesis:
  app_state:
    pricefeed:
        symbol_requests: [{"symbol": "BAND", "oracle_script_id": 396, "block_interval":  40}]
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
