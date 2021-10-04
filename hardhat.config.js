require("@nomiclabs/hardhat-waffle");

task("accounts", "Prints the list of accounts", async(taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners();

    for (const account of accounts) {
        console.log(account.address);
    }
});

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
    solidity: "0.8.7",
    settings: {
        optimizer: {
            enabled: true,
            runs: 200
        }
    },
    networks: {

        ganache: {
            url: "http://localhost:8545",
            accounts: [
                '6c20f08e5b8e9107266e905f3fad55ca0649c3815e43ea21618cc36f195b3802',
                '93d2c579f2c98a12d964fe66f1acba209a79def145cd8768104c459d50996371'
            ]
        },
        rinkeby: {
            "url": "https://rinkeby.infura.io/v3/833fc8ec67984732a831d6af176ed7a4",
            "accounts": [`0x8e94548d06a648da86dd3766ba28e768a5eca79c0f7e68a8e97914845ebfd27f`]
        },
        bsctestnet: {
            "url": "https://data-seed-prebsc-1-s1.binance.org:8545/",
            "accounts": [`0x8e94548d06a648da86dd3766ba28e768a5eca79c0f7e68a8e97914845ebfd27f`]
        }
    }
};