cd /home/shivam/go/src/github.com/goyal-shivam/fabric-samples-older/test-network

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

i=1
while [ "$i" -le 6 ]; do
    # printf $i
    # printf " - Next execution of MakeDoor Smart Contract\n"

    peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"MakeDoor","Args":["0.5", "1.0"]}'

    sleep 2.5

    # printf "\n\nOutput of GetAllAssets - "

    # peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAssets"]}'
    # printf "\nEnded this loop execution\n\n"
    i=$(( i + 1 ))
done 


