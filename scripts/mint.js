const hre = require("hardhat");

async function main() {
  //const tokenAddress = "0x8bE9209D1f3d51C59D5BB3c5Cc30d8c266652B3f";
  //const account = "0x6Aea400CC30498A7C758071029122cC94A448ffd";
  //const amount = "10000000000000000000000";

  //const GLXToken = await hre.ethers.getContractFactory("GLXToken");
  //const glxToken = await GLXToken.attach(tokenAddress);

  //const tx = await glxToken.mint(account, amount);
  //console.log("Token minted:", tx.hash);

  const GLXShip = await hre.ethers.getContractFactory("GLXShip");
  const glxShip = await GLXShip.attach("0xCe79242434c3101bd8d0ae302134f332dC05a18f");
  const tx = await glxShip.mint("0x6Aea400CC30498A7C758071029122cC94A448ffd", 3);
  console.log("Ship minted:", tx.hash);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
