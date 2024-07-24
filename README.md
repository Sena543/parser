# JSON Parser

This command-line utility provides functionality to tell the validity of a json file. It supports both file input and input via standard input (stdin).

## Usage

To use the utility, execute the compiled binary with the following command-line arguments:

./parser filename.json

## Example Usage

./parser filename.json

### Notes

The utility handles both direct file input (./parser filename) and piped input (cat filename | parser ).
If no filename is provided and stdin is empty or not piped, the utility attempts to use the last argument as a filename.

Error handling for file reading and command-line argument parsing is minimal and may need extension depending on deployment needs.

### TODO
- [ ] Work on test suite.
- [x] Handling unicode characters.
- [ ] Escaping  characters.
