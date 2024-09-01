#!/bin/bash

ABIS_DIR="contracts/abis"
BINDINGS_DIR="contracts"

# create the bindings directory if it doesn't exist
mkdir -p $BINDINGS_DIR

for abi_file in $ABIS_DIR/*.abi; do

    base_name=$(basename -- "$abi_file")
    contract_name="${base_name%.abi}"

    mkdir -p $BINDINGS_DIR/$contract_name

    abigen --abi "$abi_file" --pkg $contract_name --out "$BINDINGS_DIR/$contract_name/$contract_name.go"
done