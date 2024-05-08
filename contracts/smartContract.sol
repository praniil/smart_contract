// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract BalanceCheck {
    uint256 balance;
    address public admin;

    constructor() {
        admin = msg.sender;
        balance = 0;
        updateBalance();
    }

    function updateBalance() internal {
        balance += msg.value;
    }

    function Withdrawl (uint256 amount) public {
        require(msg.sender == admin);
        balance = balance - amount;
    }

    function Deposit (uint256 amount) public returns (uint256) {
        return balance = balance + amount;
    }

    function Balance () public view returns (uint256) {
        return balance;
    }
}