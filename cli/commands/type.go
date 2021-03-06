package commands

type Command func(args []string)

var Aliases = map[string]Command{
	"help":       help,
	"auth":       auth,
	"clip":       clip,
	"exit":       exit,
	"config":     configCmd,
	"connect":    connect,
	"disconnect": disconnect,
	"rcon":       rcon,
}

var HelpList = map[string]string{
	"help": `Displays this list.`,
	"auth": `auth user:
Attempts to log in with current user credentials.

auth admin:
Attempts to log in with current admin credentials.

auth pass [admin/user] [password]:
Updates the details on file for the admin or user.

auth check:
Checks what the client is currently authorized for.`,
	"config": `config server [ip:port]:
Sets the server IP and port.

config buffer [size]:
Sets the buffer size.

config save:
Saves the config.

config reload:
Reloads the config from the file. Does not disconnect or de-auth.`,
	"clip": `clip force [push/refresh]:
Forcefully pushes or refreshes the clipboard to/from the server.

clip recheck:
Forcefully checks the clipboard for an update.

clip empty:
Empties the clipboard.`,
	"exit":       `Shuts down the server.`,
	"connect":    `Connects you to the server.`,
	"disconnect": `Disconnects you from the server.`,
	"rcon": `rcon [command]:
Issues remote console command to server.

ex. 
rcon config buffer 256
sets the server buffer size to 256 bytes.`,
	"security": `security tls:
	displays TLS configuration

	security tls toggle:
		Toggles TLS on/off
	
	security tls toggleVerify:
		Toggles TLS verification on/off. Use only for testing.`,
}
