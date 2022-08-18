package middleware

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	token,_ := CreateToken(1)
	 ParseToken(token)

}
