// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface VRFCoordinatorInterface {
  function randomnessRequest(bytes32 keyHash) external returns (bytes32);
}
