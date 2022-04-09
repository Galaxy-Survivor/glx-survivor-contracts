const hre = require("hardhat");

async function main() {
  const USDC = await hre.ethers.getContractFactory("USDC");
  const usdc = await USDC.deploy();

  await usdc.deployed();

  console.log("USD coin deployed to:", usdc.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
