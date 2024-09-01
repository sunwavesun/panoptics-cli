const { ethers } = require('ethers');
const { getProvider } = require('./rpc');
const { PANOPTIC_FACTORY } = require('./config');
const fs = require('fs');
const path = require('path');

const PanopticFactoryABI = JSON.parse(fs.readFileSync(path.join(__dirname, '../../../contracts/abis/PanopticFactory.abi')));

function getPanopticFactoryContract() {
  const provider = getProvider();
  return new ethers.Contract(PANOPTIC_FACTORY, PanopticFactoryABI, provider);
}

async function getPanopticPool(univ3pool) {
  const panopticFactory = getPanopticFactoryContract();
  return await panopticFactory.getPanopticPool(univ3pool);
}

async function deployNewPool(tokenA, tokenB, fee, salt) {
  const panopticFactory = getPanopticFactoryContract();
  return await panopticFactory.deployNewPool(tokenA, tokenB, fee, salt);
}

module.exports = { getPanopticPool, deployNewPool };
