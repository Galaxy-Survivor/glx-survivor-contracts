const hre = require("hardhat");

async function main() {
  const vrfAddress = "0x2eD832Ba664535e5886b75D64C46EB9a228C2610";
  const subscriptionID = 131;
  const keyHash = "0x354d2f95da55398f44b7cff77da56283d9c6c829a4bdf1bbcaf2ad6a4d081f61";

  const glxItem = "0x7f99494ca31951c523301b1B613ED144852521C4";
  const glxShip = "0xD3faF99b7C4D03459D9b08dFaae2B15666ccFCee";
  const token = "0x37dC11fA655B837A5411E30795d64805F5EE15D8";
  const ticket = "0x03a827263b2C65317091e08DB90892e9722eB949";

  var startTime = Math.trunc(new Date().getTime() / 1000);
  var endWhitelistTime = startTime + 86400 * 2;
  var endTime = startTime + 8640000;

  const GLXAtroposSale = await hre.ethers.getContractFactory("GLXAtroposSale");
  const glxAtroposSale = await GLXAtroposSale.deploy(
    glxItem,
    glxShip,
    token,
    ticket,
    vrfAddress,
    keyHash,
    subscriptionID,
    startTime,
    endWhitelistTime,
    endTime
  );

  await glxAtroposSale.deployed();

  console.log("GlxAtroposSale deployed to:", glxAtroposSale.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
