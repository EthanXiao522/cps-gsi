// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/MyToken.sol";

contract DMyToken is Script {
    function run() external {
        string memory tokenName = "MyToken";

        vm.startBroadcast();
        MyToken token = new MyToken(tokenName);
        vm.stopBroadcast();
    }
}
