// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./ERC20Capped.sol";
import "./AccessControl.sol";

contract GLXToken is ERC20Capped, AccessControl {
	bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

	constructor()
		ERC20("Galaxy Token", "GLX")
		ERC20Capped(1e9 * 1e18)
	{
		_setupRole(MINTER_ROLE, msg.sender);
	}

	modifier onlyMinter(address account) {
		require(hasRole(MINTER_ROLE, account), "GLXToken: require minter role");
		_;
	}

	function mint(address _to, uint256 _amount)
		external
		onlyMinter(msg.sender)
	{
		super._mint(_to, _amount);
	}
}
