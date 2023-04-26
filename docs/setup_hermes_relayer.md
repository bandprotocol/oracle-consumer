# Connecting Hermes Relayer to BandChain
This document describes methods on how to set up Hermes Relayer to send IBC messages between oracle-consumer chain and BandChain.

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

##### Create mnemonic file on consumer chain and BandChain
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

###### BandChain key

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
