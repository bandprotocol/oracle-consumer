# Consumer Chain

The Consumer is an application of the Cosmos SDK that demonstrates the use of the [pricefeed-module](https://) implemented by BandProtocol. This module allows other Cosmos SDK applications to easily obtain data from BandChain through IBC.

- `consumerd`: The Consumer Daemon and command-line interface (CLI). runs a full-node of the comsumer application. 

consumer is built on the Cosmos SDK using the following modules:

- `x/consumer`: Consume data from pricefeed module.
- `x/pricefeed`: Reuqest BandChain logic.

### Running Consumer Chain
This guide will provide instructions on how to install the `consumerd` binary and run the cli on a server. By following these steps, you will have the binary properly set up and ready to use.

#### Installation

##### Install Go on Ubuntu

At the time of this writing, the latest release is `1.19.5`. We're going to download the tarball, extract it to `/usr/local`, and export `GOROOT` to our ``$PATH`

```
curl -OL https://golang.org/dl/go1.19.5.linux-amd64.tar.gz

sudo tar -C /usr/local -xvf go1.19.5.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin
```

Remember to add `GOPATH` to your `$PATH` environment variable. If you're not sure where that is, run go env GOPATH. This will allow us to run the gaiad binary in the next step.

```
export PATH=$PATH:$(go env GOPATH)/bin
```

##### Install ignite CLI
Ignite CLI is an easy-to-use CLI tool for creating and maintaining sovereign application-specific blockchains. Blockchains created with Ignite CLI use Cosmos SDK and Tendermint. 

To install the `ignite` binary in `/usr/local/bin` run the following command:
```
curl https://get.ignite.com/cli! | bash
```


##### Install the binaries

```
git clone -b <latest-release-tag> https://github.com/bandprotocol/oracle-consumer.git
cd consumer && ignite chain serve -v
```

That will install the consumerd binary. Verify that everything installed successfully by running:

```
consumerd version --long
```

You should see something similar to the following:

```
build_tags: ""
commit: 69b96f207cafaab827c3ddea78e80303eca58dfd
cosmos_sdk_version: v0.46.7
go: go version go1.19.5 darwin/amd64
name: Consumer
server_name: consumerd
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
```yaml
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
id = 'consumer'
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

##### create mnemonic file
> Note: The consumer mnemonic can be either Alice's or Bob's mnemonic, which you will receive when you first launch the consumer chain. It's important to note that you need to have funds in your key in order to send transactions on each chain.
- mem-consumer.txt -> add your consumer mnemonic
- mem-band.txt -> add your BandChain mnemonic

##### add keys to hermes by following command

###### consumer key

```
target/release/hermes --config config_relayer.toml keys add --chain consumer --mnemonic-file "mem-consumer.txt"
```

###### and BandChian key

```
target/release/hermes --config config_relayer.toml keys add --chain band-laozi-testnet6 --mnemonic-file "mem-band.txt"  --hd-path "m/44'/494'/0'/0/0"
```

#### Create client connection

```
target/release/hermes --config config_relayer.toml create channel --a-chain band-laozi-testnet6 --b-chain consumer --a-port oracle --b-port pricefeed --order unordered --channel-version bandchain-1 --new-client-connection
```

#### Start hermes relayer

```
target/release/hermes --config config_relayer.toml start
```


### Open and Vote The Proposal

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

#### Submmit proposal

```
consumerd tx gov submit-legacy-proposal update-symbol-request proposal.json --from alice
```

#### Vote the proposal

```
consumerd tx gov vote 1 yes --from alice
```

```
consumerd tx gov vote 1 yes --from bob
```

#### Check proposal status

```
consumerd query gov proposals
```


### Query latest Price that got from BandChain

Query latrst price that got from BandChin via pricefeed moudle

```
consumerd query pricefeed price BTC
```

Query latrst price that got from BandChin via consumer moudle

```
consumerd query consumer price BTC
```
