#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

docker-compose -f docker-compose.yml down

docker-compose -f docker-compose.yml up -d

# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=10
#echo ${FABRIC_START_TIMEOUT}
sleep ${FABRIC_START_TIMEOUT}

# Create the channel
docker exec cli peer channel create -o orderer.orai.com:7050 -c orai -f /etc/hyperledger/configtx/channel.tx

sleep 5
# Join peer0.org1.user.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.user.com/msp" peer0.org1.user.com peer channel join -b /etc/hyperledger/configtx/orai.block
sleep 5

# Join peer0.org2.insurance.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=Org2MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org2.insurance.com/msp" peer0.org2.insurance.com peer channel join -b /etc/hyperledger/configtx/orai.block
sleep 5


# Join peer0.org3.manufacturer.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=Org3MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org3.manufacturer.com/msp" peer0.org3.manufacturer.com peer channel join -b /etc/hyperledger/configtx/orai.block
sleep 5




