const hre = require("hardhat");

async function main() {
  const glxLogic = "0xe44a252435Ad21b5aECBE5BDEAB648f668095d52";

  const GLXShip = await hre.ethers.getContractFactory("GLXShip");
  const glxShip = await GLXShip.attach("0x5226AE7C172f99B090fA7c6D76B604bf20e0c2c7");
  const tx1 = await glxShip.addMinter(glxLogic);
  console.log("Add minter for GLXShip:", tx1.hash);

  const GLXEquipment = await hre.ethers.getContractFactory("GLXEquipment");
  const glxEquipment = await GLXEquipment.attach("0x43348F0803D2BC9251F2A338F15A1eF74B0A3D64");
  const tx2 = await glxEquipment.addMinter(glxLogic);
  console.log("Add minter for GLXEquipment:", tx2.hash);

  const GLXItem = await hre.ethers.getContractFactory("GLXItem");
  const glxItem = await GLXItem.attach("0xcC96d690B4D029E605C0B8aC517fb719911a6719");
  const tx3 = await glxItem.addMinter(glxLogic);
  console.log("Add minter for GLXItem:", tx3.hash);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
