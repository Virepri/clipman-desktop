package internal_command

type Command struct {
	Cmd CommandID
	Params []string
}

type CommandID uint8

const (
	CONNECT CommandID = iota //resets the connection.

	AUTH_USER //Requires an active connection.
	AUTH_ADMIN

	PUSH_CLIP //Requires authentication.
	REFRESH_CLIP

	RCON_COMMAND
)