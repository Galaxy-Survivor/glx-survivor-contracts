const hre = require("hardhat");

async function main() {
  const VRFCoordinator = await hre.ethers.getContractFactory("VRFCoordinator");
  const vrfCoordinator = await VRFCoordinator.deploy();

  await vrfCoordinator.deployed();

  console.log("VRFCoordinator deployed to:", vrfCoordinator.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
