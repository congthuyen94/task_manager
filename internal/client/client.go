// package client

// import (
// 	socketio "github.com/googollee/go-socket.io"
// 	"github.com/romana/rlog"
// )

// func Connect(url string) error {
// 	rlog.Info("Connect to: ", url)
// 	_, err := socketio.NewClient(url, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return err
// }
