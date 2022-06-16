const hre = require("hardhat");

async function main() {
  const glxTicket = "0xF9A4acF2d891E8C7b6D567F65acdF87CCb23D733";
  const token = "0x37dC11fA655B837A5411E30795d64805F5EE15D8";

  var startTime = Math.trunc(new Date().getTime() / 1000) + 3600;
  var endTime = startTime + 8640000;

  const GLXTicketSale = await hre.ethers.getContractFactory("GLXTicketSale");
  const glxTicketSale = await GLXTicketSale.deploy(
    glxTicket,
    token,
    startTime,
    endTime
  );

  await glxTicketSale.deployed();

  console.log("GLXTicketSale deployed to:", glxTicketSale.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
