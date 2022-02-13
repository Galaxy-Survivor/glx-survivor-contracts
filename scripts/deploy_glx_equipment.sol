const hre = require("hardhat");

async function main() {
  const vrfAddress = "0x0BF0dC7e27011bd811fc929Bf8B64423f979fE57";
  const keyHash = "0xe3062de5196e3ff957141d1c314bcd776d4d08288ddb8370aaeccdc17733b74e";

  const GLXEquipment = await hre.ethers.getContractFactory("GLXEquipment");
  const glxEquipment = await GLXEquipment.deploy(
    "https://test.galaxysurvivor.io/api/equipments/",
    vrfAddress,
    keyHash,
  );

  await glxEquipment.deployed();

  console.log("GLXEquipment deployed to:", glxEquipment.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
