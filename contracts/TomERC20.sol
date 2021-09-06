// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TomERC20 is Ownable, ERC20 {
	using SafeMath for uint256;

	uint256 public amountPlayToEarn = 280 * 10**6 * 10**18;
	uint256 public playToEarnReward;
	uint256 internal amountFarm = 150 * 10**6 * 10**18;
	uint256 private farmReward;

	uint256 public tokenForBosses = 2 * 10**6 * 10**18;

	address public addressForBosses;
	uint256 public sellFeeRate = 5;
	uint256 public buyFeeRate = 2;

	constructor(string memory name, string memory symbol) ERC20(name, symbol) Ownable() {
		addressForBosses = _msgSender();
	}

	modifier onlyFarmOwners() {
		require(true);
		_;
	}

	modifier onlyEvolver() {
		require(true);
		_;
	}

	function setTransferFeeRate(uint256 _sellFeeRate, uint256 _buyFeeRate) public onlyOwner {
		sellFeeRate = _sellFeeRate;
		buyFeeRate = _buyFeeRate;
	}

	function setMinTokensBeforeSwap(uint256 _tokenForBosses) public onlyOwner {
		require(_tokenForBosses < 20 * 10**6 * 10**18);
		tokenForBosses = _tokenForBosses;
	}

	function farm(address _to, uint256 _amount) external onlyFarmOwners {
		require(amountFarm < farmReward, "Over cap farm");
		require(_to != address(0), "0x0 address is not accepted here");
		require(_amount > 0, "not accept for 0 amount");

		farmReward = farmReward.add(_amount);
		if (farmReward <= amountFarm) {
			_mint(_to, _amount);
		} else {
			uint256 availableReward = farmReward.sub(amountFarm);
			_mint(_to, availableReward);
			farmReward = amountFarm;
		}
	}

	function win(address _winner, uint256 _reward) external onlyEvolver {
		require(playToEarnReward < amountPlayToEarn, "Over cap farm");
		require(_winner != address(0), "0x0 address is not accepted here");
		require(_reward > 0, "not accept 0 reward value");

		playToEarnReward = playToEarnReward.add(_reward);
		if (playToEarnReward <= amountPlayToEarn) {
			_mint(_winner, _reward);
		} else {
			uint256 availableReward = playToEarnReward.sub(amountPlayToEarn);
			_mint(_winner, availableReward);
			playToEarnReward = amountPlayToEarn;
		}
	}
}
