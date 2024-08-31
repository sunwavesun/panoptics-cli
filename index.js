const { Command } = require('commander');
const program = new Command();

const keyManager = require('./lib/keyManager');
const poolManager = require('./lib/poolManager');
const optionManager = require('./lib/optionManager');

program
  .name('pcli')
  .description('CLI to interact with Panoptic Protocol');
  
// key management
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

// pool management
const pool = program.command('pool').description('Manage Uniswap v3 pools');

// option management
const option = program.command('option').description('Manage options');


if (!process.argv.slice(2).length) {
  program.outputHelp();
}

program.parse(process.argv);