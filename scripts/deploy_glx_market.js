const hre = require("hardhat");

async function main() {
  const GLXMarket = await hre.ethers.getContractFactory("GLXMarket");
  const glxMarket = await GLXMarket.deploy(
    ["0x0000000000000000000000000000000000000000", "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"]
  );

  await glxMarket.deployed();

  console.log("GLXMarket deployed to:", glxMarket.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
