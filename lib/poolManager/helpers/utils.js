const { BigNumber, ethers } = require('ethers');

// refer to the following link for explanation on tick and sqrtPriceX96
// https://docs.uniswap.org/contracts/v4/concepts/managing-positions#tick
function calculateSqrtPriceX96(tick) {
  const sqrt1_0001 = Math.sqrt(1.0001);
  
  // calculate the price ratio using the tick value
  const ratio = sqrt1_0001 ** tick;

  // calculate sqrtPriceX96: ratio * 2^96
  const sqrtPriceX96 = ratio * 2 ** 96;

  // convert to BigNumber and return
  return BigNumber.from(Math.floor(sqrtPriceX96).toString());
}

function calculateSalt(univ3pool, deployer, nonce = 0) {
  return ethers.utils.solidityKeccak256(
    ['address', 'address', 'uint96'],
    [univ3pool, deployer, nonce]
  );
}

module.exports = { calculateSqrtPriceX96, calculateSalt };