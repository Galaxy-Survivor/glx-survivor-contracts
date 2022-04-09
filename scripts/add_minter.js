const hre = require("hardhat");

async function main() {
  //const glxLogic = "0xe44a252435Ad21b5aECBE5BDEAB648f668095d52";
  const glxEarlySale = "0xe63d9E6869A28F4838832803f944C9723a61E618";

  const GLXShip = await hre.ethers.getContractFactory("GLXShip");
  const glxShip = await GLXShip.attach("0x68F92B07A4dA9B789E1079400726487a1130fD51");
  const tx1 = await glxShip.addMinter(glxEarlySale);
  console.log("Add minter for GLXShip:", tx1.hash);

  //const GLXEquipment = await hre.ethers.getContractFactory("GLXEquipment");
  //const glxEquipment = await GLXEquipment.attach("0x228Fdfc58149DC2E8B9e7A72e4d21bf51F0a034D");
  //const tx2 = await glxEquipment.addMinter(glxEarlySale);
  //console.log("Add minter for GLXEquipment:", tx2.hash);

  //const GLXItem = await hre.ethers.getContractFactory("GLXItem");
  //const glxItem = await GLXItem.attach("0x7f99494ca31951c523301b1B613ED144852521C4");
  //const tx3 = await glxItem.addMinter(glxEarlySale);
  //console.log("Add minter for GLXItem:", tx3.hash);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
