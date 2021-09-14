// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/presets/ERC1155PresetMinterPauser.sol";

contract AircraftNFT is ERC1155PresetMinterPauser {
	uint256 MAX_GENESIS_SUPPLY = 10000;
	uint256 startAircraftId = 1000;
	uint256 aircraftCount = 0;

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

		require(
			aircraftCount + _amount <= MAX_GENESIS_SUPPLY,
			"AircarfNFT: exceed max genesis supply"
		);

		uint256 currentId = startAircraftId + aircraftCount;
		aircraftCount += _amount;

		for (uint256 i = 1; i <= _amount; i++) {
			_mint(_to, currentId + i, 1, "");
		}
	}

	function mintComponent(address _to, uint256 _id, uint256 _amount) external {
		require(
			hasRole(MINTER_ROLE, _msgSender()),
			"AircraftNFT: must have minter role to mint"
		);

		require(
			_id < startAircraftId,
			"AircraftNFT: id must less than start aircraft index"
		);

		_mint(_to, _id, _amount, "");
	}
}
