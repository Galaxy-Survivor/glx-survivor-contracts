// SPDX-License-Identifier: UNLICENSE

pragma solidity ^0.8.0;

import "./Context.sol";
import "./Ownable.sol";

interface IERC721Mint {
	function mint(address to, uint256 id) external;
}

interface IERC1155Mint {
	function mint(address to, uint256 id, uint256 amount) external;
	function burn(address from, uint256 id, uint256 amount) external;
}

contract GLXLogic is Context, Ownable {
	uint256 constant MAX_BOX_ID = 1e6;

	uint256 constant GLX_SHIP = 0;
	uint256 constant GLX_EQUIPMENT = 1;
	uint256 constant GLX_ITEM = 2;

	IERC721Mint private _ship;
	IERC721Mint private _equipment;
	IERC1155Mint private _item;

	event CreateBox(address indexed to, uint256 indexed boxId, uint256 amount);
	event Unbox(address indexed operator, uint256 indexed boxId, uint256 amount);

	constructor(address glxShip, address glxEquipment, address glxItem) {
		require(glxShip != address(0x0), "ship contract address is zero");
		require(glxEquipment != address(0x0), "equipment contract address is zero");
		require(glxItem != address(0x0), "item contract address is zero");

		_ship = IERC721Mint(glxShip);
		_equipment = IERC721Mint(glxEquipment);
		_item = IERC1155Mint(glxItem);
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
		_item.mint(to, boxId, amount);
	}

	function unbox(uint256 boxId, uint256 amount)
		external
		onlyValidBoxId(boxId)
	{
		emit Unbox(_msgSender(), boxId, amount);
		_item.burn(_msgSender(), boxId, amount);
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
			_ship.mint(recipient, itemId);
		} else if (itemType == GLX_EQUIPMENT) {
			_equipment.mint(recipient, itemId);
		} else if (itemType == GLX_ITEM) {
			_item.mint(recipient, itemId, amount);
		} else {
			revert("invalid item type");
		}
	}
}
