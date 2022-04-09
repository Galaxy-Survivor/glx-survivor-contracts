// SPDX-License-Identifier: UNLICENSE

pragma solidity ^0.8.0;

import "./Context.sol";
import "./Ownable.sol";

interface IERC721Mint {
    function mint(address to, uint64 rarity) external;
}

interface IERC1155Mint {
    function mint(address to, uint256 id, uint256 amount) external;
    function burn(address from, uint256 id, uint256 amount) external;
}

contract GLXLogic is Context, Ownable {
    IERC721Mint private _ship;
    IERC721Mint private _equipment;
    IERC1155Mint private _item;

    event BoxCreated(address indexed to, uint256 indexed boxId, uint256 amount);
    event BoxOpened(address indexed operator, uint256 indexed boxId, uint256 amount);

    constructor(address glxShip, address glxEquipment, address glxItem) {
        require(glxShip != address(0x0), "ship contract address is zero");
        require(glxEquipment != address(0x0), "equipment contract address is zero");
        require(glxItem != address(0x0), "item contract address is zero");

        _ship = IERC721Mint(glxShip);
        _equipment = IERC721Mint(glxEquipment);
        _item = IERC1155Mint(glxItem);
    }

    function craftBox(address to, uint256 boxId, uint256 amount)
        external
        onlyOwner
    {
        emit BoxCreated(to, boxId, amount);
        _item.mint(to, boxId, amount);
    }

    function unbox(uint256 boxId, uint256 amount)
        external
    {
        emit BoxOpened(_msgSender(), boxId, amount);
        _item.burn(_msgSender(), boxId, amount);
    }
}
