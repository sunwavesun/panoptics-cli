const fs = require('fs');
const path = require('path');
const { storeKey, viewKey, deleteKey } = require('../lib/keyManager');

jest.mock('fs');

describe('keyManager', () => {
  const ENV_FILE = path.resolve(__dirname, '../.env');  // absolute path for consistency

  beforeEach(() => {
    jest.clearAllMocks();
    jest.spyOn(console, 'log').mockImplementation(() => {});
    jest.spyOn(console, 'error').mockImplementation(() => {});
  });

  test('should store a valid private key', () => {
    const validKey = '1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef';

    storeKey(validKey);

    expect(fs.writeFileSync).toHaveBeenCalledWith(ENV_FILE, `PRIVATE_KEY=${validKey}`);
    expect(console.log).toHaveBeenCalledWith('Key stored successfully.');
  });

  test('should not store an invalid private key', () => {
    const invalidKey = 'invalid_key';

    storeKey(invalidKey);

    expect(fs.writeFileSync).not.toHaveBeenCalled();
    expect(console.error).toHaveBeenCalledWith('Error: Invalid private key format. It must be a 64-character hexadecimal string.');
  });

  test('should view the stored private key', () => {
    const storedKey = 'PRIVATE_KEY=1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef\n';
  
    fs.readFileSync.mockReturnValue(storedKey);
    fs.existsSync.mockReturnValue(true);
  
    viewKey();
  
    expect(fs.readFileSync).toHaveBeenCalledWith(ENV_FILE, 'utf8');
    expect(console.log).toHaveBeenCalledWith(expect.stringMatching('1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef'));
  });   

  test('should handle view when no key is set', () => {
    fs.existsSync.mockReturnValue(false);

    viewKey();

    expect(console.error).toHaveBeenCalledWith('Error: No key is set.');
  });

  test('should delete the private key', () => {
    const storedContent = 'PRIVATE_KEY=1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef\nOTHER_VAR=somevalue';

    fs.readFileSync.mockReturnValue(storedContent);
    fs.existsSync.mockReturnValue(true);

    deleteKey();

    expect(fs.writeFileSync).toHaveBeenCalledWith(ENV_FILE, 'OTHER_VAR=somevalue');
    expect(console.log).toHaveBeenCalledWith('Key deleted successfully.');
  });

  test('should handle delete when no key is set', () => {
    fs.existsSync.mockReturnValue(false);

    deleteKey();

    expect(console.error).toHaveBeenCalledWith('Error: No key to delete.');
  });
});
