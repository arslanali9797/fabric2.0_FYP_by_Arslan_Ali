module.exports.addClient= function(CNIC,fname,lname,age,gender,martialStatus,assetsValue,depositPerMonth,withdrawPerMonth,savingYears,noTransaction,typeOfBussiness,savingAmount,callback){


/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');


async function main() {
    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..',  'connection', 'connection-org2.json');
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('org2appUser1');
        if (!identity) {
            console.log('An identity for the user "org2appUser1" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'org2appUser1', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('fabcar');

       
        await contract.submitTransaction('AddRecord',CNIC,fname,lname,age,gender,martialStatus,assetsValue,depositPerMonth,withdrawPerMonth,savingYears,noTransaction,typeOfBussiness,savingAmount,0,0,0,"org2MSP");
        console.log('Transaction has been submitted');
        await gateway.disconnect();
        callback('200');


    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        process.exit(1);
    }
}

main();
}