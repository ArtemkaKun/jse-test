package main

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

func addNewUserToBuffer(newUser User) {
	userId := newUser.Id

	usersCache[userId] = &newUser
	usersStatsCache[userId] = &UserStats{0, 0,
		0, 0, 0, 0}
}

func GetAllUserInfo(id uint64) AllUserStats {
	mainStats := usersCache[id]
	otherStats := usersStatsCache[id]

	return AllUserStats{mainStats.Id, mainStats.Balance, otherStats.DepositCount,
		otherStats.DepositSum, otherStats.BetCount, otherStats.BetSum,
		otherStats.WinCount, otherStats.WinSum}
}

func GetUserBalance(id uint64) float32 {
	return usersCache[id].Balance
}

func SetUserBalance(id uint64, newBalance float32) {
	usersCache[id].Balance = newBalance
}

func IncreaseUserDepositCount(id uint64) {
	usersStatsCache[id].DepositCount++
}

func IncreaseUserDepositSum(id uint64, newBalance float32) {
	usersStatsCache[id].DepositSum += newBalance
}

func IncreaseUserBetCount(id uint64) {
	usersStatsCache[id].BetCount++
}

func IncreaseUserBetSum(id uint64, newBalance float32) {
	usersStatsCache[id].BetSum += newBalance
}

func IncreaseUserWinCount(id uint64) {
	usersStatsCache[id].WinCount++
}

func IncreaseUserWinSum(id uint64, newBalance float32) {
	usersStatsCache[id].WinSum += newBalance
}
