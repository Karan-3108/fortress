KEY="mykey"
CHAINID="fortress_9000-1"
MONIKER="localtestnet"
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# Reinstall daemon
rm -rf ~/.fortressd*
make install

# Set client config
fortressd config keyring-backend $KEYRING
fortressd config chain-id $CHAINID

# if $KEY exists it should be deleted
fortressd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO

# Set moniker and chain-id for Fortress (Moniker can be anything, chain-id must be an integer)
fortressd init $MONIKER --chain-id $CHAINID

# Change parameter token denominations to afortress
cat $HOME/.fortressd/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="afortress"' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json
cat $HOME/.fortressd/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="afortress"' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json
cat $HOME/.fortressd/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="afortress"' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json
cat $HOME/.fortressd/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="afortress"' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json
cat $HOME/.fortressd/config/genesis.json | jq '.app_state["inflation"]["params"]["mint_denom"]="afortress"' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json

# Set gas limit in genesis
cat $HOME/.fortressd/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json

# Set claims start time
node_address=$(fortressd keys list | grep  "address: " | cut -c12-)
current_date=$(date -u +"%Y-%m-%dT%TZ")
cat $HOME/.fortressd/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["airdrop_start_time"]=$current_date' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json

# Set claims records for validator account
amount_to_claim=10000
#cat $HOME/.fortressd/config/genesis.json | jq -r --arg node_address "$node_address" --arg amount_to_claim "$amount_to_claim" '.app_state["claims"]["claims_records"]=[{"initial_claimable_amount":$amount_to_claim, "actions_completed":[false, false, false, false],"address":fortress1dlrx2we7hdrak87x7mxu2jxfg7j2rf9rpymtd8}]' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json

# Set claims decay
cat $HOME/.fortressd/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["duration_of_decay"]="1000000s"' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json
cat $HOME/.fortressd/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["duration_until_decay"]="100000s"' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json

# Claim module account:
# 0xA61808Fe40fEb8B3433778BBC2ecECCAA47c8c47 || fortress1dlrx2we7hdrak87x7mxu2jxfg7j2rf9rpymtd8
cat $HOME/.fortressd/config/genesis.json | jq -r --arg amount_to_claim "$amount_to_claim" '.app_state["bank"]["balances"] += [{"address":"fortress1dlrx2we7hdrak87x7mxu2jxfg7j2rf9rpymtd8","coins":[{"denom":"afortress", "amount":$amount_to_claim}]}]' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json

# disable produce empty block
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.fortressd/config/config.toml
  else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.fortressd/config/config.toml
fi

if [[ $1 == "pending" ]]; then
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.fortressd/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.fortressd/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.fortressd/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.fortressd/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.fortressd/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.fortressd/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.fortressd/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.fortressd/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.fortressd/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.fortressd/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.fortressd/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.fortressd/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.fortressd/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.fortressd/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.fortressd/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.fortressd/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.fortressd/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.fortressd/config/config.toml
  fi
fi

# Allocate genesis accounts (cosmos formatted addresses)
fortressd add-genesis-account $KEY 100000000000000000000000000afortress --keyring-backend $KEYRING

# Update total supply with claim values
validators_supply=$(cat $HOME/.fortressd/config/genesis.json | jq -r '.app_state["bank"]["supply"][0]["amount"]')
# Bc is required to add this big numbers
# total_supply=$(bc <<< "$amount_to_claim+$validators_supply")
total_supply=100000000000000000000010000
cat $HOME/.fortressd/config/genesis.json | jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' > $HOME/.fortressd/config/tmp_genesis.json && mv $HOME/.fortressd/config/tmp_genesis.json $HOME/.fortressd/config/genesis.json

# Sign genesis transaction
fortressd gentx $KEY 1000000000000000000000afortress --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
fortressd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
fortressd validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
fortressd start --pruning=nothing $TRACE --log_level $LOGLEVEL --minimum-gas-prices=0.0001afortress --json-rpc.api eth,txpool,personal,net,debug,web3
