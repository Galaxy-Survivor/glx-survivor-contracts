// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./IERC20.sol";
import "./SafeERC20.sol";
import "./IERC721.sol";
import "./IERC1155.sol";
import "./Context.sol";
import "./Ownable.sol";
import "./EnumerableSet.sol";

/**
 * @title Galaxy Survivor Marketplace.
 * @dev Implement a contract for buying and selling ships and items.
 */
contract GLXMarket is Context, Ownable {
    using EnumerableSet for EnumerableSet.UintSet;
    using SafeERC20 for IERC20;

    enum ItemType {
    	GLXShip,
    	GLXItem,
    	GLXBox
    }

    enum OrderSide {
    	Buy,
	Sell
    }

    /**
     * @dev Data structure that contains information for an order.
     * @param creator is the address of the order's creator.
     * @param itemId is the id of item want to buy/sell.
     * @param price is price for buying/selling one item.
     * @param createdTime is timestamp when the order is created.
     * @param expiredTime is the timestamp when the order will be expired.
     * @param amount is the number of items want to buy/sell.
     * @param remainAmount is the remaining amount can be bought/sold.
     * @param itemType is the type of an item.
     * @param side is whether buy or sell.
     */
    struct Order {
    	address creator;
    	uint256 itemId;
    	uint256 price;
    	uint64 createdTime;
    	uint64 expiredTime;
    	uint32 amount;
    	uint32 remainAmount;
    	ItemType itemType;
    	OrderSide side;
    }

    uint256 constant private ONE_BPS = 10000;
    uint256 constant private MIN_ORDER_DURATION = 10 minutes;

    IERC20 private _token;
    IERC721 private _ship;
    IERC721 private _item;
    IERC1155 private _fungibleItem;

    uint256 private _baseFee = 100;

    uint256 public totalOrders;

    mapping(uint256 => Order) private orders;
    mapping(address => EnumerableSet.UintSet) userOrderIds;

    /**
     * @dev Emitted when a new fee is set.
     */
    event SetFee(address indexed operator, uint256 oldFee, uint256 newFee);

    /**
     * @dev Emitted when an order is successfully created.
     */
    event CreateOrder(address indexed operator, uint256 indexed orderId);

    /**
     * @dev Emitted when an order is successfully cancelled.
     */
    event CancelOrder(address indexed operator, uint256 indexed orderId);

    /**
     * @dev Emitted when an order is taken.
     */
    event TakeOrder(
    	address indexed seller,
    	address indexed buyer,
    	uint256 indexed orderId,
    	uint256 amount,
    	uint256 totalPrice,
    	uint256 totalFee
    );

    /**
     * @dev Sets addresses for contracts GLXToken, GLXShip, GLXItem, and GLXFungibleItem.
     */
    constructor(address token, address ship, address item, address fungibleItem) {
    	require(token != address(0x0), "token address is 0x0");
    	require(ship != address(0x0), "ship address is 0x0");
    	require(item != address(0x0), "item address is 0x0");
    	require(fungibleItem != address(0x00), "fungible item address is 0x0");
    	_token = IERC20(token);
    	_ship = IERC721(ship);
    	_item = IERC721(item);
    	_fungibleItem = IERC1155(fungibleItem);
    }

    /**
     * @dev Sets a new fee value for market.
     *
     * The unit for fee value is 1/10000.
     */ 
    function setFee(uint256 fee) external onlyOwner {
    	require(fee > 0, "zero fee");
    	uint256 oldFee = _baseFee;
    	_baseFee = fee;
    	emit SetFee(_msgSender(), oldFee, _baseFee);
    }

    /**
     * @dev Withdraws {amount} GLXToken to {to} address.
     */
    function withdraw(address to, uint256 amount) external onlyOwner {
    	_token.safeTransfer(to, amount);
    }

    /**
     * @dev Creates a new order for buying/selling items.
     * @param itemId is id of item want to buy/sell.
     * @param itemType is the type of item.
     * @param price is the price for an item.
     * @param amount is the number of items want to buy/sell. This value is ignore for GLXShip and GLXItem.
     * @param duration is the valid duration time in seconds of the order.
     * @param side is whether buy or sell.
     */
    function createOrder(
    	uint256 itemId,
    	ItemType itemType,
    	uint256 price,
    	uint32 amount,
    	uint64 duration,
    	OrderSide side
    ) external returns (uint256 orderId) {
    	require(price > 0, "zero price");
    	require(duration >= MIN_ORDER_DURATION, "duration too short");

    	if (itemType == ItemType.GLXBox) {
    		require(amount > 0, "amount is zero");
    	} else {
    		amount = 1;
    	}

    	if (side == OrderSide.Sell) {
    		bool isValidOwner = false;
    		if (itemType == ItemType.GLXShip) {
    			address itemOwner = _ship.ownerOf(itemId);
    			isValidOwner = _msgSender() == itemOwner;
    		} else if (itemType == ItemType.GLXItem) {
    			address itemOwner = _item.ownerOf(itemId);
    			isValidOwner = _msgSender() == itemOwner;
    		} else if (itemType == ItemType.GLXBox) {
    			uint256 availableAmount = _fungibleItem.balanceOf(_msgSender(), itemId);
    			isValidOwner = availableAmount >= amount;
    		}
    		require(isValidOwner, "insufficient balance");
    	} else {
    		uint256 balance = _token.balanceOf(_msgSender());
    		require(balance >= price * amount, "insufficient balance");
    	}

    	orderId = ++totalOrders;
    	uint64 currentTime = _getBlockTimestamp();

    	orders[orderId] = Order({
    		creator: _msgSender(),
    		itemId: itemId,
    		itemType: itemType,
    		price: price,
    		amount: amount,
    		remainAmount: amount,
    		createdTime: currentTime,
    		expiredTime: currentTime + duration,
    		side: side
    	});
    	userOrderIds[_msgSender()].add(orderId);

    	emit CreateOrder(_msgSender(), orderId);

    	return orderId;
    }

    /**
     * @dev Cancels an existing order.
     */
    function cancelOrder(uint256 orderId) external {
    	Order storage order = orders[orderId];
    	require(
		(order.creator == _msgSender()) || (owner() == _msgSender()),
		"permission denied"
	);
    	require(order.remainAmount > 0, "sold out");
    	userOrderIds[order.creator].remove(orderId);
    	delete orders[orderId];
    	emit CancelOrder(_msgSender(), orderId);
    }

    /**
     * @dev Take an amount of items from an order.
     */
    function takeOrder(uint256 orderId, uint32 amount) external {
    	Order storage order = orders[orderId];

    	require(order.creator != address(0x0), "order not found");
    	require(order.creator != _msgSender(), "cannot take own order");
    	require(amount > 0, "amount is zero");
    	require(order.remainAmount >= amount, "amount is too large");
    	require(order.expiredTime >= _getBlockTimestamp(), "order expired");

    	uint256 totalPrice = amount * order.price;
    	uint256 totalFee = totalPrice * _baseFee / ONE_BPS;
    	order.remainAmount -= amount;

    	if (order.remainAmount == 0) {
    		userOrderIds[order.creator].remove(orderId);
    	}

    	address seller;
    	address buyer;
    	if (order.side == OrderSide.Sell) {
    		seller = order.creator;
    		buyer = _msgSender();
    	} else {
    		seller = _msgSender();
    		buyer = order.creator;
    	}
    	emit TakeOrder(seller, buyer, orderId, amount, totalPrice, totalFee);

    	_token.safeTransferFrom(buyer, address(this), totalFee);
    	_token.safeTransferFrom(buyer, order.creator, totalPrice-totalFee);
    	if (order.itemType == ItemType.GLXShip) {
    		_ship.safeTransferFrom(seller, buyer, order.itemId);
    	} else if (order.itemType == ItemType.GLXItem) {
    		_item.safeTransferFrom(seller, buyer, order.itemId);
    	} else {
    		_fungibleItem.safeTransferFrom(seller, buyer, order.itemId, amount, "");
    	}
    }

    /**
     * @dev Returns order's information.
     */
    function getOrder(uint256 orderId) external view returns (Order memory) {
    	return orders[orderId];
    }

    /**
     * @dev Returns a list of orders' information from {startIndex} to {endIndex}.
     */
    function getOrders(
	    uint256 startIndex,
	    uint256 endIndex
    ) external view returns (Order[] memory) {
    	Order[] memory result = new Order[](endIndex - startIndex + 1);
    	for (uint256 i = startIndex; i <= endIndex; i++) {
    		result[i-startIndex] = orders[i];
    	}
    	return result;
    }

    function _getBlockTimestamp() internal view returns (uint64) {
    	return uint64(block.timestamp);
    }
}
