const hre = require("hardhat");

async function main() {
  //const tokenAddress = "0x8bE9209D1f3d51C59D5BB3c5Cc30d8c266652B3f";
  //const account = "0x6Aea400CC30498A7C758071029122cC94A448ffd";
  //const amount = "10000000000000000000000";

  //const GLXToken = await hre.ethers.getContractFactory("GLXToken");
  //const glxToken = await GLXToken.attach(tokenAddress);

  //const tx = await glxToken.mint(account, amount);
  //console.log("Token minted:", tx.hash);

  //const GLXShip = await hre.ethers.getContractFactory("GLXShip");
  //const glxShip = await GLXShip.attach("0x5226AE7C172f99B090fA7c6D76B604bf20e0c2c7");
  //const tx = await glxShip.mint("0x6Aea400CC30498A7C758071029122cC94A448ffd");
  //console.log("Ship minted:", tx.hash);
  //const rarity = await glxShip.getRarity(1);
  //const durability = await glxShip.getDurability(1);
  //console.log("Ship rarity:", rarity, ", Ship durability:", durability);

  const GLXEquipment = await hre.ethers.getContractFactory("GLXEquipment");
  const glxEquipment = await GLXEquipment.attach("0x43348F0803D2BC9251F2A338F15A1eF74B0A3D64");
  //const tx = await glxEquipment.mint("0x6Aea400CC30498A7C758071029122cC94A448ffd");
  //console.log("Equipment minted:", tx.hash);
  const rarity = await glxEquipment.getRarity(1);
  const durability = await glxEquipment.getDurability(1);
  console.log("Equipment rarity:", rarity, ", Equipment durability:", durability);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
