## Read Functions

### 1. getAddresses
- **Inputs**: None
- **Outputs**: `Addresses`
- **Description**: Get GNS and GNS Staking addresses.

### 2. getFeeTier
- **Inputs**: `_feeTierIndex (uint256)`
- **Outputs**: `FeeTier`
- **Description**: Get fee tier by index.

### 3. getFeeTiersCount
- **Inputs**: None
- **Outputs**: `uint256`
- **Description**: Get the count of fee tiers.

### 4. getFeeTiersTraderDailyInfo
- **Inputs**: `_trader (address)`, `_day (uint32)`
- **Outputs**: `TraderDailyInfo`
- **Description**: Get fee tiers trader daily info.

### 5. getFeeTiersTraderInfo
- **Inputs**: `_trader (address)`
- **Outputs**: `TraderInfo`
- **Description**: Get fee tiers trader info.

### 6. getGroupVolumeMultiplier
- **Inputs**: `_groupIndex (uint256)`
- **Outputs**: `uint256`
- **Description**: Get group volume multiplier.

### 7. hasRole
- **Inputs**: `_account (address)`, `_role (uint8)`
- **Outputs**: `bool`
- **Description**: Check if an account has a specific role.

## Write Functions

### 1. initialize
- **Inputs**: `_rolesManager (address)`
- **Outputs**: None
- **Description**: Initialize roles.

### 2. initializeFeeTiers
- **Inputs**: `_groupIndices (uint256[])`, `_groupVolumeMultipliers (uint256[])`, `_feeTiersIndices (uint256[])`, `_feeTiers (tuple[])`
- **Outputs**: None
- **Description**: Initialize fee tiers.

### 3. setFeeTiers
- **Inputs**: `_feeTiersIndices (uint256[])`, `_feeTiers (tuple[])`
- **Outputs**: None
- **Description**: Set fee tiers.

### 4. setGroupVolumeMultipliers
- **Inputs**: `_groupIndices (uint256[])`, `_groupVolumeMultipliers (uint256[])`
- **Outputs**: None
- **Description**: Set group volume multipliers.

### 5. setRoles
- **Inputs**: `_accounts (address[])`, `_roles (uint8[])`, `_values (bool[])`
- **Outputs**: None
- **Description**: Set roles.

### 6. updateTraderPoints
- **Inputs**: `_trader (address)`, `_volumeUsd (uint256)`, `_pairIndex (uint256)`
- **Outputs**: None
- **Description**: Update trader points.
