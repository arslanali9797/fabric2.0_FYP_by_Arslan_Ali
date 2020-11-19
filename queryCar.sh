set -ex

. ./envVar.sh


CC_RUNTIME_LANGUAGE="golang" # chaincode runtime language is node.js
CC_SRC_PATH="./go/"
CHANNEL_NAME="mychannel"
VERSION="1"
CC_NAME="fabcar"

chaincodeQuery(){
      setGlobalsForPeer0Org1

      # Query all cars
       peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"function":"CalculateScore","Args":["5440503812223"]}'

       #queryAllCars

      # Query Car by Id
      #peer chaincode query -C mychannel -n users -c '{"function": "QueryAllBossId","Args":[""]}'
    #'{"Args":["GetSampleData","Key1"]}'
}
chaincodeQuery
