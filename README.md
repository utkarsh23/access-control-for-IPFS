# Access Control For IPFS
This is an attempt to provide access control for IPFS using the ethereum blockchain.
## Requirements
* [go-ethereum](https://github.com/ethereum/go-ethereum)
* Blockchain API on [Infura](https://infura.io/)
* [Solidity Compiler](https://solidity.readthedocs.io/en/v0.5.3/installing-solidity.html)
## Compiling & Deploying The Smart Contract
1. Smart contract (.sol) gets compiled into a golang equivalent (.go) by executing the following.
```
abigen --sol=Access_Control_IPFS.sol --pkg=main --out=Access_Control_IPFS.go
```
2. The smart contract is deployed onto the Ethereum's Goerli test network. Infura is a service that an API to deploy & interact with your blockchain. First you need to copy an ethereum wallet's private key. This is how you can do it from metamask: [https://metamask.zendesk.com/hc/en-us/articles/360015289632-How-to-Export-an-Account-Private-Key](https://metamask.zendesk.com/hc/en-us/articles/360015289632-How-to-Export-an-Account-Private-Key)
```
WALLET_PRIVATE_KEY=<your wallet's private key> go run contract_deploy.go Access_Control_IPFS.go
```
## Testing Access Control With IPFS
We need to inject the access control programming logic into the IPFS source so that it interacts with the blockchain to provide access control over various files.

### Prerequisite
Install go 1.14.4:
```
curl -O https://storage.googleapis.com/golang/go1.14.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```
Install build essentials:
```
sudo apt install build-essential
```

This repository is one such example, but is extremely hacky and unsuitable for stable usage at this stage: https://github.com/utkarsh23/go-ipfs

1. Clone the above repo and change directory to the root of this repo.
```
git clone https://github.com/utkarsh23/go-ipfs
cd go-ipfs
```
2. Execute `make install`.
3. Export environment variables for ipfs.
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
4. Export the WALLET_PRIVATE_KEY variable.
```
export WALLET_PRIVATE_KEY=<your wallet's private key>
```
5. Initialize the IPFS node: `ipfs init`
6. Take your node online: `ipfs daemon`
