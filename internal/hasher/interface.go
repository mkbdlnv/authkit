package hasher

type Hasher interface{
	Hash(password string) (string, error)
	Compare(hashedPassword, plainPassword string) (error)
}