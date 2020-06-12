 # ORAI

![diagram](./img/network.png)



# hyperledger fabric sample 
## pre-condition
* curl, docker, docker-compose, go, nodejs, python
* hyperledger fabric-docker images are installed
* GOPATH are configured
* hyperledger bineries are installed (cryptogen, configtxgen ... etcs)
# -network
## 1. generating crypto-config directory, genesis.block, channel
cd network
./generate.sh
## 2. starting the network, create channel and join 
./start.sh
