// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./ERC721.sol";
import "./ERC721Enumerable.sol";
import "./Context.sol";
import "./AccessControl.sol";

contract GLXShip is Context, AccessControl, ERC721Enumerable {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    uint256 public constant DEFAULT_DURABILITY = 3;

    struct Ship {
        uint32 empire;
        uint32 rarity;
        uint32 durability;
        uint32 unlockTime;
    }

    string private _baseTokenURI;
    uint256 private _currentID;
    uint256 private immutable _startID;
    uint256 private immutable _endID;
    mapping(uint256 => Ship) internal ships;

    event ShipCreated(address indexed owner, uint256 indexed shipID, uint32 empire, uint32 rarity, uint32 durability);
    event ShipRepaired(uint256 indexed shipID, uint32 durability);

    constructor(string memory baseURI, uint256 startID) ERC721("Galaxy Ship", "GLXShip") {
        _baseTokenURI = baseURI;
        _startID = startID;
        _endID = _startID + 1000000;
        _currentID = _startID;

        _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _grantRole(MINTER_ROLE, _msgSender());
    }

    function _baseURI() internal view override returns (string memory) {
        return _baseTokenURI;
    }

    function addMinter(address minter) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _grantRole(MINTER_ROLE, minter);
    }

    function removeMinter(address minter) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _revokeRole(MINTER_ROLE, minter);
    }

    function mint(address to, uint32 empire, uint32 rarity, uint32 unlockTime) external onlyRole(MINTER_ROLE) {
        _currentID++;
        require(_currentID <= _endID, "limit exceed");

        Ship storage ship = ships[_currentID];
        ship.empire = empire;
        ship.rarity = rarity;
        ship.durability = uint32(DEFAULT_DURABILITY);
        ship.unlockTime = unlockTime;

        emit ShipCreated(to, _currentID, ship.empire, ship.rarity, ship.durability);

        _safeMint(to, _currentID);
    }

    function mint(address to, uint256 id, uint32 empire, uint32 rarity, uint32 durability, uint32 unlockTime) external onlyRole(MINTER_ROLE) {
        require(ownerOf(id) == address(0x0), "ship already exists");

        Ship storage ship = ships[id];
        ship.empire = empire;
        ship.rarity = rarity;
        ship.durability = durability;
        ship.unlockTime = unlockTime;

        _safeMint(to, _currentID);
    }

    function repair(uint256 shipID) external {
        address owner = ownerOf(shipID);
        require(_msgSender() == owner, "only ship's owner");

        Ship storage ship = ships[shipID];
        require(ship.durability > 0, "ship run out of durability");
        ship.durability--;
        emit ShipRepaired(shipID, ship.durability);
    }

    function getRarity(uint256 shipID) external view returns (uint32) {
        return ships[shipID].rarity;
    }

    function getDurability(uint256 shipID) external view returns (uint32) {
        return ships[shipID].durability;
    }

    function getUnlockTime(uint256 shipID) external view returns (uint32) {
        return ships[shipID].unlockTime;
    }

    function _beforeTokenTransfer(
        address,
        address,
        uint256 tokenId
    ) internal override {
        require(ships[tokenId].unlockTime <= _getBlockTimestamp(), "GLXShip: lock transfer");
    }

    function _getBlockTimestamp() private view returns (uint32) {
        return uint32(block.timestamp);
    }

    function supportsInterface(bytes4 interfaceID)
        public
        view
        override(AccessControl, ERC721Enumerable)
        returns (bool)
    {
        return super.supportsInterface(interfaceID);
    }
}
