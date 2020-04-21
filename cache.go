package jse_test

var usersCache map[uint64]*User
var usersStatsCache map[uint64]*UserStats

func init() {
	usersCache = make(map[uint64]*User)
	usersStatsCache = make(map[uint64]*UserStats)
}

func isIdExist(id uint64) bool {
	if usersCache[id] == nil {
		return false
	}

	return true
}

func addNewUserToBuffer(new_user User) {
	user_id := new_user.Id

	usersCache[user_id] = &new_user
	usersStatsCache[user_id] = &UserStats{0, 0,
		0, 0, 0, 0}
}

func GetAllUserInfo(id uint64) AllUserStats {
	main_stats := usersCache[id]
	other_stats := usersStatsCache[id]

	return AllUserStats{main_stats.Id, main_stats.Balance, other_stats.DepositCount,
		other_stats.DepositSum, other_stats.BetCount, other_stats.BetSum,
		other_stats.WinCount, other_stats.WinSum}
}

func GetUserBalance(id uint64) float32 {
	return usersCache[id].Balance
}

func SetUserBalance(id uint64, new_balance float32) {
	usersCache[id].Balance = new_balance
}

func IncreaseUserDepositCount(id uint64) {
	usersStatsCache[id].DepositCount++
}

func IncreaseUserDepositAmount(id uint64, new_balance float32) {
	usersStatsCache[id].DepositSum += new_balance
}
