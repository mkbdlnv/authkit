package hasher

import "golang.org/x/crypto/bcrypt"

type BcryptHahser struct {
	cost int
}

func NewBcryptHasher() *BcryptHahser {
	return &BcryptHahser{cost: bcrypt.DefaultCost}
}

func NewBcryptHasherWithCost(cost int) *BcryptHahser {
	return &BcryptHahser{cost: cost}
}

func (h *BcryptHahser) Hash(password string) (string, error){
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(hashed), err
}

func (h *BcryptHahser) Compare(hashedPassword, plainPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword));
}