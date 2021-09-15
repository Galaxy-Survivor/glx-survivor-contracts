// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/presets/ERC1155PresetMinterPauser.sol";

contract AircraftNFT is ERC1155PresetMinterPauser {
	uint256 LUCKY_BOX_ID = 0;
	uint256 MAX_COMPONENT_ID = 10000;

	uint256 currentId = 10000;

	constructor(string memory _uri) ERC1155PresetMinterPauser(_uri) {}

	modifier onlyMinter() {
		require(
			hasRole(MINTER_ROLE, _msgSender()),
			"AircraftNFT: must have minter role to mint"
		);
		_;
	}

	function addMinter(address _minter) external {
		require(
			hasRole(DEFAULT_ADMIN_ROLE, _msgSender()),
			"AircraftNFT: must have admin role to add minter"
		);

		grantRole(MINTER_ROLE, _minter);
	}

	function mintAircraft(address _to, uint256 _amount) external onlyMinter {
		uint256 startId = currentId;
		currentId += _amount;

		for (uint256 i = 1; i <= _amount; i++) {
			_mint(_to, startId + i, 1, "");
		}
	}

	function mintLuckyBox(address _to, uint256 _amount) external onlyMinter {
		_mint(_to, LUCKY_BOX_ID, _amount, "");
	}

	function mintComponent(address _to, uint256 _id, uint256 _amount) external onlyMinter {
		require(
			_id < MAX_COMPONENT_ID && _id > LUCKY_BOX_ID,
			"AircraftNFT: id must less than max component id"
		);

		_mint(_to, _id, _amount, "");
	}

	function mintComponents(
		address _to,
		uint256[] memory _ids,
		uint256[] memory _amounts
	) external {
		bool isValid = true;
		for (uint256 i = 0; i < _ids.length; i++) {
			if (_ids[i] >= MAX_COMPONENT_ID || _ids[i] == LUCKY_BOX_ID) {
				isValid = false;
				break;
			}
		}
		require(isValid, "AircraftNFT: id must less than max component id");

		_mintBatch(_to, _ids, _amounts, "");
	}
}
