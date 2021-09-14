// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/presets/ERC1155PresetMinterPauser.sol";

contract AircraftNFT is ERC1155PresetMinterPauser {
	uint256 maxComponentId = 1000;
	uint256 currentId = 1000;

	constructor(string memory _uri) ERC1155PresetMinterPauser(_uri) {}

	function addMinter(address _minter) external {
		require(
			hasRole(DEFAULT_ADMIN_ROLE, _msgSender()),
			"AircraftNFT: must have admin role to add minter"
		);

		grantRole(MINTER_ROLE, _minter);
	}

	function mintAircraft(address _to, uint256 _amount) external {
		require(
			hasRole(MINTER_ROLE, _msgSender()),
			"AircraftNFT: must have minter role to mint"
		);

		uint256 startId = currentId;
		currentId += _amount;

		for (uint256 i = 1; i <= _amount; i++) {
			_mint(_to, startId + i, 1, "");
		}
	}

	function mintComponent(address _to, uint256 _id, uint256 _amount) external {
		require(
			hasRole(MINTER_ROLE, _msgSender()),
			"AircraftNFT: must have minter role to mint"
		);

		require(
			_id < maxComponentId,
			"AircraftNFT: id must less than or equal to max component id"
		);

		_mint(_to, _id, _amount, "");
	}
}
