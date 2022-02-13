const hre = require("hardhat");

async function main() {
  const GLXLogic = await hre.ethers.getContractFactory("GLXLogic");
  const glxLogic = await GLXLogic.deploy(
    "0x5226AE7C172f99B090fA7c6D76B604bf20e0c2c7",
    "0x43348F0803D2BC9251F2A338F15A1eF74B0A3D64",
    "0xcC96d690B4D029E605C0B8aC517fb719911a6719"
  );

  await glxLogic.deployed();

  console.log("GLXLogic deployed to:", glxLogic.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
