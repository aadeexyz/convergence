// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {MockUSDC} from "src/mocks/MockUSDC.sol";

contract AdminMintScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address mockUsdcAddr = vm.envAddress("MOCK_USDC");
        address recipient = vm.envAddress("MINT_RECIPIENT");
        uint256 amount = vm.envUint("MINT_AMOUNT");

        MockUSDC mockUSDC = MockUSDC(mockUsdcAddr);

        vm.startBroadcast(deployerPrivateKey);
        mockUSDC.mint(recipient, amount);
        vm.stopBroadcast();

        console.log("Minted to:", recipient);
        console.log("Amount:", amount);
    }
}
