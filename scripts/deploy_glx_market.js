const hre = require("hardhat");

async function main() {
  const GLXMarket = await hre.ethers.getContractFactory("GLXMarket");
  const glxMarket = await GLXMarket.deploy(
	  "0xB65D1635f5518533Ab9aB4520462e3BD24DD0474",
	  "0x2d6A90ad19F658F0EF59226BAC0AE067364F9ee6",
	  "0xb794f95e2C6cF686BF6656611B75a92C8D31d07C",
	  "0xe44a252435Ad21b5aECBE5BDEAB648f668095d52"
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
