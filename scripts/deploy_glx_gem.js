const hre = require("hardhat");

async function main() {
  const GLXGem = await hre.ethers.getContractFactory("GLXGem");
  const glxGem = await GLXGem.deploy("https://nft.galaxysurvivor.xyz/gems/", "0");

  await glxGem.deployed();

  console.log("GLXGem deployed to:", glxGem.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
