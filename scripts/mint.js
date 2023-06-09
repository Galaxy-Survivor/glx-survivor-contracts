const hre = require("hardhat");

async function main() {
  const tokenAddress = "0x37dC11fA655B837A5411E30795d64805F5EE15D8";
  const account = "0x07a35a5ddA7F50f471eDfaB96666a6b70a154403";
  const amount = "1000000000000";

  //const GLXToken = await hre.ethers.getContractFactory("GLXToken");
  const USDCToken = await hre.ethers.getContractFactory("USDC");
  const token = await USDCToken.attach(tokenAddress);

  const tx = await token.mint(account, amount);
  console.log("Token minted:", tx.hash);


  //const GLXShip = await hre.ethers.getContractFactory("GLXShip");
  //const glxShip = await GLXShip.attach("0x5226AE7C172f99B090fA7c6D76B604bf20e0c2c7");
  //const tx = await glxShip.mint("0x6Aea400CC30498A7C758071029122cC94A448ffd");
  //console.log("Ship minted:", tx.hash);
  //const rarity = await glxShip.getRarity(1);
  //const durability = await glxShip.getDurability(1);
  //console.log("Ship rarity:", rarity, ", Ship durability:", durability);

  //const GLXEquipment = await hre.ethers.getContractFactory("GLXEquipment");
  //const glxEquipment = await GLXEquipment.attach("0x43348F0803D2BC9251F2A338F15A1eF74B0A3D64");
  //const tx = await glxEquipment.mint("0x6Aea400CC30498A7C758071029122cC94A448ffd");
  //console.log("Equipment minted:", tx.hash);
  //const rarity = await glxEquipment.getRarity(1);
  //const durability = await glxEquipment.getDurability(1);
  //console.log("Equipment rarity:", rarity, ", Equipment durability:", durability);

  //const GLXLogic = await hre.ethers.getContractFactory("GLXLogic");
  //const glxLogic = await GLXLogic.attach("0xe44a252435Ad21b5aECBE5BDEAB648f668095d52");
  //const tx = await glxLogic.craftBox("0x6Aea400CC30498A7C758071029122cC94A448ffd", 1, 10);
  //const tx = await glxLogic.unbox(1, 1);
  //console.log("Transaction hash:", tx.hash);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
