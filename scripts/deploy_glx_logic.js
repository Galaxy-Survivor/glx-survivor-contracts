const { ethers, upgrades } = require("hardhat");

async function main() {
  const vrfAddress = "0x2eD832Ba664535e5886b75D64C46EB9a228C2610";
  const subscriptionID = 131;
  const keyHash = "0x354d2f95da55398f44b7cff77da56283d9c6c829a4bdf1bbcaf2ad6a4d081f61";

  const glxItem = "0x1C0cEc88c8baFDE3e7818c10e553F54844F197C5";
  const glxShip = "0x33adE2bE7D41aF792eCD251C456f62526E718e5a";
  const glxEquipment = "0x040Bf32d79FE8A46788Bdc0005832627dc5AE381";

  const GLXLogic = await ethers.getContractFactory("GLXLogic");

  const logic = await upgrades.deployProxy(GLXLogic, [glxShip, glxEquipment, glxItem, vrfAddress, keyHash, subscriptionID], { initializer: '__GLXLogic_init' });

  await logic.deployed();
  console.log("GLXLogic deployed to:", logic.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
