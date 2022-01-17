package erros

import "thiagofelipe.com.br/sistema-faculdade-backend/erros"

const (
	decodificarBase64 = iota + 1
	decodificarHashInválido
	decodificarArgon2id
	decodificarArgon2idVersão
	verificarSenhaHash
	desencriptarAES
	desencriptarAESNonceSize
)

var criarErroPadrão = erros.NovoPadrãoFunc("ENTIDADES") //nolint: gochecknoglobals

var (
	ErroDecodificarBase64 = criarErroPadrão(
		"Erro ao decodificar a palavra na base64",
		decodificarBase64,
	)
	ErroDecodificarHashInválido = criarErroPadrão(
		"O hash informado é inválido para a decodificação",
		decodificarHashInválido,
	)
	ErroDecodificarArgon2id = criarErroPadrão(
		"Erro ao tentar decodificar o hash com argon2id",
		decodificarArgon2id,
	)
	ErroDecodificarArgon2idVersão = criarErroPadrão(
		"Versão do argon2id inválida",
		decodificarArgon2idVersão,
	)
	ErroVerificarSenhaHash = criarErroPadrão(
		"Erro ao tentar verificar a senha com o hash",
		verificarSenhaHash,
	)
	ErroDesencriptarAES = criarErroPadrão(
		"Erro ao tentar desencriptar a senha com o algoritmo AES",
		desencriptarAES,
	)
	ErroDesencriptarAESNonceSize = criarErroPadrão(
		"A senha cifrada é menor que o nonce",
		desencriptarAESNonceSize,
	)
)
