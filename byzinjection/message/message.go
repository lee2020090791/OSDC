package message

import "encoding/gob"

func init() {
	gob.Register(Transaction{})
	gob.Register(TransactionReply{})
	gob.Register(Query{})
	gob.Register(QueryReply{})
	gob.Register(Read{})
	gob.Register(ReadReply{})
	gob.Register(Register{})
	gob.Register(config.Config{})
}
