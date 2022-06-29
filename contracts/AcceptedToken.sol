// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./Ownable.sol";
import "./EnumerableSet.sol";

contract AcceptedToken is Ownable {
    using EnumerableSet for EnumerableSet.AddressSet;

    EnumerableSet.AddressSet internal acceptedTokens;

    event TokensAdded(address[] tokens);
    event TokensRemoved(address[] tokens);

    function addAcceptedTokens(address[] memory tokens) public onlyOwner {
        for (uint256 i = 0; i < tokens.length; i++) {
            acceptedTokens.add(tokens[i]);
        }
        emit TokensAdded(tokens);
    }

    function removeAcceptedTokens(address[] memory tokens) external onlyOwner {
        for (uint256 i = 0; i < tokens.length; i++) {
            acceptedTokens.remove(tokens[i]);
        }
        emit TokensRemoved(tokens);
    }

    function isAcceptedToken(address token) external view returns (bool) {
        return acceptedTokens.contains(token);
    }

    function _isAcceptedToken(address token) internal view returns (bool) {
        return acceptedTokens.contains(token);
    }

    function totalAcceptedTokens() external view returns (uint256) {
        return acceptedTokens.length();
    }

    function acceptedTokenAt(uint256 i) external view returns (address) {
        return acceptedTokens.at(i);
    }
}
