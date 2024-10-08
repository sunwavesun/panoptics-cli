const { ethers } = require('ethers');
const { getProvider, getSigner, getGasLimit } = require('./rpc');
const { UNISWAP_V3_FACTORY } = require('./config');
const fs = require('fs');
const path = require('path');

const IUniswapV3FactoryABI = JSON.parse(fs.readFileSync(path.join(__dirname, '../../../contracts/abis/IUniswapV3Factory.abi')));
const IUniswapV3PoolABI = JSON.parse(fs.readFileSync(path.join(__dirname, '../../../contracts/abis/IUniswapV3Pool.abi')));

function getUniswapFactoryContract() {
  const signer = getProvider();
  return new ethers.Contract(UNISWAP_V3_FACTORY, IUniswapV3FactoryABI, signer);
}

async function getUniswapV3Pool(tokenA, tokenB, fee) {
  const uniswapFactory = getUniswapFactoryContract();
  return await uniswapFactory.getPool(tokenA, tokenB, fee);
}

async function createUniswapV3Pool(tokenA, tokenB, fee) {
  if (!ethers.utils.isAddress(tokenA) || !ethers.utils.isAddress(tokenB)) {
    throw new Error("Invalid token addresses provided");
  }

  const uniswapFactory = getUniswapFactoryContract();
  return await uniswapFactory.createPool(tokenA, tokenB, fee);
}

function getUniswapV3PoolContract(poolAddress) {
  const signer = getSigner();
  return new ethers.Contract(poolAddress, IUniswapV3PoolABI, signer);
}

// initialize the Uniswap V3 pool with the desired tick
async function initializeUniswapV3Pool(poolAddress, sqrtPriceX96) {
  const uniswapV3Pool = getUniswapV3PoolContract(poolAddress);

  const tx = await uniswapV3Pool.initialize(sqrtPriceX96, {
    gasLimit: getGasLimit()
  });
  await tx.wait();

  console.log(`Uniswap V3 pool at ${poolAddress} initialized with sqrtPriceX96: ${sqrtPriceX96.toString()}`);
}

module.exports = { getUniswapV3Pool, createUniswapV3Pool, initializeUniswapV3Pool };
