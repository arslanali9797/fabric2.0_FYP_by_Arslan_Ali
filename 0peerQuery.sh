set -ex

. ./envVar.sh


CC_RUNTIME_LANGUAGE=node # chaincode runtime language is node.js
CC_SRC_PATH="./chaincode/"
CHANNEL_NAME="mychannel"
VERSION="1"
CC_NAME="fabcar"

chaincodeQuery(){
      setGlobalsForPeer0Org1

      # Query all cars
       peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"Args":["queryAllCars"]}'

      # Query Car by Id
      #peer chaincode query -C mychannel -n users -c '{"function": "QueryAllBossId","Args":[""]}'
    #'{"Args":["GetSampleData","Key1"]}'
}
chaincodeQuery
