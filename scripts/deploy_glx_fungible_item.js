const hre = require("hardhat");

async function main() {
  const GLXFungibleItem = await hre.ethers.getContractFactory("GLXFungibleItem");
  const glxFungibleItem = await GLXFungibleItem.deploy();

  await glxFungibleItem.deployed();

  console.log("GLXFungibleItem deployed to:", glxFungibleItem.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
