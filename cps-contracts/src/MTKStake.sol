// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract MTKStake {
    //枚举质押期限
    enum StakingPeriod {
        ThirtyDays,
        NinetyDays,
        HundredEightyDays,
        OneYear
    }

    //质押订单结构体
    struct Stake {
        uint256 stakeId; //质押订单ID
        uint256 amount; //质押数量
        uint256 startTime; //质押开始时间
        uint256 endTime; //质押结束时间
        uint256 rewardRateYear; //质押奖励率(收益率)
        bool isActive; //质押订单是否有效
        StakingPeriod period; //质押期限
    }

    IERC20 public stakingToken; //质押的MTK代币

    mapping(address => Stake[]) public userStakes; //用户地址到质押订单数组的映射
    mapping(uint256 => address) public stakeIdOwner; //质押订单ID到用户地址的映射
    mapping(StakingPeriod => uint256) public apy; //质押期限>奖励率
    mapping(StakingPeriod => uint256) public durations; //质押期限>持续时间(查询时换算为秒)
    uint256 private nonce; //质押订单ID自增计数器

    event Staked(
        address indexed user,
        uint256 stakeId,
        uint256 amount,
        StakingPeriod period,
        uint256 timestamp
    );
    event Withdraw(
        address indexed user,
        uint256 stakeId,
        uint256 principal,
        uint256 reward,
        uint256 totalAmount
    );

    constructor(IERC20 _mtkToken) {
        stakingToken = _mtkToken;

        //设置不同质押期限的年化收益率
        apy[StakingPeriod.ThirtyDays] = 10; //30天质押年化10%
        apy[StakingPeriod.NinetyDays] = 15; //90天质押年化15%
        apy[StakingPeriod.HundredEightyDays] = 18; //180天质押年化18%
        apy[StakingPeriod.OneYear] = 20; //365天质押年化20%

        //设置不同质押期限的持续时间
        durations[StakingPeriod.ThirtyDays] = 30 days;
        durations[StakingPeriod.NinetyDays] = 90 days;
        durations[StakingPeriod.HundredEightyDays] = 180 days;
        durations[StakingPeriod.OneYear] = 365 days;
    }

    //质押函数
    function stake(uint256 amount, StakingPeriod period) external {
        require(amount > 0, "Amount must be greater than zero");
        require(apy[period] > 0, "Invalid staking period");

        //转移用户的MTK代币到质押合约
        require(stakingToken.transferFrom(msg.sender, address(this), amount),"Token transfer failed");

        //创建质押订单
        uint256 stakeId = _generateStakeId(); //生成唯一的质押订单ID
        uint256 duration = _getDuration(period); //获取质押期限对应的持续时间

        uint256 rewardRateYear = apy[period] * durations[period]; //计算质押奖励率(年化收益率*时间) 确保精度不损失,最后/365 days

        Stake memory newStake = Stake({
            stakeId: stakeId,
            amount: amount,
            startTime: block.timestamp,
            endTime: block.timestamp + duration,
            rewardRateYear: rewardRateYear,
            isActive: true,
            period: period
        });

        userStakes[msg.sender].push(newStake);
        stakeIdOwner[stakeId] = msg.sender;

        emit Staked(msg.sender, stakeId, amount, period, block.timestamp);
    }

    //生成唯一的质押订单ID
    function _generateStakeId() private returns (uint256) {
        return
            uint256(
                keccak256(
                    abi.encodePacked(block.timestamp, msg.sender, nonce++)
                )
            );
    }

    //获取质押期限对应的持续时间（测试环境用分钟）
    function _getDuration(
        StakingPeriod period
    ) internal pure returns (uint256) {
        if (period == StakingPeriod.ThirtyDays) {
            return 1 minutes;
        } else if (period == StakingPeriod.NinetyDays) {
            return 3 minutes;
        } else if (period == StakingPeriod.HundredEightyDays) {
            return 5 minutes;
        } else {
            return 10 minutes;
        }
    }

    //提取函数
    function withdraw(uint256 stakeId) external {
        require(stakeIdOwner[stakeId] == msg.sender, "Not the owner of this stake");
        Stake storage stk;
        uint256 stakeIndex;
        (stk, stakeIndex) = _getStakeById(msg.sender, stakeId);

        require(stk.isActive, "Stake is not active");
        require(block.timestamp >= stk.endTime, "Staking period is not over");

        stk.isActive = false;

        //计算收益
        uint256 reward = (stk.amount * stk.rewardRateYear) / 100 / 365 days; //确保精度不损失,先乘后除
        uint256 totalAmount = stk.amount + reward;

        //转移本金和奖励给用户
        require(stakingToken.transfer(msg.sender, totalAmount), "Token transfer failed");

        emit Withdraw(msg.sender, stakeId, stk.amount, reward, totalAmount);
    }

    //根据质押订单ID获取质押订单
    function _getStakeById(
        address user,
        uint256 stakeId
    ) internal view returns (Stake storage, uint256) {
        Stake[] storage stakes = userStakes[user];
        for (uint256 i = 0; i < stakes.length; i++) {
            if (stakes[i].stakeId == stakeId) {
                return (stakes[i], i);
            }
        }
        revert("Stake not found");
    }

    //获取用户的所有活跃的质押订单
    function getUserActiveStakes(
        address user
    ) external view returns (Stake[] memory) {
        Stake[] storage stakes = userStakes[user];
        uint256 activeCount = 0;

        // 计算活跃质押订单的数量
        for (uint256 i = 0; i < stakes.length; i++) {
            if (stakes[i].isActive) {
                activeCount++;
            }
        }

        // 创建并填充活跃质押订单数组
        Stake[] memory activeStakes = new Stake[](activeCount);
        uint256 index = 0;
        for (uint256 i = 0; i < stakes.length; i++) {
            if (stakes[i].isActive) {
                activeStakes[index] = stakes[i];
                index++;
            }
        }

        return activeStakes;
    }
}
