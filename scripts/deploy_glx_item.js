const hre = require("hardhat");

async function main() {
  const GLXItem = await hre.ethers.getContractFactory("GLXItem");
  const glxItem = await GLXItem.deploy("https://test.galaxy-survivor.com/api/items/");

  await glxItem.deployed();

  console.log("GLXItem deployed to:", glxItem.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
