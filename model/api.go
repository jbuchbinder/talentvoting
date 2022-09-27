package model

func IsValidContest(contestId string) bool {
	var count int
	DB.Raw("SELECT COUNT(*) FROM contest WHERE code = ?", contestId).Scan(&count)
	return count == 1
}

func IsValidVoter(contestId, voterId string) bool {
	var count int
	DB.Raw("SELECT COUNT(*) FROM voter v LEFT OUTER JOIN contest c ON v.contest = c.id WHERE v.code = ? AND c.code = ?", voterId, contestId).Scan(&count)
	return count == 1
}
