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

function calculateSalt(univ3pool, address, nonce = 0) {
  // need to shorten the address to 64
  const addressUint64 = ethers.BigNumber.from(address).mask(64);
  const fullHash = ethers.utils.solidityKeccak256(
    ['address', 'uint64', 'uint96'],
    [univ3pool, addressUint64, nonce]
  );

  const shortHash = ethers.BigNumber.from(fullHash).mod(ethers.constants.MaxUint256).mask(96);
  return shortHash;
}

function generateNonce() {
  return ethers.BigNumber.from(ethers.utils.randomBytes(4)).toNumber(); // 4 bytes
}

module.exports = { calculateSqrtPriceX96, calculateSalt, generateNonce };