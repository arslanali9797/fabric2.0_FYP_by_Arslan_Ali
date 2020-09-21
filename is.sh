set -ex

. ./envVar.sh


CC_RUNTIME_LANGUAGE=node # chaincode runtime language is node.js
CC_SRC_PATH="./chaincode/"
CHANNEL_NAME="mychannel"
VERSION="1"
CC_NAME="fabcar"



checkCommitReadyness(){
    setGlobalsForPeer0Org1
    peer lifecycle chaincode checkcommitreadiness \
    --channelID $CHANNEL_NAME --name ${CC_NAME} --version ${VERSION} \
    --sequence ${VERSION} --output json --init-required
    echo "===================== checking commit readyness from peer0 org 1 ===================== "
}
checkCommitReadyness






 approveForMyOrg1(){
     setGlobalsForPeer0Org1
     peer lifecycle chaincode approveformyorg -o localhost:7050 \
     --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED \
     --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name ${CC_NAME} \
     --version ${VERSION} --init-required --package-id ${PACKAGE_ID} \
     --sequence ${VERSION}

     echo "===================== chaincode approved from org 1 ===================== "
 }
approveForMyOrg1

checkCommitReadyness(){
    echo "===================== checking commit readyness  from peer0 org 1 ===================== "
    setGlobalsForPeer0Org1
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA --name ${CC_NAME} --version ${VERSION} --sequence ${VERSION} --output json --init-required
    echo "===================== checked commit readyness  from peer0 org 1 ===================== "
}
checkCommitReadyness

checkCommitReadyness(){
    echo "===================== checking commit readyness  from peer1 org 1 ===================== "
    setGlobalsForPeer1Org1
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA --name ${CC_NAME} --version ${VERSION} --sequence ${VERSION} --output json --init-required
    echo "===================== checked commit readyness  from peer1 org 1 ===================== "
}
checkCommitReadyness






approveForMyOrg2(){
    setGlobalsForPeer0Org2
    peer lifecycle chaincode approveformyorg -o localhost:7050 \
    --ordererTLSHostnameOverride orderer.example.com --tls $CORE_PEER_TLS_ENABLED \
    --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name ${CC_NAME} \
    --version ${VERSION} --init-required --package-id ${PACKAGE_ID} \
    --sequence ${VERSION}
    echo "===================== chaincode approved from org 2 ===================== "
}
approveForMyOrg2

checkCommitReadyness(){
    echo "===================== checking commit readyness from peer0 org 2 ===================== "
    setGlobalsForPeer0Org2
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA --name ${CC_NAME} --version ${VERSION} --sequence ${VERSION} --output json --init-required
    echo "===================== checked commit readyness from peer0 org 2 ===================== "
}
checkCommitReadyness

checkCommitReadyness(){
    echo "===================== checking commit readyness from peer1 org 2 ===================== "
    setGlobalsForPeer1Org2
    peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA --name ${CC_NAME} --version ${VERSION} --sequence ${VERSION} --output json --init-required
    echo "===================== checked commit readyness from peer1 org 2 ===================== "
}
checkCommitReadyness






commitChaincodeDefination(){
    setGlobalsForPeer0Org1
    peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com \
    --tls $CORE_PEER_TLS_ENABLED  --cafile $ORDERER_CA \
    --channelID $CHANNEL_NAME --name ${CC_NAME} \
    --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
    --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
    --version ${VERSION} --sequence ${VERSION} --init-required
}
commitChaincodeDefination




queryCommitted(){
    setGlobalsForPeer0Org1
    peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name ${CC_NAME}
}
queryCommitted



chaincodeInvokeInit(){
    setGlobalsForPeer0Org1
    peer chaincode invoke -o localhost:7050 \
    --ordererTLSHostnameOverride orderer.example.com \
    --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA \
    -C $CHANNEL_NAME -n ${CC_NAME} \
    --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
    --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
    --isInit -c '{"function":"initLedger","Args":[]}'
}
 chaincodeInvokeInit


 chaincodeQuery(){
       setGlobalsForPeer0Org1

       # Query all cars
        peer chaincode query -C $CHANNEL_NAME -n ${CC_NAME} -c '{"Args":["queryAllCars"]}'

       # Query Car by Id
       #peer chaincode query -C mychannel -n users -c '{"function": "QueryAllBossId","Args":[""]}'
     #'{"Args":["GetSampleData","Key1"]}'
 }

 # chaincodeQuery
