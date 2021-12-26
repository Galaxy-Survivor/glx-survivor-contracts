const hre = require("hardhat");

async function main() {
  const GLXLogic = await hre.ethers.getContractFactory("GLXLogic");
  const glxLogic = await GLXLogic.deploy(
	  "0x2d6A90ad19F658F0EF59226BAC0AE067364F9ee6",
	  "0xb794f95e2C6cF686BF6656611B75a92C8D31d07C",
	  "0xe44a252435Ad21b5aECBE5BDEAB648f668095d52"
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
