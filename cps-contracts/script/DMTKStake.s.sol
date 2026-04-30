// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../src/MTKStake.sol";

contract DMTKStake is Script {
    function run() external {
        IERC20 _mtkToken = IERC20(0x5e6BBd0266c1E69F0F384ed0EEd7BB982C732d1E); // 替换为实际的MTK代币地址

        vm.startBroadcast();
        MTKStake stake = new MTKStake(_mtkToken);
        vm.stopBroadcast();
    }
}
