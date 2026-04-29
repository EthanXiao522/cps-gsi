// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../src/MyToken.sol";

contract MyTokenTest is Test {
    MyToken token;
    address alice = address(0xABCD);
    address bob = address(0xBEEF);

    function setUp() public {
        token = new MyToken("ABC");
    }

    function testMetadata() public {
        assertEq(token.name(), "ABC");
        assertEq(token.symbol(), "MTK");
        assertEq(token.decimals(), 18);
    }

    function testInitialSupplyAssigned() public {
        uint256 expected = 120 * 10 ** uint256(token.decimals());
        assertEq(token.totalSupply(), expected);
        assertEq(token.balanceOf(address(this)), expected);
    }

    function testTransfer() public {
        uint256 amount = 1 ether;
        token.transfer(alice, amount);
        assertEq(token.balanceOf(alice), amount);
        assertEq(token.balanceOf(address(this)), token.totalSupply() - amount);
    }

    function testApproveAndTransferFrom() public {
        uint256 amount = 100 * 10 ** uint256(token.decimals());
        token.approve(bob, amount);
        assertEq(token.allowance(address(this), bob), amount);

        vm.prank(bob);
        token.transferFrom(address(this), alice, amount);

        assertEq(token.balanceOf(alice), amount);
        assertEq(token.allowance(address(this), bob), 0);
    }
}
