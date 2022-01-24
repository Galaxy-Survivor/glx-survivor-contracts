const hre = require("hardhat");

async function main() {
  const GLXMarket = await hre.ethers.getContractFactory("GLXMarket");
  const glxMarket = await GLXMarket.deploy(["0x0000000000000000000000000000000000000000"]);

  await glxMarket.deployed();

  console.log("GLXMarket deployed to:", glxMarket.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
