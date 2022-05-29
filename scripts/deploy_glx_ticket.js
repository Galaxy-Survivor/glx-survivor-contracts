const hre = require("hardhat");

async function main() {
  const GLXTicket = await hre.ethers.getContractFactory("GLXTicket");
  const glxTicket = await GLXTicket.deploy("https://test.galaxysurvivor.io/api/tickets/");

  await glxTicket.deployed();

  console.log("GLXTicket deployed to:", glxTicket.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
