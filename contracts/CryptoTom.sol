// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "./TomERC20.sol";

contract CryptoTom is TomERC20, ReentrancyGuard {
	using SafeMath for uint256;

	uint256 public maxSupply = 1000 * 10**6 * 10**18;

	constructor(string memory _name, string memory _symbol) TomERC20(_name, _symbol) {
		_mint(_msgSender(), maxSupply.sub(amountFarm).sub(amountPlayToEarn));
	}

	function burn(uint256 _amount) public {
		_burn(_msgSender(), _amount);
	}

	function setAddressForBosses(address _addressForBosses) external onlyOwner {
		require(_addressForBosses != address(0), "0x0 address is not accepted here");

		addressForBosses = _addressForBosses;
	}
}
