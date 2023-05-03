DIR=`dirname "$0"`/nodes

rm -r $DIR

make install

declare -a vals=("alice" "bob")

# init main node and account
oracle-consumerd init main --chain-id oracle-consumer --home $DIR/main
main_node=$(oracle-consumerd tendermint show-node-id --home $DIR/main)
oracle-consumerd keys add main --keyring-backend test --home $DIR/main
oracle-consumerd add-genesis-account main 10000000000000stake --keyring-backend test --home $DIR/main

for val in "${vals[@]}"
do
  # init node and store variable
  oracle-consumerd init $val --chain-id oracle-consumer --home $DIR/$val
  node_id=$(oracle-consumerd tendermint show-node-id --home $DIR/$val)
  node_pub=$(oracle-consumerd tendermint show-validator --home $DIR/$val)

  # create account with money
  oracle-consumerd keys add $val --keyring-backend test --home $DIR/main
  oracle-consumerd add-genesis-account $val 10000000000000stake --keyring-backend test --home $DIR/main

  # register validator
  oracle-consumerd gentx $val 100000000stake \
      --chain-id oracle-consumer \
      --node-id $node_id \
      --pubkey "$node_pub" \
      --moniker $val \
      --keyring-backend test \
      --home $DIR/main
done

# collect genesis transactions
oracle-consumerd collect-gentxs --home $DIR/main

# add persistent_peers to use main node
sed -i '' -e \
  "s/persistent_peers = .*/persistent_peers = \"$main_node@main:26656\"/" \
  $DIR/main/config/config.toml

cat <<< $(jq '.app_state.gov.voting_params.voting_period = "50s"' $DIR/main/config/genesis.json) > $DIR/main/config/genesis.json

for val in "${vals[@]}"
do
  # copy and paster config file and genesis
  cp $DIR/main/config/genesis.json $DIR/$val/config/genesis.json
  cp $DIR/main/config/app.toml $DIR/$val/config/app.toml
  cp $DIR/main/config/config.toml $DIR/$val/config/config.toml
done

# adjust configuration of main node
sed -i '' \
  '/\[api\]/,+10 s/enable = .*/enable = true/' \
  $DIR/main/config/app.toml

sed -i '' \
  '/\[rpc\]/,+10 s/laddr = .*/laddr = "tcp:\/\/0.0.0.0:26657"/' \
  $DIR/main/config/config.toml
