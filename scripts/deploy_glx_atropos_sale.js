const hre = require("hardhat");

async function main() {
  const vrfAddress = "0x2eD832Ba664535e5886b75D64C46EB9a228C2610";
  const subscriptionID = 131;
  const keyHash = "0x354d2f95da55398f44b7cff77da56283d9c6c829a4bdf1bbcaf2ad6a4d081f61";

  const glxItem = "0x1C0cEc88c8baFDE3e7818c10e553F54844F197C5";
  const glxShip = "0x33adE2bE7D41aF792eCD251C456f62526E718e5a";
  const token = "0x37dC11fA655B837A5411E30795d64805F5EE15D8";
  const ticket = "0xF9A4acF2d891E8C7b6D567F65acdF87CCb23D733";

  var startTime = Math.trunc(new Date().getTime() / 1000);
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
