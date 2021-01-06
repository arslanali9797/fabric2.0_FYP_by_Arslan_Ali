


//,'4550503812444','car loan','org1MSP','azhar' ,'durani',100000,'2/3/2020',2,50000,'28');
module.exports.loanClientInfo = function(cnic,loanType,loanAmount,yearPlan,depositeDate,callback){




    /*
     * SPDX-License-Identifier: Apache-2.0
     */
    
    'use strict';
    
    
    const { Gateway, Wallets } = require('fabric-network');
    const fs = require('fs');
    const path = require('path');
    
    async function main() {
        try {
            // load the network configuration
            const ccpPath = path.resolve(__dirname, '..',  'connection', 'connection-org1.json');
            let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));
    
            // Create a new file system based wallet for managing identities.
            const walletPath = path.join(process.cwd(), 'wallet');
            const wallet = await Wallets.newFileSystemWallet(walletPath);
            console.log(`Wallet path: ${walletPath}`);
    
            // Check to see if we've already enrolled the user.
            const identity = await wallet.get('appUser');
            if (!identity) {
                console.log('An identity for the user "appUser" does not exist in the wallet');
                console.log('Run the registerUser.js application before retrying');
                return;
            }
    
            // Create a new gateway for connecting to our peer node.
            const gateway = new Gateway();
            await gateway.connect(ccp, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });
    
            // Get the network (channel) our contract is deployed to.
            const network = await gateway.getNetwork('mychannel');
    
            // Get the contract from the network.
            const contract = network.getContract('fabcar');

            await contract.submitTransaction('ClientPrivateRecord',cnic,loanType,loanAmount,'24/10/2020',yearPlan,depositeDate,'org1loan');
            console.log('Transaction has been submitted');
            await gateway.disconnect();
            callback('200');

            
        } catch (error) {
            console.error(`Failed to submit transaction: ${error}`);
            process.exit(1);
        }
    }
    
    main();
    
}