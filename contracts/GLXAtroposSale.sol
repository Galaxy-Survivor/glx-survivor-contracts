// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./Context.sol";
import "./Ownable.sol";
import "./IERC20.sol";
import "./SafeERC20.sol";
import "./VRFConsumerBaseV2.sol";
import "./VRFCoordinatorV2Interface.sol";

interface IERC721Mint {
    function mint(address to, uint32 empire, uint32 rarity, uint32 unlockTime) external;
}

interface IERC1155Mint {
    function mint(address to, uint256 id, uint256 amount) external;
    function burn(address from, uint256 id, uint256 amount) external;
    function balanceOf(address owner, uint256 id) external view returns (uint256);
}

contract GLXAtroposSale is VRFConsumerBaseV2, Context, Ownable {
    using SafeERC20 for IERC20;

    uint256 public constant MAX_TOTAL_NORMAL_BPS = 644;
    uint256 public constant MAX_TOTAL_RARE_BPS = 270;
    uint256 public constant MAX_TOTAL_MYTHICAL_BPS = 86;

    uint256 public constant NORMAL_BP_ID = 1;
    uint256 public constant RARE_BP_ID = 2;
    uint256 public constant MYTHICAL_BP_ID = 3;

    uint256 public constant NORMAL_BP_PRICE = 100 * 10**6;
    uint256 public constant RARE_BP_PRICE = 150 * 10 ** 6;
    uint256 public constant MYTHICAL_BP_PRICE = 300 * 10 ** 6;

    uint256 private soldNormalBps = 0;
    uint256 private soldRareBps = 0;
    uint256 private soldMythicalBps = 0;

    uint256 private openedNormalBps = 0;
    uint256 private openedRareBps = 0;
    uint256 private openedMythicalBps = 0;

    uint256 private remainShips = 1000;
    uint256 private remainNormalBps = 644;
    uint256 private remainRareBps = 270;
    uint256 private remainMythicalBps = 86;

    uint256 private remainNN = 500;
    uint256 private remainNR = 116;
    uint256 private remainNSR = 28;

    uint256 private remainRR = 184;
    uint256 private remainRSR = 73;
    uint256 private remainRSSR = 13;

    uint256 private remainMSR = 39;
    uint256 private remainMSSR = 37;
    uint256 private remainMUR = 10;

    struct RandomnessRequest {
        address sender;
        uint256 blueprintId;
    }

    bytes32 private _keyHash;
    VRFCoordinatorV2Interface private _vrfCoordinator;
    mapping(uint256 => RandomnessRequest) private _randomnessRequests;
    uint32 private _callbackGasLimit = 500000;
    uint16 private _requestConfirmations = 3;
    uint64 private _subscriptionId;

    uint256 public startTime;
    uint256 public endTime;

    IERC1155Mint private _item;
    IERC721Mint private _ship;
    IERC20 private _token;
    IERC1155Mint private _ticket;

    event BlueprintMinted(address indexed owner, uint256 blueprintId, uint256 amount);
    event BlueprintOpened(address indexed owner, uint256 blueprintId, uint256 amount);

    constructor(
        address glxItem,
        address glxShip,
        address token,
        address ticket,
        address vrfCoordinator,
        bytes32 keyHash,
        uint64 subscriptionId,
        uint256 stime,
        uint256 etime
    )
        VRFConsumerBaseV2(vrfCoordinator)
    {
        _vrfCoordinator = VRFCoordinatorV2Interface(vrfCoordinator);
        _keyHash = keyHash;
        _subscriptionId = subscriptionId;

        startTime = stime;
        endTime = etime;

        _item = IERC1155Mint(glxItem);
        _ship = IERC721Mint(glxShip);
        _token = IERC20(token);
        _ticket = IERC1155Mint(ticket);
    }

    modifier onlyValidTime() {
        require(block.timestamp >= startTime, "GLXAtroposSale: sale not start yet");
        require(block.timestamp <= endTime, "GLXAtroposSale: sale ended");
        _;
    }

    function buyNormal(uint256 amount, uint256 ticketId) external onlyValidTime {
        require(ticketId == 1 || ticketId == 2 || ticketId == 3, "GLXAtroposSale: invalid ticket");
        require(soldNormalBps + amount <= MAX_TOTAL_NORMAL_BPS, "GLXAtroposSale: amount exceed limit");

        uint256 ticketBalance = _ticket.balanceOf(_msgSender(), ticketId);
        require(ticketBalance >= amount, "GLXAtroposSale: not enough ticket");

        uint256 totalPrice = amount * (NORMAL_BP_PRICE - _getDiscountByTicketId(ticketId));
        uint256 balance = _token.balanceOf(_msgSender());
        require(totalPrice <= balance, "insufficient balance");

        soldNormalBps += amount;

        _ticket.burn(_msgSender(), ticketId, amount);
        _token.safeTransferFrom(_msgSender(), address(this), totalPrice);
        _item.mint(_msgSender(), NORMAL_BP_ID, amount);

        emit BlueprintMinted(_msgSender(), NORMAL_BP_ID, amount);
    }

    function buyRare(uint256 amount, uint256 ticketId) external onlyValidTime {
        require(ticketId == 2 || ticketId == 3, "GLXAtroposSale: invalid ticket");
        require(soldRareBps + amount <= MAX_TOTAL_RARE_BPS, "GLXAtroposSale: amount exceed limit");

        uint256 ticketBalance = _ticket.balanceOf(_msgSender(), ticketId);
        require(ticketBalance >= amount, "GLXAtroposSale: not enough ticket");

        uint256 totalPrice = amount * (RARE_BP_PRICE - _getDiscountByTicketId(ticketId));
        uint256 balance = _token.balanceOf(_msgSender());
        require(totalPrice <= balance, "insufficient balance");

        soldRareBps += amount;

        _ticket.burn(_msgSender(), ticketId, amount);
        _token.safeTransferFrom(_msgSender(), address(this), totalPrice);
        _item.mint(_msgSender(), RARE_BP_ID, amount);

        emit BlueprintMinted(_msgSender(), RARE_BP_ID, amount);
    }

    function buyMythical(uint256 amount, uint256 ticketId) external onlyValidTime {
        require(ticketId == 3, "GLXAtroposSale: invalid ticket");
        require(soldMythicalBps + amount <= MAX_TOTAL_MYTHICAL_BPS, "GLXAtroposSale: amount exceed limit");

        uint256 ticketBalance = _ticket.balanceOf(_msgSender(), ticketId);
        require(ticketBalance >= amount, "GLXAtroposSale: not enough ticket");

        uint256 totalPrice = amount * (MYTHICAL_BP_PRICE - _getDiscountByTicketId(ticketId));
        uint256 balance = _token.balanceOf(_msgSender());
        require(totalPrice <= balance, "insufficient balance");

        soldMythicalBps += amount;

        _ticket.burn(_msgSender(), ticketId, amount);
        _token.safeTransferFrom(_msgSender(), address(this), totalPrice);
        _item.mint(_msgSender(), MYTHICAL_BP_ID, amount);

        emit BlueprintMinted(_msgSender(), MYTHICAL_BP_ID, amount);
    }

    function unbox(uint256 blueprintId, uint256 amount) external {
        require(
            blueprintId == NORMAL_BP_ID || blueprintId == RARE_BP_ID || blueprintId == MYTHICAL_BP_ID,
            "GLXAtroposSale: invalid blueprint id"
        );
        require(amount > 0, "GLXAtroposSale: amount is zero");

        if (blueprintId == NORMAL_BP_ID) {
            openedNormalBps += amount;
        } else if (blueprintId == RARE_BP_ID) {
            openedRareBps += amount;
        } else {
            openedMythicalBps += amount;
        }

        for (uint256 i = 0; i < amount; i++) {
            uint256 requestID = _vrfCoordinator.requestRandomWords(
                _keyHash,
                _subscriptionId,
                _requestConfirmations,
                _callbackGasLimit,
                1
            );
            RandomnessRequest storage req = _randomnessRequests[requestID];
            req.sender = _msgSender();
            req.blueprintId = blueprintId;
        }

        _item.burn(_msgSender(), blueprintId, amount);

        emit BlueprintOpened(_msgSender(), blueprintId, amount);
    }

    function fulfillRandomWords(uint256 requestID, uint256[] memory randomWords) internal override {
        RandomnessRequest memory req = _randomnessRequests[requestID];
        require(req.sender != address(0x0), "GLXAtroposSale: invalid request id");
        require(randomWords.length >= 1, "GLXAtroposSale: no random words provided");
        require(remainShips > 0, "GLXAtroposSale: all blueprint opened");

        delete _randomnessRequests[requestID];
        remainShips -= 1;

        uint256 randomness = randomWords[0];
        if (req.blueprintId == NORMAL_BP_ID) {
            randomness = randomness % remainNormalBps + 1;
            remainNormalBps -= 1;
            if (randomness <= remainNN) {
                remainNN -= 1;
                _ship.mint(req.sender, 2, 1, 0);
                return;
            }

            randomness -= remainNN;
            if (randomness <= remainNR) {
                remainNR -= 1;
                _ship.mint(req.sender, 2, 2, 0);
                return;
            }

            remainNSR -= 1;
            _ship.mint(req.sender, 2, 3, 0);
            return;
        } else if (req.blueprintId == RARE_BP_ID) {
            randomness = randomness % remainRareBps + 1;
            remainRareBps -= 1;
            if (randomness <= remainRR) {
                remainRR -= 1;
                _ship.mint(req.sender, 2, 2, 0);
                return;
            }

            randomness -= remainRR;
            if (randomness <= remainRSR) {
                remainRSR -= 1;
                _ship.mint(req.sender, 2, 3, 0);
                return;
            }

            remainRSSR -= 1;
            _ship.mint(req.sender, 2, 4, 0);
            return;
        } else {
            randomness = randomness % remainMythicalBps + 1;
            remainMythicalBps -= 1;
            if (randomness <= remainMSR) {
                remainMSR -= 1;
                _ship.mint(req.sender, 2, 3, 0);
                return;
            }

            randomness -= remainMSR;
            if (randomness <= remainMSSR) {
                remainMSSR -= 1;
                _ship.mint(req.sender, 2, 4, 0);
                return;
            }

            remainMUR -= 1;
            _ship.mint(req.sender, 2, 5, 0);
            return;
        }
    }

    function withdraw(address to, uint256 amount) external onlyOwner {
        _token.safeTransfer(to, amount);
    }

    function getSoldNormals() external view returns (uint256) {
        return soldNormalBps;
    }

    function getSoldRares() external view returns (uint256) {
        return soldRareBps;
    }

    function getSoldMythicals() external view returns (uint256) {
        return soldMythicalBps;
    }

    function _getDiscountByTicketId(uint256 ticketId) private pure returns (uint256) {
        if (ticketId == 1) {
            return 10**6;
        }
        if (ticketId == 2) {
            return 5*10**6;
        }
        if (ticketId == 3) {
            return 10*10**6;
        }
        return 0;
    }
}
