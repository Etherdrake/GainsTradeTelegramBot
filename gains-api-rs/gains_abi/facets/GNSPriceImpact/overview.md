### Read Functions

1. **getAddresses**
    - **Description**: Retrieves addresses stored in the contract.
    - **Parameters**: None
    - **Returns**: Tuple containing addresses (`gns` and `gnsStaking`).
    - **State Mutability**: View

2. **getOiWindow**
    - **Description**: Retrieves open interest data for a specific window.
    - **Parameters**: `_windowsDuration` (uint48), `_pairIndex` (uint256), `_windowId` (uint256)
    - **Returns**: Tuple containing open interest data for long and short positions.
    - **State Mutability**: View

3. **getOiWindows**
    - **Description**: Retrieves open interest data for multiple windows.
    - **Parameters**: `_windowsDuration` (uint48), `_pairIndex` (uint256), `_windowIds` (uint256[])
    - **Returns**: Array of tuples containing open interest data for long and short positions.
    - **State Mutability**: View

4. **getOiWindowsSettings**
    - **Description**: Retrieves settings for open interest windows.
    - **Parameters**: None
    - **Returns**: Tuple containing settings for open interest windows (start timestamp, duration, count).
    - **State Mutability**: View

5. **getPairDepth**
    - **Description**: Retrieves depth data for a specific pair.
    - **Parameters**: `_pairIndex` (uint256)
    - **Returns**: Tuple containing depth data (one percent depth above and below).
    - **State Mutability**: View

6. **getPairDepths**
    - **Description**: Retrieves depth data for multiple pairs.
    - **Parameters**: `_indices` (uint256[]), `_depthsAboveUsd` (uint128[]), `_depthsBelowUsd` (uint128[])
    - **Returns**: Array of tuples containing depth data (one percent depth above and below) for each pair.
    - **State Mutability**: View

7. **getPriceImpactOi**
    - **Description**: Retrieves price impact open interest for a specific pair and position.
    - **Parameters**: `_pairIndex` (uint256), `_long` (bool)
    - **Returns**: Active open interest value.
    - **State Mutability**: View

8. **getTradePriceImpact**
    - **Description**: Calculates trade price impact for a specific trade.
    - **Parameters**: `_openPrice` (uint256), `_pairIndex` (uint256), `_long` (bool), `_tradeOpenInterestUsd` (uint256)
    - **Returns**: Price impact percentage and price after impact.
    - **State Mutability**: View

9. **hasRole**
    - **Description**: Checks if an account has a specific role.
    - **Parameters**: `_account` (address), `_role` (enum IAddressStore.Role)
    - **Returns**: Boolean indicating role existence.
    - **State Mutability**: View

### Write Functions

1. **initialize**
    - **Description**: Initializes the contract.
    - **Parameters**: `_rolesManager` (address)
    - **State Mutability**: Non-payable

2. **initializePriceImpact**
    - **Description**: Initializes price impact settings.
    - **Parameters**: `_windowsDuration` (uint48), `_windowsCount` (uint48)
    - **State Mutability**: Non-payable

3. **addPriceImpactOpenInterest**
    - **Description**: Adds price impact open interest.
    - **Parameters**: `_openInterestUsd` (uint256), `_pairIndex` (uint256), `_long` (bool)
    - **State Mutability**: Non-payable

4. **removePriceImpactOpenInterest**
    - **Description**: Removes price impact open interest.
    - **Parameters**: `_openInterestUsd` (uint256), `_pairIndex` (uint256), `_long` (bool), `_addTs` (uint48)
    - **State Mutability**: Non-payable

5. **setPairDepths**
    - **Description**: Sets depth data for multiple pairs.
    - **Parameters**: `_indices` (uint256[]), `_depthsAboveUsd` (uint128[]), `_depthsBelowUsd` (uint128[])
    - **State Mutability**: Non-payable

6. **setPriceImpactWindowsCount**
    - **Description**: Sets the count of price impact windows.
    - **Parameters**: `_newWindowsCount` (uint48)
    - **State Mutability**: Non-payable

7. **setPriceImpactWindowsDuration**
    - **Description**: Sets the duration of price impact windows.
    - **Parameters**: `_newWindowsDuration` (uint48)
    - **State Mutability**: Non-payable

8. **setRoles**
    - **Description**: Sets roles for multiple accounts.
    - **Parameters**: `_accounts` (address[]), `_roles` (enum IAddressStore.Role[]), `_values` (bool[])
    - **State Mutability**: Non-payable
