package hash

import "golang.org/x/crypto/bcrypt"


func Hash(value string, cost ...int) (string, error) {
	c := bcrypt.DefaultCost

	if len(cost) > 0 {
		if cost[0] >= bcrypt.MinCost && cost[0] <= bcrypt.MaxCost {
			c = cost[0]
		}
	}
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(value),
		c,
	)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func Compare(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(plain),
	)

	return err == nil
}