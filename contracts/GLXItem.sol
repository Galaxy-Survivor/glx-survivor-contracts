// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./ERC1155.sol";
import "./Context.sol";
import "./Strings.sol";
import "./AccessControl.sol";

contract GLXItem is Context, AccessControl, ERC1155 {
	using Strings for uint256;

	bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

	uint256 public constant MAX_BOX_ID = 1000;

	uint256 private _currentItemId = 1000;

	event Unbox(address indexed operator, uint256 boxId, uint256 amount);

	constructor(string memory baseURI) ERC1155(baseURI) {
		_setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
		_setupRole(MINTER_ROLE, _msgSender());
	}

	modifier onlyAdmin() {
		require(hasRole(DEFAULT_ADMIN_ROLE, _msgSender()), "GLXItem: require admin role");
		_;
	}

	modifier onlyMinter() {
		require(hasRole(MINTER_ROLE, _msgSender()), "GLXItem: require minter role");
		_;
	}

	function addMinter(address minter) public onlyAdmin {
		grantRole(MINTER_ROLE, minter);
	}

	function removeMinter(address minter) public onlyAdmin {
		revokeRole(MINTER_ROLE, minter);
	}

	function uri(uint256 tokenId) public view override returns (string memory) {
		string memory baseURI = super.uri(tokenId);
		return string(abi.encodePacked(baseURI, tokenId.toString()));
	}

	function mintBox(address to, uint256 id, uint256 amount) public onlyMinter {
		require(id <= MAX_BOX_ID, "GLXItem: invalid box id");
		_mint(to, id, amount, "");
	}

	function mintItem(address to) public onlyMinter {
		_currentItemId++;
		_mint(to, _currentItemId, 1, "");
	}

	function unbox(uint256 id, uint256 amount) public {
		require(id <= MAX_BOX_ID, "GLXItem: invalid box id");
		_burn(_msgSender(), id, amount);
		emit Unbox(_msgSender(), id, amount);
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
