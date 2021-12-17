const hre = require("hardhat");

async function main() {
  const GLXMarket = await hre.ethers.getContractFactory("GLXMarket");
  const glxMarket = await GLXMarket.deploy(
	  "0xC5ce35F36CcC96a0ca878fDF05553C82a0bC64E2",
	  "0xf49B1b3807876ed7EBFC73164D28eAE1A265b800",
	  "0x4b09e73E2f2aA55a585eafD5D26410E19882e39f",
	  "0x1D32E9F60CF0650Bd89f64BFCcc597fb0eE47107"
  );

  await glxMarket.deployed();

  console.log("GLXMarket deployed to:", glxMarket.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
