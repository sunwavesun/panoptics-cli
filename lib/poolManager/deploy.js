const { ethers } = require('ethers');
const { checkEnvVariables } = require('./helpers/config');
const { getUniswapV3Pool, createUniswapV3Pool, initializeUniswapV3Pool } = require('./helpers/uniswap');
const { getPanopticPool, deployNewPool } = require('./helpers/panoptic');
const { calculateSqrtPriceX96, calculateSalt } = require('./helpers/utils');

async function deploy(tokenA, tokenB, fee, tick = null) {
  checkEnvVariables();

  let uniswapV3Pool = await getUniswapV3Pool(tokenA, tokenB, fee);

  // create the Uniswap V3 pool if it does not exist
  if (!uniswapV3Pool || uniswapV3Pool === ethers.constants.AddressZero) {
    console.log('Uniswap V3 pool does not exist, creating one...');
    uniswapV3Pool = await createUniswapV3Pool(tokenA, tokenB, fee);
    console.log('Uniswap V3 pool created at:', uniswapV3Pool);

    // if tick is not provided, use a default sqrtPriceX96 value for initialization
    let sqrtPriceX96;
    if (tick === null) {
      console.log('No tick provided. using default initialization.');
      sqrtPriceX96 = ethers.BigNumber.from('0x1000000000000000000000000'); // corresponds to 1:1 price ratio
    } else {
      console.log(`Using provided tick: ${tick}`);
      sqrtPriceX96 = calculateSqrtPriceX96(tick);
    }
  
    await initializeUniswapV3Pool(uniswapV3Pool, sqrtPriceX96);
  } else {
    console.log('Uniswap V3 pool already exists at:', uniswapV3Pool);
  }

  let panopticPool = await getPanopticPool(uniswapV3Pool);

  // create the Panoptic pool if it does not exist
  if (!panopticPool || panopticPool === ethers.constants.AddressZero) {
    console.log('Panoptic pool does not exist, creating one...');
    const deployer = await ethers.Signer.getAddress();
    const salt = calculateSalt(uniswapV3Pool, deployer, nonce);
    panopticPool = await deployNewPool(tokenA, tokenB, fee, salt);
    console.log('Panoptic pool created at:', panopticPool);
  } else {
    console.log('Panoptic pool already exists at:', panopticPool);
  }

  // return { uniswapV3Pool, panopticPool };
}

module.exports = { deploy };