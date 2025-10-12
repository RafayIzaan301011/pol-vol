package security

type Encryption interface {
	Encrypt(data string) (string, error)
	Decrypt(data string) (string, error)
}
