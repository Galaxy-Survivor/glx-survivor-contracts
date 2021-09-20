let contractData;
let abi;
let contractAddress = '0x3Bce080b858D6d680899e8F2045b56d74F6c7ac5';
let contract;
let provider;
let signer;
fetch("/artifacts/contracts/CryptoTom.sol/CryptoTom.json").then(res => res.json()).then(data => {
    contractData = data;
    abi = contractData.abi;
})


function enable() {
    window.ethereum.enable();

    ethereum.request({
        method: 'eth_requestAccounts'
    });

    provider = new ethers.providers.Web3Provider(window.ethereum)

    // The Metamask plugin also allows signing transactions to
    // send ether and pay to change state within the blockchain.
    // For this, you need the account signer...
    signer = provider.getSigner()


}

// async function isEnable() {
//     if (window.ethereum) {

//         const accounts = web3.eth.accounts;
//         // window.web3 = new Web3(window.ethereum);
//         return true;
//     }
//     return false;
// }

async function loadContract() {
    // contract = new web3.eth.Contract(abi, contractAddress)
    // console.log(abi)
    // console.log(contractAddress);
    contract = new ethers.Contract(contractAddress, abi, signer);
    // contract.on("Transfer", (from, to, amount, event) => {
    //     console.log(`${ from } sent ${ (amount) } to ${ to}`);
    //     // The event object contains the verbatim log data, the
    //     // EventFragment and functions to fetch the block,
    //     // transaction and receipt and event functions
    // });

    // // A filter for when a specific address receives tokens
    // myAddress = "0x8ba1f109551bD432803012645Ac136ddd64DBA72";
    // filter = contract.filters.Transfer(null, myAddress)
    //     // {
    //     //   address: 'dai.tokens.ethers.eth',
    //     //   topics: [
    //     //     '0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef',
    //     //     null,
    //     //     '0x0000000000000000000000008ba1f109551bd432803012645ac136ddd64dba72'
    //     //   ]
    //     // }

    // // Receive an event when that filter occurs
    // contract.on(filter, (from, to, amount, event) => {
    //     // The to will always be "address"
    //     console.log(`I got ${ formatEther(amount) } from ${ from }.`);
    // });
}

async function greet() {
    // await contract.methods.balanceOf(window.ethereum.selectedAddress).call(function(err, res) {
    //     if (err) {
    //         console.log("An error occured", err)
    //         return
    //     }
    //     console.log("RES ", res)
    // })
    // console.log(await contract.methods.approve(window.ethereum.selectedAddress, '1487595000024000000').send({
    //     from: window.ethereum.selectedAddress
    // }))
    // let data = await contract.methods.transferFrom(window.ethereum.selectedAddress, '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266', '1487595000024000000').send({
    //     from: window.ethereum.selectedAddress
    // })
    let balance = ethers.utils.formatUnits(await contract.balanceOf(ethereum.selectedAddress), 18)
    console.log(balance)

    const daiWithSigner = contract.connect(signer);

    // Each DAI has 18 decimal places
    const dai = ethers.utils.parseUnits("1.0", 18);

    // Send 1 DAI to "ricmoo.firefly.eth"
    tx = daiWithSigner.transfer("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", dai);
    // console.log(data)
    // try {
    //     const transactionHash = await ethereum.request({
    //         method: 'eth_sendTransaction',
    //         params: [{
    //             to: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    //             'from': ethereum.selectedAddress,
    //             value: '1487595000024000000',
    //             // And so on...
    //         }, ],
    //     });
    //     // Handle the result
    //     console.log(transactionHash);
    // } catch (error) {
    //     console.error(error);
    // }



}