set -ex

. ./envVar.sh


#presetup(){
#    echo Vendoring Go dependencies ...
#    pushd ./go
#    GO111MODULE=on go mod vendor
#    popd
#    echo Finished vendoring Go dependencies
#}
#presetup



CC_RUNTIME_LANGUAGE="golang" # chaincode runtime language is node.js
CC_SRC_PATH="./go/"
CHANNEL_NAME="mychannel"
VERSION="1"
CC_NAME="fabcar"




packageChaincode() {
  rm -rf ${CC_NAME}.tar.gz
  setGlobalsForPeer0Org1
  peer lifecycle chaincode package ${CC_NAME}.tar.gz \
  --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} \
  --label ${CC_NAME}_${VERSION}
  echo "===================== Chaincode is packaged ===================== "
}
packageChaincode



installChaincode(){
    setGlobalsForPeer0Org1
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer0.org1 ===================== "

    setGlobalsForPeer1Org1
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer1.org1 ===================== "

    setGlobalsForPeer0Org2
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer0.org2 ===================== "

    setGlobalsForPeer1Org2
    peer lifecycle chaincode install ${CC_NAME}.tar.gz
    echo "===================== Chaincode is installed on peer1.org2 ===================== "

  }
installChaincode


queryInstalled(){
    setGlobalsForPeer0Org1
    peer lifecycle chaincode queryinstalled >&log.txt
    cat log.txt
    PACKAGE_ID=$(sed -n "/${CC_NAME}_${VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    echo PackageID is ${PACKAGE_ID}
    echo "===================== Query installed successful on peer0.org1 on channel ===================== "
}

 queryInstalled






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


 queryCommitted(){
     setGlobalsForPeer0Org2
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

  sleep 5
