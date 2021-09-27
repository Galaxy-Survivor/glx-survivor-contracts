// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/Context.sol";
import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";
import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Pausable.sol";

contract AircraftNFT is Context, AccessControlEnumerable, ERC1155Pausable {
	event UpgradeAircraft( address indexed operator, uint256 indexed aircraftId, uint256 indexed componentId);
	event event_MintedAircraft(address indexed operator, uint256 indexed aircraftId);
	bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
	bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

	uint256 public constant LUCKY_BOX_ID = 0;
	uint256 public constant MAX_COMPONENT_ID = 10000;

	uint256 currentId = 10000;

	mapping (uint256 => address) public aircraftIdToOwner;

	constructor(string memory uri) ERC1155(uri) {
		_setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
		_setupRole(MINTER_ROLE, _msgSender());
		_setupRole(PAUSER_ROLE, _msgSender());
	}

	modifier onlyAdmin() {
		require(
			hasRole(DEFAULT_ADMIN_ROLE, _msgSender()),
			"AircraftNFT: must have admin role"
		);
		_;
	}

	modifier onlyMinter() {
		require(
			hasRole(MINTER_ROLE, _msgSender()),
			"AircraftNFT: must have minter role"
		);
		_;
	}

	modifier onlyPauser() {
		require(
			hasRole(PAUSER_ROLE, _msgSender()),
			"AircraftNFT: must have pauser role"
		);
		_;
	}

	function addMinter(address minter) external onlyAdmin {
		grantRole(MINTER_ROLE, minter);
	}

	function mintAircraft(address to, uint256 amount) external onlyMinter {
		uint256 startId = currentId;
		currentId += amount;

		for (uint256 i = 1; i <= amount; i++) {
			uint256 id = startId + i;
			_mint(to, id, 1, "");
			aircraftIdToOwner[id] = to;
			emit event_MintedAircraft(to,id);
		}
	}

	function mintLuckyBox(address to, uint256 amount) external onlyMinter {
		_mint(to, LUCKY_BOX_ID, amount, "");
	}

	function mintComponent(address to, uint256 id) external onlyMinter {
		require(
			balanceOf(to, LUCKY_BOX_ID) > 0,
			"AircraftNFT: must have lucky box to mint component"
		);
		require(
			id < MAX_COMPONENT_ID && id > LUCKY_BOX_ID,
			"AircraftNFT: invalid component id"
		);

		_burn(to, LUCKY_BOX_ID, 1);
		_mint(to, id, 1, "");
	}

	function upgrade(address to, uint256 aircraftId, uint256 componentId) external onlyMinter {
		require(aircraftId > MAX_COMPONENT_ID, "AircraftNFT: invalid aircraft id");
		require(
			componentId > LUCKY_BOX_ID && componentId < MAX_COMPONENT_ID,
			"AircraftNFT: invalid component id"
		);
		require(
			to == aircraftIdToOwner[aircraftId],
			"AircraftNFT: must be owner of aircraft"
		);
		require(
			balanceOf(to, componentId) > 0,
			"AircraftNFT: must have component to upgrade"
		);

		_burn(to, componentId, 1);

		emit UpgradeAircraft(to, aircraftId, componentId);
	}

	function getAircrafts(address account) external view returns (uint256[] memory) {
		uint256 count;
		for (uint256 i = MAX_COMPONENT_ID+1; i < currentId; i++) {
			if (aircraftIdToOwner[i] == account) {
				count++;
			}
		}

		uint256[] memory res = new uint256[](count);
		uint256 j;
		for (uint256 i = MAX_COMPONENT_ID+1; i < currentId; i++) {
			if (aircraftIdToOwner[i] == account) {
				res[j] = i;
				j++;
			}
		}
		return res;
	}

	function pause() public onlyPauser {
		_pause();
	}

	function unpause() public onlyPauser {
		_unpause();
	}

	function supportsInterface(bytes4 interfaceId)
		public
		view
		override(AccessControlEnumerable, ERC1155)
		returns (bool)
	{
		return super.supportsInterface(interfaceId);
	}
}
