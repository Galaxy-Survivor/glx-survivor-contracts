// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./ERC1155.sol";
import "./Context.sol";
import "./AccessControl.sol";

contract GLXFungibleItem is Context, AccessControl, ERC1155 {
	bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

	constructor() ERC1155("") {
		_setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
		_setupRole(MINTER_ROLE, _msgSender());
	}

	modifier onlyAdmin() {
		require(hasRole(DEFAULT_ADMIN_ROLE, _msgSender()), "GLXBox: require admin role");
		_;
	}

	modifier onlyMinter() {
		require(hasRole(MINTER_ROLE, _msgSender()), "GLXBox: require minter role");
		_;
	}

	function addMinter(address minter) public onlyAdmin {
		grantRole(MINTER_ROLE, minter);
	}

	function removeMinter(address minter) public onlyAdmin {
		revokeRole(MINTER_ROLE, minter);
	}

	function mint(address to, uint256 id, uint256 amount) public onlyMinter {
		_mint(to, id, amount, "");
	}

	function burn(address from, uint256 id, uint256 amount) public onlyMinter {
		_burn(from, id, amount);
	}

	function supportsInterface(bytes4 interfaceId)
		public
		view
		override(AccessControl, ERC1155)
		returns (bool)
	{
		return super.supportsInterface(interfaceId);
	}
}
