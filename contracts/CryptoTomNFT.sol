// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";

contract CryptoTomNFT is ERC721Enumerable, Ownable {
	using SafeMath for uint256;
	using EnumerableSet for EnumerableSet.UintSet;
	using Counters for Counters.Counter;

	enum Tribe {
		SKYLER,
		HYDREIN,
		PLASMER,
		STONIC
	}

	event LayEgg(uint256 indexed _tokenId, address _buyer);
	event Evolve(uint256 indexed _tokenId, uint256 _dna);
	event Buy(
		uint256 indexed _tokenId,
		address _buyer,
		address _seller,
		uint256 _price
	);
	event UpdateTribe(uint256 indexed _tokenId, Tribe _tribe);
	event UpgradeGeneration(uint256 indexed _tokenId, uint256 _newGeneration);
	event Exp(uint256 indexed _tokenId, uint256 _exp);
	event Working(uint256 indexed _tokenId, uint256 _time);
	event PlaceOrder(uint256 indexed _tokenId, address _seller, uint256 _price);
	event CancelOrder(uint256 indexed _tokenId, address _seller);
	event UpdatePrice(
		uint256 indexed _tokenId,
		address _seller,
		uint256 _newPrice
	);
	event FillOrder(uint256 indexed _tokenId, address _seller);

	struct CryptoTom {
		uint256 generation;
		Tribe tribe;
		uint256 exp;
		uint256 dna;
		uint256 farmTime;
		uint256 bornTime;
	}

	struct ItemSale {
		uint256 tokenId;
		address owner;
		uint256 price;
	}

	Counters.Counter private _tokenIds;
	mapping(uint256 => bool) public isEvolved;

	mapping(uint256 => CryptoTom) internal tommers;
	mapping(uint256 => ItemSale) internal markets;

	EnumerableSet.UintSet private tokenSales;
	mapping(address => EnumerableSet.UintSet) private sellerTokens;

	IERC20 public tomERC20;

	uint256 public priceEgg = 12000 * 10**18;
	uint256 public feeEvolve = 24000 * 10**18;
	uint256 public feeChangeTribe = 10000 * 10**18;
	uint256 public feeUpgradeGeneration = 30000 * 10**18;
	uint256 public feeMarketRate = 10;
	uint256 public divPercent = 1000;

	constructor(
		string memory _name,
		string memory _symbol,
		address _tomERC20
	) ERC721(_name, _symbol) {
		tomERC20 = IERC20(_tomERC20);
	}

	modifier onlyFarmer() {
		require(true);
		_;
	}

	modifier onlySpawner() {
		require(true);
		_;
	}

	modifier onlyBattlefield() {
		require(true);
		_;
	}

	function working(uint256 _tokenId, uint256 _time) public onlyFarmer {
		require(_time > 0, "no time provided");
		CryptoTom storage tommer = tommers[_tokenId];
		tommer.farmTime = tommer.farmTime.add(_time);

		emit Working(_tokenId, _time);
	}

	function exp(uint256 _tokenId, uint256 _exp) public onlyBattlefield {
		require(_exp > 0, "no exp provided");
		CryptoTom storage tommer = tommers[_tokenId];
		tommer.exp = tommer.exp.add(_exp);

		emit Exp(_tokenId, _exp);
	}

	function evolve(
		uint256 _tokenId,
		address _owner,
		uint256 _dna
	) public onlySpawner {
		require(ownerOf(_tokenId) == _owner, "not own");
		CryptoTom storage tommer = tommers[_tokenId];
		require(!isEvolved[_tokenId], "already evolved");

		tommer.bornTime = block.timestamp;
		tommer.dna = _dna;
		isEvolved[_tokenId] = true;

		emit Evolve(_tokenId, _dna);
	}

	function changeTribe(
		uint256 _tokenId,
		address _owner,
		Tribe _tribe
	) external onlySpawner {
		require(ownerOf(_tokenId) == _owner, "not own");
		tomERC20.transferFrom(_msgSender(), owner(), feeChangeTribe);

		CryptoTom storage tommer = tommers[_tokenId];
		tommer.tribe = _tribe;

		emit UpdateTribe(_tokenId, _tribe);
	}

	function upgradeGeneration(
		uint256 _tokenId,
		address _owner,
		uint256 _generation
	) external onlySpawner {
		require(ownerOf(_tokenId) == _owner, "not own");
		tomERC20.transferFrom(_msgSender(), owner(), feeUpgradeGeneration);

		CryptoTom storage tommer = tommers[_tokenId];
		tommer.generation = _generation;

		emit UpgradeGeneration(_tokenId, _generation);
	}

	function layEgg(address _receiver, Tribe[] memory _tribes) external onlySpawner {
		uint256 amount = _tribes.length;
		require(amount > 0, "0 amount");
		if (amount == 1) {
			_layEgg(_receiver, _tribes[0]);
		} else {
			for (uint256 i = 0; i < amount; i++) {
				_layEgg(_receiver, _tribes[i]);
			}
		}
	}

	function _layEgg(address _receiver, Tribe _tribe) internal {
		_tokenIds.increment();
		uint256 newTokenId = _tokenIds.current();

		_mint(_receiver, newTokenId);

		tommers[newTokenId] = CryptoTom({
			generation: 1,
			tribe: _tribe,
			exp: 0,
			dna: 0,
			farmTime: 0,
			bornTime: block.timestamp
		});

		emit LayEgg(newTokenId, _receiver);
	}

	function getTommer(uint256 _tokenId) public view returns (CryptoTom memory) {
		return tommers[_tokenId];
	}

	function tommerLevel(uint256 _tokenId) public view returns (uint256) {
		return getLevel(getTommer(_tokenId).exp);
	}

	function getRare(uint256 _tokenId) public view returns (uint256) {
		uint256 dna = getTommer(_tokenId).dna;
		if (dna == 0)
			return 0;

		uint256 rareScore = dna / 10**26;
		if (rareScore < 5225) {
			return 1;
		} else if (rareScore < 7837) {
			return 2;
		} else if (rareScore < 8707) {
			return 3;
		} else if (rareScore < 9360) {
			return 4;
		} else if (rareScore < 9708) {
			return 5;
		} else {
			return 6;
		}
	}

	function getLevel(uint256 _exp) internal pure returns (uint256) {
		if (_exp < 100) {
			return 1;
		} else if (_exp < 350) {
			return 2;
		} else if (_exp < 1000) {
			return 3;
		}  else if (_exp < 2000) {
			return 4;
		} else if (_exp < 4000) {
			return 5;
		} else {
			return 6;
		}
	}

	function placeOrder(uint256 _tokenId, uint256 _price) public {
		require(ownerOf(_tokenId) == _msgSender(), "not own");
		require(_price > 0, "0 price");
		require(isEvolved[_tokenId], "not evolved");

		tokenOrder(_tokenId, true, _price);

		emit PlaceOrder(_tokenId, _msgSender(), _price);
	}

	function cancelOrder(uint256 _tokenId) public {
		require(tokenSales.contains(_tokenId), "not sale");
		ItemSale storage itemSale = markets[_tokenId];
		require(itemSale.owner == _msgSender(), "not own");

		tokenOrder(_tokenId, false, 0);

		emit CancelOrder(_tokenId, _msgSender());
	}

	function updatePrice(uint256 _tokenId, uint256 _price) public {
		require(_price > 0, "0 price");
		require(tokenSales.contains(_tokenId), "not sale");
		ItemSale storage itemSale = markets[_tokenId];
		require(itemSale.owner == _msgSender(), "not own");

		itemSale.price = _price;

		emit UpdatePrice(_tokenId, _msgSender(), _price);
	}

	function fillOrder(uint256 _tokenId) public {
		require(tokenSales.contains(_tokenId), "not sale");

		ItemSale storage itemSale = markets[_tokenId];
		uint256 fee = itemSale.price.mul(feeMarketRate).div(divPercent);
		tomERC20.transferFrom(_msgSender(), owner(), fee);
		tomERC20.transferFrom(_msgSender(), itemSale.owner, itemSale.price.sub(fee));
		tokenOrder(_tokenId, false, 0);

		emit FillOrder(_tokenId, _msgSender());
	}

	function tokenOrder(
		uint256 _tokenId,
		bool _sell,
		uint256 _price
	) internal {
		ItemSale storage itemSale = markets[_tokenId];
		if (_sell) {
			transferFrom(_msgSender(), address(this), _tokenId);
			tokenSales.add(_tokenId);
			sellerTokens[_msgSender()].add(_tokenId);

			markets[_tokenId] = ItemSale({
				tokenId: _tokenId,
				price: _price,
				owner: _msgSender()
			});
		} else {
			transferFrom(address(this), _msgSender(), _tokenId);

			tokenSales.remove(_tokenId);
			sellerTokens[itemSale.owner].remove(_tokenId);
			markets[_tokenId] = ItemSale({
				tokenId: 0,
				price: 0,
				owner: address(0)
			});
		}
	}

	function marketsSize() public view returns (uint256) {
		return tokenSales.length();
	}

	function orders(address _seller) public view returns (uint256) {
		return sellerTokens[_seller].length();
	}

	function tokenSaleByIndex(uint256 index) public view returns (uint256) {
		return tokenSales.at(index);
	}

	function tokenSaleOfOwnerByIndex(address _seller, uint256 index) public view returns (uint256) {
		return sellerTokens[_seller].at(index);
	}

	function getSale(uint256 _tokenId) public view returns (ItemSale memory) {
		if (tokenSales.contains(_tokenId)) {
			return markets[_tokenId];
		} else {
			return ItemSale({tokenId: 0, owner: address(0), price: 0});
		}
	}
}
