// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {FactoryOfMarketFactories} from "src/FactoryOfMarketFactories.sol";
import {MockUSDC} from "src/mocks/MockUSDC.sol";

contract CreateMarketScript is Script {
    using SafeTransferLib for address;

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        address fomf = vm.envAddress("FACTORY_OF_MARKET_FACTORIES");
        string memory name = vm.envString("MARKET_NAME");
        string memory symbol = vm.envString("MARKET_SYMBOL");

        FactoryOfMarketFactories factoryOfMarketFactories = FactoryOfMarketFactories(fomf);
        MockUSDC mockUSDC = MockUSDC(factoryOfMarketFactories.collateralToken());
        uint256 creationFee = factoryOfMarketFactories.creationFee();

        vm.startBroadcast(deployerPrivateKey);
        mockUSDC.mint(deployer, creationFee);
        address(mockUSDC).safeApprove(fomf, creationFee);
        address marketFactory = factoryOfMarketFactories.createMarketFactory(name, symbol);
        vm.stopBroadcast();

        console.log("MarketFactory:", marketFactory);
        console.log("Name:", name);
        console.log("Symbol:", symbol);
    }
}
