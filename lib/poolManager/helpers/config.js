require('dotenv').config();
const { PRIVATE_KEY, UNISWAP_V3_FACTORY, SFPM, PANOPTIC_FACTORY, RPC_ENDPOINT } = process.env;

function checkEnvVariables() {
  if (!PRIVATE_KEY || !UNISWAP_V3_FACTORY || !SFPM || !PANOPTIC_FACTORY || !RPC_ENDPOINT) {
    throw new Error('Environment variables PRIVATE_KEY, UNISWAP_V3_FACTORY, SFPM, PANOPTIC_FACTORY, and RPC_ENDPOINT must be set in .env');
  }
}

module.exports = { checkEnvVariables, PRIVATE_KEY, UNISWAP_V3_FACTORY, SFPM, PANOPTIC_FACTORY, RPC_ENDPOINT };
