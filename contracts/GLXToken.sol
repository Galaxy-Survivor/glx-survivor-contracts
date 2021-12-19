// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./ERC20Capped.sol";
import "./AccessControl.sol";

contract GLXToken is ERC20Capped, AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    constructor()
        ERC20("Galaxy Token", "GLX")
        ERC20Capped(1e9 * 1e18)
    {
        _grantRole(MINTER_ROLE, msg.sender);
    }

    function mint(address to, uint256 amount)
        external
        onlyRole(MINTER_ROLE)
    {
        super._mint(to, amount);
    }
}
