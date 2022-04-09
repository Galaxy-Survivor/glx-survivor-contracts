const hre = require("hardhat");

function toHexString(byteArray) {
  return Array.from(byteArray, function(byte) {
    return ('0' + (byte & 0xFF).toString(16)).slice(-2);
  }).join('')
}

async function main() {
  const oracle = "0x6Aea400CC30498A7C758071029122cC94A448ffd";
  const vrfAddress = "0x5935BDb591e53E532744bD74dA9D69FfFf49049d";

  const VRFCoordinator = await hre.ethers.getContractFactory("VRFCoordinator");
  const vrfCoordinator = await VRFCoordinator.attach(vrfAddress);

  //const tx = await vrfCoordinator.registerProvingKey(
  //  oracle,
  //  [
  //    "0xb25506e03d7509363d331aaf48cae6e28fcb7c3476fa14def7153d985ca9d658",
  //    "0xfd9a4ed1352bcf58813036510c7dec3db1c57212cae40eba627e23ccc5c1f7b8",
  //  ],
  //);
  //console.log("Successfully register proving key:", tx.hash);

  //const keyHash = await vrfCoordinator.hashOfKey([
  //  "0xb25506e03d7509363d331aaf48cae6e28fcb7c3476fa14def7153d985ca9d658",
  //  "0xfd9a4ed1352bcf58813036510c7dec3db1c57212cae40eba627e23ccc5c1f7b8",
  //])
  //console.log("Public key hash:", keyHash);

  const tx = await vrfCoordinator.addConsumer("0xe63d9E6869A28F4838832803f944C9723a61E618");
  console.log("Successfully added consumer:", tx.hash)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
