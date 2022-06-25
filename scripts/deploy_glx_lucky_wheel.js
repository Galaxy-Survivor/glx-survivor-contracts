const hre = require("hardhat");

async function main() {
  const token = "0x37dC11fA655B837A5411E30795d64805F5EE15D8";
  const glxShip = "0x33adE2bE7D41aF792eCD251C456f62526E718e5a";
  const glxItem = "0x1C0cEc88c8baFDE3e7818c10e553F54844F197C5";

  const GLXLuckyWheel = await hre.ethers.getContractFactory("GLXLuckyWheel");
  const glxLuckyWheel = await GLXLuckyWheel.deploy(token, glxShip, glxItem);

  await glxLuckyWheel.deployed();

  console.log("GLXLuckyWheel deployed to:", glxLuckyWheel.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
