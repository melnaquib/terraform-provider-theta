// Hello world example
// curl localhost:4220/services/1/blake3/$(lightning-node dev store services/js-poc/examples/example.js | awk '{print $1}')
// curl localhost:4220/services/1/blake3/$(./target/debug/lightning-node dev store services/js-poc/examples/example.js | awk '{print $1}')

export const main = () => {
    return 'Hello Terraform from Theta Network!';
  };
  