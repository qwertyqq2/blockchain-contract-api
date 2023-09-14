// SPDX-License-Identifier: GPL-3.0

pragma solidity 0.5.16;

contract Coin {
    address public minter;
    mapping(address => uint) public balances;


    constructor() public {
        minter = msg.sender;
    }

    function mint(uint amount) public {
        require(msg.sender == minter);
        balances[minter] += amount;
    }


    function send(address receiver, uint amount) public {
        require(msg.sender == minter);
        if (balances[minter] < amount){
            revert("not enough coins");
        }

        balances[minter] -= amount;
        balances[receiver] += amount;
    }

    function getBalance(address _addr) public view returns (uint) {
        return balances[_addr];
    }
}