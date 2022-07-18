const hre = require("hardhat");

async function main() {
  const GLXEquipment = await hre.ethers.getContractFactory("GLXEquipment");
  const glxEquipment = await GLXEquipment.deploy("https://nft.galaxysurvivor.xyz/equipments/", "0");

  await glxEquipment.deployed();

  console.log("GLXEquipment deployed to:", glxEquipment.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
