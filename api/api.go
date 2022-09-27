package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitApi(m *gin.Engine) {
	g := m.Group("/api")
	g.GET("/contest/:contest/:voter", apiContest)
	g.POST("/registerVote/:contest/:voter", apiRegisterVote)
}

func apiContest(c *gin.Context) {
	contest := c.Param("contest")
	voter := c.Param("voter")
	if !validContest(contest) {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid contest."})
		return
	}
	if !validVoter(contest, voter) {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid voter id -- may have been previously used."})
		return
	}
	log.Printf("apiOptions(): contest = %s", contest)

	// TODO: FIXME: IMPLEMENT: XXXX

	c.JSON(http.StatusOK,
		gin.H{
			"name":         "Talent Show Contest",
			"description":  "Some description here",
			"maximumVotes": 2,
			"options": []gin.H{
				{"name": "Singer 1", "id": 1},
				{"name": "Singer 2", "id": 2},
				{"name": "Singer 3", "id": 3},
				{"name": "Singer 4", "id": 4},
				{"name": "Singer 5", "id": 5},
			},
		},
	)
}

func apiRegisterVote(c *gin.Context) {
	log.Printf("apiRegisterVote()")
	// TODO: FIXME: IMPLEMENT: XXXX
}

func validContest(contestId string) bool {
	// TODO: FIXME: IMPLEMENT: XXXX
	return contestId == "Y"
}

func validVoter(contestId, voterId string) bool {
	// TODO: FIXME: IMPLEMENT: XXXX
	return contestId == "Y" && voterId == "X"
}
