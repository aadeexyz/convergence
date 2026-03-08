.PHONY: deploy deploy-broadcast verify create-market create-market-broadcast

deploy:
	. .env && forge script script/Deploy.s.sol:DeployScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v

deploy-broadcast:
	. .env && forge script script/Deploy.s.sol:DeployScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v --broadcast

verify:
	. .env && forge script script/Deploy.s.sol:DeployScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v --broadcast --resume --verify --verifier custom

create-market:
	. .env && forge script script/CreateMarket.s.sol:CreateMarketScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v

create-market-broadcast:
	. .env && forge script script/CreateMarket.s.sol:CreateMarketScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v --broadcast

admin-mint:
	. .env && forge script script/AdminMint.s.sol:AdminMintScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v

admin-mint-broadcast:
	. .env && forge script script/AdminMint.s.sol:AdminMintScript --rpc-url $$RPC_URL --private-keys $$DEPLOYER_PRIVATE_KEY -v --broadcast
