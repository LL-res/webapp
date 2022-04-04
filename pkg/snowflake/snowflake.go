package snowflake

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixMilli()
	node, err = snowflake.NewNode(machineID)
	return
}
func GenID() string {
	s := node.Generate().String()
	fmt.Println(s)
	return s
}
