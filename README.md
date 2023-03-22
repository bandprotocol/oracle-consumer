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
commit: f81a9b4646d6e627ad3bbe5c7f20446ee2615155
cosmos_sdk_version: v0.46.10
go: go version go1.19.5 darwin/amd64
name: Oracle-Consumer
server_name: oracle-consumerd
version: ""
```

### Connecting Hermes Relayer to BandChain

#### Build hermes for BandChain

```
# Clone Hermes version 1.1.0-band
git clone https://github.com/bandprotocol/hermes.git
cd hermes
git checkout v1.1.0-band

# Build Hermes
cargo build --release
```



#### Create config_relayer.toml
```
# The global section has parameters that apply globally to the relayer operation.
[global]

# Specify the verbosity for the relayer logging output. Default: 'info'
# Valid options are 'error', 'warn', 'info', 'debug', 'trace'.
log_level = 'trace'


# Specify the mode to be used by the relayer. [Required]
[mode]

# Specify the client mode.
[mode.clients]

# Whether or not to enable the client workers. [Required]
enabled = true

# Whether or not to enable periodic refresh of clients. [Default: true]
# Note: Even if this is disabled, clients will be refreshed automatically if
#      there is activity on a connection or channel they are involved with.
refresh = true

# Whether or not to enable misbehaviour detection for clients. [Default: false]
misbehaviour = true

# Specify the connections mode.
[mode.connections]

# Whether or not to enable the connection workers for handshake completion. [Required]
enabled = true

# Specify the channels mode.
[mode.channels]

# Whether or not to enable the channel workers for handshake completion. [Required]
enabled = true

# Specify the packets mode.
[mode.packets]

# Whether or not to enable the packet workers. [Required]
enabled = true

# Parametrize the periodic packet clearing feature.
# Interval (in number of blocks) at which pending packets
# should be eagerly cleared. A value of '0' will disable
# periodic packet clearing. [Default: 100]
clear_interval = 100

# Whether or not to clear packets on start. [Default: false]
clear_on_start = true

# Toggle the transaction confirmation mechanism.
# The tx confirmation mechanism periodically queries the `/tx_search` RPC
# endpoint to check that previously-submitted transactions
# (to any chain in this config file) have delivered successfully.
# Experimental feature. Affects telemetry if set to false.
# [Default: true]
tx_confirmation = true

# The REST section defines parameters for Hermes' built-in RESTful API.
# https://hermes.informal.systems/rest.html
[rest]

# Whether or not to enable the REST service. Default: false
enabled = true

# Specify the IPv4/6 host over which the built-in HTTP server will serve the RESTful
# API requests. Default: 127.0.0.1
host = '127.0.0.1'

# Specify the port over which the built-in HTTP server will serve the restful API
# requests. Default: 3000
port = 3000


# The telemetry section defines parameters for Hermes' built-in telemetry capabilities.
# https://hermes.informal.systems/telemetry.html
[telemetry]

# Whether or not to enable the telemetry service. Default: false
enabled = true

# Specify the IPv4/6 host over which the built-in HTTP server will serve the metrics
# gathered by the telemetry service. Default: 127.0.0.1
host = '127.0.0.1'

# Specify the port over which the built-in HTTP server will serve the metrics gathered
# by the telemetry service. Default: 3001
port = 3001

[[chains]]
id = 'oracleconsumer'
rpc_addr = 'http://localhost:26657'
grpc_addr = 'http://localhost:9090'
websocket_addr = 'ws://localhost:26657/websocket'
rpc_timeout = '10s'
account_prefix = 'cosmos'
key_name = 'requester'
store_prefix = 'ibc'
default_gas = 5000000
max_gas = 15000000
gas_price = { price = 0, denom = 'ustake' }
gas_multiplier = 1.1
max_msg_num = 20
max_tx_size = 209715
clock_drift = '20s'
max_block_time = '10s'
trusting_period = '10days'
trust_threshold = { numerator = '1', denominator = '3' }
address_type = { derivation = 'cosmos' }
ignore_port_channel = []
# [chains.packet_filter]
# policy = 'allow'
# list = [
#    ['wasm.*', '*'],
# ]

[[chains]]
id = 'band-laozi-testnet6'
rpc_addr = 'https://rpc.laozi-testnet6.bandchain.org:443'
grpc_addr = 'https://laozi-testnet6.bandchain.org:443'
websocket_addr = 'wss://rpc.laozi-testnet6.bandchain.org:443/websocket'
rpc_timeout = '10s'
account_prefix = 'band'
key_name = 'testkey'
store_prefix = 'ibc'
default_gas = 100000
max_gas = 10000000
gas_price = { price = 0.0025, denom = 'uband' }
gas_multiplier = 1.1
max_msg_num = 30
max_tx_size = 2097152
clock_drift = '5s'
max_block_time = '10s'
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }
address_type = { derivation = 'cosmos' }
ignore_port_channel = []
# [chains.packet_filter]
# policy = 'allow'
# list = [
#    ['oracle', '*'],
# ]
```

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
