package zbprotocol

// SingleMessageHeader is used to represent Consumer model of communication.
type SingleMessageHeader struct{} // 0 byte

// NewSingleMessageHeader constructor.
func NewSingleMessageHeader() *SingleMessageHeader {
	return &SingleMessageHeader{}
}
