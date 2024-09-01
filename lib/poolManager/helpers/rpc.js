require('dotenv').config();
const { ethers } = require('ethers');

const { RPC_ENDPOINT } = process.env;

if (!RPC_ENDPOINT) {
  throw new Error('RPC_ENDPOINT must be set in the .env file');
}

function getProvider() {
  return new ethers.providers.JsonRpcProvider(RPC_ENDPOINT);
}

module.exports = { getProvider };
