

# install
docker exec cli peer chaincode install -n reitscc -v 1.0 -p github.com/reitscc

# instantiate
docker exec cli peer chaincode instantiate -n reitscc -v 1.0 -C mychannel -c '{"Args":["user0","100","100","project0","10000","10000"]}' -P 'OR ("Org1MSP.member","Org2MSP.member","Org3MSP.member")'
sleep 5

# addUser
docker exec cli peer chaincode invoke -n reitscc -C mychannel -c '{"Args":["addUser","user1"]}'
sleep 5

# invoke
docker exec cli peer chaincode invoke -n reitscc -C mychannel -c '{"Args":["invest","user0","project0","50"]}'
sleep 5

# query user0
docker exec cli peer chaincode query -n reitscc -C mychannel -c '{"Args":["query","user0"]}'
sleep 5

# query user1
docker exec cli peer chaincode query -n reitscc -C mychannel -c '{"Args":["query","user1"]}'
sleep 5

