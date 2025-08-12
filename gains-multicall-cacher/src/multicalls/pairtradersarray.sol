// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@makerdao/multicall/contracts/Multicall.sol";

contract GetPairTradersArray {
    Multicall public multicall; // Multicall contract instance

    constructor(address _multicallAddress) {
        multicall = Multicall(_multicallAddress);
    }

    function getPairTradersArrays(address[] memory _addresses, uint256[] memory _pairIndexes) public view returns (address[][] memory) {
        uint256 len = _addresses.length;
        address[][] memory results = new address[][](len);

        for (uint256 i = 0; i < len; i++) {
            bytes memory data = abi.encodeWithSignature("pairTradersArray(address,uint256)", _addresses[i], _pairIndexes[i]);
            (bool success, bytes memory result) = address(multicall).staticcall(data);

            require(success, "Call failed");

            // Parse the result
            (address[] memory pairTradersArray) = abi.decode(result, (address[]));

            results[i] = pairTradersArray;
        }

        return results;
    }
}
