const { ethers } = require("ethers");
const { checkEnvVariables } = require("./helpers/config");
const { getUniswapV3Pool } = require("./helpers/uniswap");
const { getPanopticPool } = require("./helpers/panoptic");

async function get(tokenA, tokenB, fee) {
  checkEnvVariables();

  let uniswapV3Pool = await getUniswapV3Pool(tokenA, tokenB, fee);

  if (!uniswapV3Pool || uniswapV3Pool === ethers.constants.AddressZero) {
    throw new Error('Uniswap V3 pool does not exist for the provided options.');
  } else {
    console.log('Uniswap V3 pool:', uniswapV3Pool);
  }

  let panopticPool = await getPanopticPool(uniswapV3Pool);
  if (!panopticPool || panopticPool === ethers.constants.AddressZero) {
    throw new Error('Panoptic pool does not exist for the provided options.');
  }

  console.log('Panoptic pool:', panopticPool);
  // return { uniswapV3Pool, panopticPool };
}

module.exports = { get };