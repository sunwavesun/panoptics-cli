require('dotenv').config();
const { UNISWAP_V3_FACTORY, SFPM, PANOPTIC_FACTORY, RPC_ENDPOINT } = process.env;

function checkEnvVariables() {
  if (!UNISWAP_V3_FACTORY || !SFPM || !PANOPTIC_FACTORY || !RPC_ENDPOINT) {
    throw new Error('Environment variables UNISWAP_V3_FACTORY, SFPM, PANOPTIC_FACTORY, and RPC_ENDPOINT must be set in .env');
  }
}

module.exports = { checkEnvVariables, UNISWAP_V3_FACTORY, SFPM, PANOPTIC_FACTORY, RPC_ENDPOINT };
