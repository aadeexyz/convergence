// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {IERC20} from "forge-std/interfaces/IERC20.sol";
import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {SafeTransferLib} from "solady/utils/SafeTransferLib.sol";
import {FactoryOfMarketFactories} from "src/FactoryOfMarketFactories.sol";
import {MockUSDC} from "src/mocks/MockUSDC.sol";

contract DeployScript is Script {
    using SafeTransferLib for address;

    uint256 deployerPrivateKey;

    address deployer;
    address admin;
    address forwarderAddress;

    uint256 liquidityFee;
    uint256 protocolFee;
    uint8 oracleDecimals;
    uint256 liquidityFeeUnit;
    uint256 protocolFeeUnit;

    MockUSDC mockUSDC;
    FactoryOfMarketFactories factoryOfMarketFactories;

    function run() external {
        deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        admin = vm.envAddress("ADMIN_ADDRESS");
        forwarderAddress = vm.envAddress("FORWARDER_ADDRESS");
        oracleDecimals = uint8(vm.envUint("ORACLE_DECIMALS"));
        liquidityFeeUnit = vm.envUint("LIQUIDITY_FEE_UNIT");
        protocolFeeUnit = vm.envUint("PROTOCOL_FEE_UNIT");

        deployer = vm.addr(deployerPrivateKey);

        vm.startBroadcast(deployerPrivateKey);
        mockUSDC = new MockUSDC();
        vm.stopBroadcast();

        liquidityFee = liquidityFeeUnit * 10 ** mockUSDC.decimals();
        protocolFee = protocolFeeUnit * 10 ** mockUSDC.decimals();

        vm.startBroadcast(deployerPrivateKey);
        factoryOfMarketFactories = new FactoryOfMarketFactories(
            address(mockUSDC), liquidityFee, protocolFee, forwarderAddress, oracleDecimals, admin
        );
        vm.stopBroadcast();

        uint256 creationFee = factoryOfMarketFactories.creationFee();

        vm.startBroadcast(deployerPrivateKey);
        mockUSDC.mint(deployer, creationFee);
        address(mockUSDC).safeApprove(address(factoryOfMarketFactories), creationFee);
        address trumpFactory = factoryOfMarketFactories.createMarketFactory("Donald Trump", "TRUMP");
        vm.stopBroadcast();

        console.log("Deployer:", deployer);
        console.log("Admin:", admin);
        console.log("MockUSDC:", address(mockUSDC));
        console.log("ForwarderAddress:", forwarderAddress);
        console.log("LiquidityFee:", liquidityFee);
        console.log("ProtocolFee:", protocolFee);
        console.log("OracleDecimals:", oracleDecimals);
        console.log("FactoryOfMarketFactories:", address(factoryOfMarketFactories));
        console.log("TrumpMarketFactory:", trumpFactory);
    }
}
