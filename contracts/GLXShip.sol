// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./ERC721.sol";
import "./ERC721Enumerable.sol";
import "./Counters.sol";
import "./Context.sol";
import "./AccessControl.sol";

contract GLXShip is Context, AccessControl, ERC721Enumerable {
	using Counters for Counters.Counter;

	bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

	Counters.Counter private _tokenIdTracker;

	string private _baseTokenURI;

	constructor(string memory baseURI) ERC721("Galaxy Ship", "GLXShip") {
		_baseTokenURI = baseURI;

		_setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
		_setupRole(MINTER_ROLE, _msgSender());
	}

	modifier onlyAdmin() {
		require(hasRole(DEFAULT_ADMIN_ROLE, _msgSender()), "GLXShip: require admin role");
		_;
	}

	modifier onlyMinter() {
		require(hasRole(MINTER_ROLE, _msgSender()), "GLXShip: require minter role");
		_;
	}

	function _baseURI() internal view override returns (string memory) {
		return _baseTokenURI;
	}

	function addMinter(address minter) external onlyAdmin {
		grantRole(MINTER_ROLE, minter);
	}

	function removeMinter(address minter) external onlyAdmin {
		revokeRole(MINTER_ROLE, minter);
	}

	function mint(address to) public onlyMinter {
		_tokenIdTracker.increment();
		_safeMint(to, _tokenIdTracker.current());
	}

	function supportsInterface(bytes4 interfaceId)
		public
		view
		override(AccessControl, ERC721Enumerable)
		returns (bool)
	{
		return super.supportsInterface(interfaceId);
	}
}
