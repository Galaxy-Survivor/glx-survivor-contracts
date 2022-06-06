const hre = require("hardhat");

async function main() {
  const vrfAddress = "0xE541244c550B28DAa48AcEA38b6942c530fdf1e6";
  const keyHash = "0xe3062de5196e3ff957141d1c314bcd776d4d08288ddb8370aaeccdc17733b74e";

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
