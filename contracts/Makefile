.PHONY: deploy deploy-broadcast verify

deploy:
	. .env && forge script script/Deploy.s.sol:DeployScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v

deploy-broadcast:
	. .env && forge script script/Deploy.s.sol:DeployScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v --broadcast

verify:
	. .env && forge script script/Deploy.s.sol:DeployScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v --broadcast --resume --verify --verifier etherscan --etherscan-api-key $$BASESCAN_API_KEY
