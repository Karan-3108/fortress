#!/bin/bash

KEY="mykey"
CHAINID="fortress_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t fortress-datadir.XXXXX)

echo "create and add new keys"
./fortressd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Fortress with moniker=$MONIKER and chain-id=$CHAINID"
./fortressd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./fortressd add-genesis-account \
"$(./fortressd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000afortress,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./fortressd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./fortressd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./fortressd validate-genesis --home $DATA_DIR

echo "starting fortress node $i in background ..."
./fortressd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started fortress node"
tail -f /dev/null