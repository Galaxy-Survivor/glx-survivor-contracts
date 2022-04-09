const hre = require("hardhat");

async function main() {
  const vrfAddress = "0x5935BDb591e53E532744bD74dA9D69FfFf49049d";
  const keyHash = "0xe3062de5196e3ff957141d1c314bcd776d4d08288ddb8370aaeccdc17733b74e";

  const glxItem = "0x7f99494ca31951c523301b1B613ED144852521C4";
  const glxShip = "0x68F92B07A4dA9B789E1079400726487a1130fD51";
  const token = "0x37dC11fA655B837A5411E30795d64805F5EE15D8";

  var startTime = Math.trunc(new Date().getTime() / 1000);
  var endTime = startTime + 864000;

  const GLXEarlySale = await hre.ethers.getContractFactory("GLXEarlySale");
  const glxEarlySale = await GLXEarlySale.deploy(
    glxItem,
    glxShip,
    token,
    vrfAddress,
    keyHash,
    startTime,
    endTime
  );

  await glxEarlySale.deployed();

  console.log("GLXEarlySale deployed to:", glxEarlySale.address);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
