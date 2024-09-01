const { Command } = require('commander');
const program = new Command();

const keyManager = require('./lib/keyManager');
const poolManager = require('./lib/poolManager');
const optionManager = require('./lib/optionManager');

program
  .name('pcli')
  .description('CLI to interact with Panoptic Protocol');
  
// KEY MANAGEMENT
const key = program.command('key').description('Manage private keys');

key
  .command('store <key>')
  .description('Store a private key')
  .action(keyManager.storeKey);

key
  .command('view')
  .description('View the stored private key')
  .action(keyManager.viewKey);

key
  .command('delete')
  .description('Delete the stored private key')
  .action(keyManager.deleteKey);

// POOL MANAGEMENT
const pool = program.command('pool').description('Manage Panoptic pools');

pool
  .command('deploy')
  .description('Deploy a new Panoptic pool, and if necessary, a Uniswap V3 pool based on the provided token pair and fee tier.')
  .requiredOption('--tokenA <address>', 'The contract address of the first token in the pair (e.g., WONE).')
  .requiredOption('--tokenB <address>', 'The contract address of the second token in the pair (e.g., USDC).')
  .requiredOption('--fee <fee>', 'The fee tier for the Uniswap V3 pool, specified in hundredths of a basis point (e.g., 500 for 0.05%).')
  .option('--tick <tick>', '(Optional) The initial tick value for setting the price of the pool. If not provided, a default tick is used that corresponds to a 1:1 price ratio.')
  .action(async (options) => {
    const { tokenA, tokenB, fee, tick } = options;
    try {
      await poolManager.deploy(tokenA, tokenB, fee, tick);
    } catch (error) {
      console.error('Error:', error.message);
    }
  });

pool
  .command('get')
  .description('Retrieve the Panoptic pool associated with the provided token pair and fee tier.')
  .requiredOption('--tokenA <address>', 'The contract address of the first token in the pair (e.g., WONE).')
  .requiredOption('--tokenB <address>', 'The contract address of the second token in the pair (e.g., USDC).')
  .requiredOption('--fee <fee>', 'The fee tier for the Uniswap V3 pool, specified in hundredths of a basis point (e.g., 500 for 0.05%).')
  .action(async (options) => {
    const { tokenA, tokenB, fee } = options;
    try {
      await poolManager.get(tokenA, tokenB, fee);
    } catch (error) {
      console.error('Error:', error.message);
    }
  });

// OPTION MANAGEMENT
const option = program.command('option').description('Manage options');


if (!process.argv.slice(2).length) {
  program.outputHelp();
}

program.parse(process.argv);