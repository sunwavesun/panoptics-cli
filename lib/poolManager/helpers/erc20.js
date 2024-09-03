const { ethers } = require('ethers');
const { getSigner } = require('./rpc');
const fs = require('fs');
const path = require('path');

const ERC20ABI = JSON.parse(fs.readFileSync(path.join(__dirname, '../../../contracts/abis/ERC20.abi')));

function getTokenContract(tokenAddress) {
  const signer = getSigner();
  return new ethers.Contract(tokenAddress, ERC20ABI, signer);
}

async function approveToken(tokenAddress, spenderAddress, amount=ethers.constants.MaxUint256) {
  console.log(`Approving token: ${tokenAddress}`);

  const tokenContract = getTokenContract(tokenAddress);
  const tx = await tokenContract.approve(spenderAddress, amount);
  console.log(`Approval transaction hash: ${tx.hash}`)
}

module.exports = { approveToken };
