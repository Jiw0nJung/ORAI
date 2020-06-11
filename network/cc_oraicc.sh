ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/orai.com/orderers/orderer.orai.com/msp/tlscacerts/tlsca.orai.com-cert.pem

# install
docker exec cli peer chaincode install -n oraicc -v 1.0 -p github.com/oraicc
sleep 5

# instantiate
docker exec cli peer chaincode instantiate -o orderer.orai.com:7050 --tls --cafile $ORDERER_CA -C orai -n oraicc -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member','Org3MSP.member')"
sleep 10

# addUser
docker exec cli peer chaincode invoke -o orderer.orai.com:7050 --tls --cafile $ORDERER_CA -C orai -n oraicc -C orai -c '{"Args":["addUser","jiwon","audi a8","audi"]}'
sleep 5

# invoke
docker exec cli peer chaincode invoke -o orderer.orai.com:7050 --tls --cafile $ORDERER_CA -C orai -n oraicc -C orai -c '{"Args":["addAccidents","jiwon","raining","daylight","wet","flooded","backing","rearend"]}'
sleep 5

# query user0
docker exec cli peer chaincode query -n oraicc -C orai -c '{"Args":["viewAccidents","jiwon"]}'
sleep 5