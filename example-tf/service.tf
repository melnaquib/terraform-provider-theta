terraform {
  required_providers {
    theta = {
      source = "melnaquib/theta"
    }
  }
}

provider "theta" {
  foo = "example_value"
  #
  # if 'foo' wasn't set here by us, then the value would default to the value 
  # assigned to the environment variable 'THETA_FOO' or the default value of nil
  # if the environment variable wasn't set.

  # mainchainIDStr = "privatenet",
  # mainchainRPC = "http://127.0.0.1:16888/rpc",
  # wTHETAAddr = "0x7d73424a8256C0b2BA245e5d5a3De8820E45F390",
  # registrarOnMainchainAddr = "0x08425D9Df219f93d5763c3e85204cb5B4cE33aAa",
  # govTokenContractAddr = "0x7ad6cea2bc3162e30a3c98d84f821b3233c22647",
  # mainchainTFuelTokenBankAddr = "0xA10A3B175F0f2641Cf41912b887F77D8ef34FAe8",
  # mainchainTNT20TokenBankAddr = "0x6E05f58eEddA592f34DD9105b1827f252c509De0",
  # mainchainTNT721TokenBankAddr = "0x79EaFd0B5eC8D3f945E6BB2817ed90b046c0d0Af",
  # mainchainTNT1155TokenBankAddr = "0x2Ce636d6240f8955d085a896e12429f8B3c7db26",

}

# resource "theta_subnet" "subnet_1" {
#     subchainID = 360777,
#     subchainIDStr = "tsub360777",
#     subchainRPC = "http://127.0.0.1:16900/rpc",
#     subchainTFuelTokenBankAddr = "0x5a443704dd4B594B382c22a083e2BD3090A6feF3",
#     subchainTNT20TokenBankAddr = "0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D",
#     # subchainTNT721TokenBankAddr = "0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA",
#     # subchainTNT1155TokenBankAddr = "0x47c5e40890bcE4a473A49D7501808b9633F29782",
#     # initialFee = 20000,

#     crossChainTransferFeeInTFuel = 10

# }
 
resource "theta_subnet" "subnet_2" {
  id = 360777
  # name = "tsub360777"
  rpc = "http://127.0.0.1:16900/rpc"
  tfuel_token_bank_addr = "0x5a443704dd4B594B382c22a083e2BD3090A6feF3"
  tfuel_tnt20_bank_addr = "0x47e9Fbef8C83A1714F1951F142132E6e90F5fa5D"
  # tfuel_tnt21_bank_addr = "0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA"
  # tfuel_tnt1155_bank_addr = "0x47c5e40890bcE4a473A49D7501808b9633F29782"
  initial_fee = 20000
  cross_chain_transfer_fee_in_tfuel = 10
  key = "~/.thetacli/keys/encrypted/2E833968E5bB786Ae419c4d13189fB081Cc43bab"

  # mainchainID="privatenet"
  # initValidatorSet="./data/init_validator_set.json"
  # admin="0x2E833968E5bB786Ae419c4d13189fB081Cc43bab"
  # fallbackReceiver="0x2E833968E5bB786Ae419c4d13189fB081Cc43bab"
  # genesis="./snapshot"

}

# resource "theta_evm_contract" "test_token" {
#   network = "subnet_1"
#   artifact = file("./internal/provider/testdata/Token.json")
#   signer   = evm_random_pk.deployer.pk
#   constructor_args = [
#     "Test1",
#     "TST",
#     1000000000 * pow(10, 18),
#     18,
#   ]
# }

# resource "theta_edge_vod" "res_store_1" {
#   file                       = "proj/video1.mp4"
#   "playback_policy"= "public",
#   "nft_collection"= "0x5d0004fe2e0ec6d002678c7fa01026cabde9e793",
# }


# resource "theta_edge_store" "res_store_1" {
#   file                       = "proj/file1.data"
# }

# resource "theta_edge_function" "terraform_lambda_func20" {
#   filename                       = "thetafn/thetafn.js"
#   function_name                  = "thetafn33"
#   handler                        = "main"
#   runtime                        = "denotheta.2"
# }

# resource "theta_edge_ai" "ai_1" {
#   name    = "example-${random_pet.this.id}"
#   runtime = "cloudfront-js-1.0"
#   vm = ""
#   model = ""
#   port = ""
#   replicas = ""
#   code    = file("${path.module}/example.py")
# }

