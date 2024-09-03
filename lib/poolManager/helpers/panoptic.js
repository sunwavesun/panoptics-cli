const { ethers } = require('ethers');
const { getProvider, getSigner, getGasLimit } = require('./rpc');
const { PANOPTIC_FACTORY } = require('./config');
const fs = require('fs');
const path = require('path');
const { approveToken } = require('./erc20');

const PanopticFactoryABI = JSON.parse(fs.readFileSync(path.join(__dirname, '../../../contracts/abis/PanopticFactory.abi')));

function getPanopticFactoryContract() {
  const signer = getSigner();
  return new ethers.Contract(PANOPTIC_FACTORY, PanopticFactoryABI, signer);
}

async function getPanopticPool(univ3pool) {
  const panopticFactory = getPanopticFactoryContract();
  return await panopticFactory.getPanopticPool(univ3pool);
}

async function deployNewPool(tokenA, tokenB, fee, salt) {
  const panopticFactory = getPanopticFactoryContract();
  await approveToken(tokenA, PANOPTIC_FACTORY);
  await approveToken(tokenB, PANOPTIC_FACTORY);
  return await panopticFactory.deployNewPool(tokenA, tokenB, fee, salt, {
    gasLimit: getGasLimit()
  });
}

module.exports = { getPanopticPool, deployNewPool };
