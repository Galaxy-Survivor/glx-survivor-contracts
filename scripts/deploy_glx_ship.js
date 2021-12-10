const hre = require("hardhat");

async function main() {
  const GLXShip = await hre.ethers.getContractFactory("GLXShip");
  const glxShip = await GLXShip.deploy("https://test.galaxy-survivor.com/api/ship/");

  await glxShip.deployed();

  console.log("GLXShip deployed to:", glxShip.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
