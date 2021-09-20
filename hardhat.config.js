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
                '144da7bac727c34f519ab0f86ff61a1bfd377070264e0925243793825c3cb96d',
                'b6abf32df0a8becfddd8637b80801140b81b3687f010316351891ef5efed3cd5'
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