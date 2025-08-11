## Read Functions

### 1. fees
- **Inputs**: `_index (uint256)`
- **Outputs**: `Fee`
- **Description**: Get fee details by index.

### 2. feesCount
- **Inputs**: None
- **Outputs**: `uint256`
- **Description**: Get the count of fees.

### 3. getAddresses
- **Inputs**: None
- **Outputs**: `Addresses`
- **Description**: Get GNS and GNS Staking addresses.

### 4. getAllPairsRestrictedMaxLeverage
- **Inputs**: None
- **Outputs**: `uint256[]`
- **Description**: Get all pairs' restricted max leverage.

### 5. groups
- **Inputs**: `_index (uint256)`
- **Outputs**: `Group`
- **Description**: Get group details by index.

### 6. groupsCount
- **Inputs**: None
- **Outputs**: `uint256`
- **Description**: Get the count of groups.

### 7. hasRole
- **Inputs**: `_account (address)`, `_role (uint8)`
- **Outputs**: `bool`
- **Description**: Check if an account has a specific role.

### 8. isPairIndexListed
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `bool`
- **Description**: Check if a pair index is listed.

### 9. isPairListed
- **Inputs**: `_from (string)`, `_to (string)`
- **Outputs**: `bool`
- **Description**: Check if a pair is listed.

### 10. pairCloseFeeP
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get pair close fee percentage.

### 11. pairCustomMaxLeverage
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get pair custom max leverage.

### 12. pairJob
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `(string, string)`
- **Description**: Get pair job details.

### 13. pairMaxLeverage
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get pair max leverage.

### 14. pairMinLeverage
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get pair min leverage.

### 15. pairMinPositionSizeUsd
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get pair min position size in USD.

### 16. pairOpenFeeP
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get pair open fee percentage.

### 17. pairOracleFeeP
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get pair oracle fee percentage.

### 18. pairSpreadP
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get pair spread percentage.

### 19. pairTriggerOrderFeeP
- **Inputs**: `_pairIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get pair trigger order fee percentage.

### 20. pairs
- **Inputs**: `_index (uint256)`
- **Outputs**: `Pair`
- **Description**: Get pair details by index.

### 21. pairsBackend
- **Inputs**: `_index (uint256)`
- **Outputs**: `(Pair, Group, Fee)`
- **Description**: Get pair, group, and fee details by index.

### 22. pairsCount
- **Inputs**: None
- **Outputs**: `uint256`
- **Description**: Get the count of pairs.

## Write Functions

### 1. addFees
- **Inputs**: `_fees (tuple[])`
- **Outputs**: None
- **Description**: Add fees.

### 2. addGroups
- **Inputs**: `_groups (tuple[])`
- **Outputs**: None
- **Description**: Add groups.

### 3. addPairs
- **Inputs**: `_pairs (tuple[])`
- **Outputs**: None
- **Description**: Add pairs.

### 4. initialize
- **Inputs**: `_rolesManager (address)`
- **Outputs**: None
- **Description**: Initialize roles.

### 5. setPairCustomMaxLeverages
- **Inputs**: `_indices (uint256[])`, `_values (uint256[])`
- **Outputs**: None
- **Description**: Set pair custom max leverages.

### 6. setRoles
- **Inputs**: `_accounts (address[])`, `_roles (uint8[])`, `_values (bool[])`
- **Outputs**: None
- **Description**: Set roles.

### 7. updateFees
- **Inputs**: `_ids (uint256[])`, `_fees (tuple[])`
- **Outputs**: None
- **Description**: Update fees.

### 8. updateGroups
- **Inputs**: `_ids (uint256[])`, `_groups (tuple[])`
- **Outputs**: None
- **Description**: Update groups.

### 9. updatePairs
- **Inputs**: `_ids (uint256[])`, `_pairs (tuple[])`
- **Outputs**: None
- **Description**: Update pairs.
