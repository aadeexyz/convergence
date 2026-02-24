.PHONY: deploy deploy-broadcast

deploy:
	. .env && forge script script/Deploy.s.sol:DeployScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v

deploy-broadcast:
	. .env && forge script script/Deploy.s.sol:DeployScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v --broadcast
