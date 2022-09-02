# Conga NFT Contract

Proof of concept containerised Fabric chaincode for use with [fabric-builder-k8s](https://github.com/hyperledgendary/fabric-builder-k8s)

Based on the [token-erc-721 example in fabric-samples](https://github.com/hyperledger/fabric-samples/tree/main/token-erc-721/chaincode-go)

See [fabric-builder-k8s samples](https://github.com/hyperledger-labs/fabric-builder-k8s/tree/main/samples) for more recent examples in Go, Java, and Node.js

## Dev container

[![Open in Remote - Containers](https://img.shields.io/static/v1?label=Remote%20-%20Containers&message=Open&color=blue&logo=visualstudiocode)](https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/hyperledgendary/conga-nft-contract)

The badge above should launch this contract inside a prototype Fabric dev container

Once it has started, there should be a running [microfab](https://github.com/IBM-Blockchain/microfab) network, which you can use to deploy the contract to as follows

Configure the `peer` command environment

```shell
export FABRIC_LOGGING_SPEC=INFO
export CORE_PEER_TLS_ENABLED=false
export CORE_PEER_MSPCONFIGPATH=/var/opt/hyperledger/microfab/admin-org1
export FABRIC_CFG_PATH=/var/opt/hyperledger/microfab/peer-org1/config
export CORE_PEER_ADDRESS=$(yq .peer.address ${FABRIC_CFG_PATH}/core.yaml)
export CORE_PEER_LOCALMSPID=$(yq .peer.localMspId ${FABRIC_CFG_PATH}/core.yaml)
```

Create a chaincode-as-a-service package

```shell
~/fabric-samples/test-network/scripts/pkgcc.sh -l conga -a 0.0.0.0:9999
```

**Note:** ignore the error, which is due to a [known issue](https://github.com/hyperledger/fabric-samples/issues/823)

Install the chaincode package

```shell
peer lifecycle chaincode install conga.tgz
```

Export a PACKAGE_ID environment variable

```shell
export PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid conga.tgz) && echo $PACKAGE_ID
```

Approve the chaincode

```shell
peer lifecycle \
  chaincode approveformyorg \
  --channelID     channel1 \
  --name          conga-nft \
  --version       1 \
  --package-id    ${PACKAGE_ID} \
  --sequence      1 \
  --orderer       orderer-api.127-0-0-1.nip.io:8080
```

Commit the chaincode

```shell
peer lifecycle \
  chaincode commit \
  --channelID     channel1 \
  --name          conga-nft \
  --version       1 \
  --sequence      1 \
  --orderer       orderer-api.127-0-0-1.nip.io:8080
```

Add the `PACKAGE_ID` environment variable to the `.vscode/launch.json` file

```shell
yq e '(.configurations[] | select(.name == "Debug chaincode") | .env.PACKAGE_ID) = strenv(PACKAGE_ID)' -i .vscode/launch.json
```

Start the chaincode in debug using vscode!

Check the chaincode works!

```shell
peer chaincode query -C channel1 -n conga-nft -c '{"Args":["org.hyperledger.fabric:GetMetadata"]}'
```
