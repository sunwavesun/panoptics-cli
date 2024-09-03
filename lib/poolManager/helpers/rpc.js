require('dotenv').config();
const { ethers } = require('ethers');
const { PRIVATE_KEY } = require('./config');

const { RPC_ENDPOINT } = process.env;

if (!RPC_ENDPOINT) {
  throw new Error('RPC_ENDPOINT must be set in the .env file');
}

const provider = new ethers.providers.JsonRpcProvider(RPC_ENDPOINT);
const signer = new ethers.Wallet(PRIVATE_KEY, provider);

// function getProvider() {
//   return new ethers.providers.JsonRpcProvider(RPC_ENDPOINT);
// }

function getProvider() {
  return provider;
}

function getSigner() {
  return signer;
}

function getGasLimit() {
  return 500000;
}

module.exports = { getProvider, getSigner, getGasLimit };
