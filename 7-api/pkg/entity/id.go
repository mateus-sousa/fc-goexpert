package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	// O ID() é como se fosse um casting para o tipo que criamos, assim como fazemso string([]byte), nesse caso parece ser uma conversão redundante
	// como no curso esta usando, vamos manter
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
