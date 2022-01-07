
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity >=0.6.0 <=0.8.9;

contract MyWallet {
    address public owner; // contract owner
    uint256 internal balance; // contract balance in wei
    event Deposited(address _from, uint256 _weiAmount);
    event Withdrawn(address _to, uint256 _weiAmount);
    event Transferred(address _from, address _to, uint256 _weiAmount);

    modifier onlyOwner() {
        require(msg.sender==owner,"sender is not the owner");
        _;
    }

    constructor() {
        owner=msg.sender;
    }

    /**
     * @notice Deposits the sent amount in the contract as credit to be withdrawn.
     * @dev msg.value contains the amount in Wei to be deposited.
     */
    function deposit() public payable {
        balance=balance+msg.value;
        emit Deposited(msg.sender, msg.value);
    }

    /**
     * @notice Withdraws to the owner address the requested amount from the contract.
     * @param weiAmount The amount in Wei to be withdrawn.
     */
    function withdraw(uint256 weiAmount) public onlyOwner {
        require(balance>=weiAmount, "insufficient funds");
        payable(msg.sender).transfer(weiAmount);
        balance -= weiAmount;
        emit Withdrawn(msg.sender,weiAmount);
    }

    /**
     * @notice Transfers the requested amount to another address.
     * @param recipient The address of the recipient.
     * @param weiAmount The amount in Wei to be transferred.
     */
    function transfer(address payable recipient, uint256 weiAmount) public onlyOwner {
        require(getBalance() >= weiAmount, "insufficient funds");
        recipient.transfer(weiAmount);
        balance -= weiAmount;
        emit Transferred(msg.sender, recipient, weiAmount);
    }

    /**
     * @notice Returns the current balance.
     * @return The current contract balance.
     */
    function getBalance() public view returns (uint256) {
        return balance;
    }
}
