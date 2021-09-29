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
                '01bce67e184a9009d3da19f260f16a69f5a9733a746f81afd6dabdb991a812ea',
                '7e3b88f484c531d2d10439ed24449f2e0eee3315869522dc6712e844cea2fd2e'
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