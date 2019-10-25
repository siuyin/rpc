package rpc

import (
	"fmt"

	"github.com/nats-io/nuid"
)

func Example() {
	// Encode req into some transport, eg. JSON or Protocol Buffers
	// then send the rpc request.
	// Below I have used NATS' unique ID generator to provide an ID.
	req := Request{Request: "helpMe", ID: nuid.Next()}

	// The RPC service will then reply with the response.
	// Unmarshal the response into the resp variable.
	// By convention the fist two items in the response are the
	// request name and ID respectively.
	resp := Response{"helpMe", "123456", "I can help!"}

	fmt.Println(req)
	fmt.Println(resp)
}
