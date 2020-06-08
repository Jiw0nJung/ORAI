

# install
docker exec cli peer chaincode install -n oraicc -v 1.0 -p github.com/oraicc

# instantiate
docker exec cli peer chaincode instantiate -n oraicc -v 1.0 -C orai -c '{"Args":[""]}' -P 'OR ("Org1MSP.member","Org2MSP.member","Org3MSP.member")'
sleep 5

# addUser
docker exec cli peer chaincode invoke -n oraicc -C orai -c '{"Args":["addUser","jiwon","audi a8","audi"]}'
sleep 5

# invoke
docker exec cli peer chaincode invoke -n oraicc -C orai -c '{"Args":["addData","jiwon","raining","daylight","wet","flooded","backing","rearend"]}'
sleep 5

# query user0
docker exec cli peer chaincode query -n oraicc -C orai -c '{"Args":["readData","jiwon"]}'
sleep 5