package server

func (q *Queue) setupRouting() {
	q.apisvr.Router().GET("/v1/:appid/:topic/:ver", q.apisvr.handleAppendV1)
}
