// SPDX-License-Identifier: UNLICENSE

pragma solidity ^0.8.0;

import "./Context.sol";
import "./Ownable.sol";

interface IERC721Mint {
	function mint(address to, uint256 id) external;
}

interface IGLXFungibleItem {
	function mint(address to, uint256 id, uint256 amount) external;
	function burn(address from, uint256 id, uint256 amount) external;
}

contract GLXLogic is Context, Ownable {
	uint256 constant MAX_BOX_ID = 1e6;

	uint256 constant GLX_SHIP = 0;
	uint256 constant GLX_ITEM = 1;
	uint256 constant GLX_FUNGIBLE_ITEM = 2;

	IERC721Mint private _glxShip;
	IERC721Mint private _glxItem;
	IGLXFungibleItem private _glxFungibleItem;

	event CreateBox(address indexed to, uint256 indexed boxId, uint256 amount);
	event Unbox(address indexed operator, uint256 indexed boxId, uint256 amount);

	constructor(address glxShip, address glxItem, address glxFungibleItem) {
		require(glxShip != address(0x0), "ship contract address is zero");
		require(glxItem != address(0x0), "item contract address is zero");
		require(glxFungibleItem != address(0x0), "fungible item contract address is zero");

		_glxShip = IERC721Mint(glxShip);
		_glxItem = IERC721Mint(glxItem);
		_glxFungibleItem = IGLXFungibleItem(glxFungibleItem);
	}

	modifier onlyValidBoxId(uint256 boxId) {
		require(boxId < MAX_BOX_ID, "invalid box id");
		_;
	}

	function craftBox(address to, uint256 boxId, uint256 amount)
		external
		onlyOwner
		onlyValidBoxId(boxId)
	{
		emit CreateBox(to, boxId, amount);
		_glxFungibleItem.mint(to, boxId, amount);
	}

	function unbox(uint256 boxId, uint256 amount)
		external
		onlyValidBoxId(boxId)
	{
		emit Unbox(_msgSender(), boxId, amount);
		_glxFungibleItem.burn(_msgSender(), boxId, amount);
	}

	function unboxPhase2(
		address recipient,
		uint256[] calldata itemIds,
		uint256[] calldata types,
		uint256[] calldata amounts
	) external {
		require(recipient != address(0x0), "repicient address is zero");
		require(
			itemIds.length == types.length && itemIds.length == amounts.length,
			"invalid input data"
		);
		for (uint256 i = 0; i < itemIds.length; i++) {
			_mintItem(recipient, itemIds[i], types[i], amounts[i]);
		}
	}

	function _mintItem(address recipient, uint256 itemId, uint256 itemType, uint256 amount) private {
		if (itemType == GLX_SHIP) {
			_glxShip.mint(recipient, itemId);
		} else if (itemType == GLX_ITEM) {
			_glxItem.mint(recipient, itemId);
		} else if (itemType == GLX_FUNGIBLE_ITEM) {
			_glxFungibleItem.mint(recipient, itemId, amount);
		} else {
			revert("invalid item type");
		}
	}
}
