package external_command

type Command struct {
	Cmd    byte
	Params []string
}

var Commands = map[byte]func(args []string){
	0: SetClip,
	1: Success,
	2: Failure,
} //Pointers because some commands can elevate privileges.
//For now, clients will have just one command available to the server: 0 updateClip

/*
Command structure:
[CMD, 10, ...., 10, ...., 0]
Every ASCII LF (\n) character is a new argument.
Every command ends with a NULL byte.
*/

/*
Server commands:
0: becomeAdmin
1: rcon
2: getClip
3: setClip
4: login

Client commands:
0 setClip
1 success
2 failure
*/
func ParseCmd(buffer []byte) Command {
	cmd := Command{Cmd: buffer[0], Params: make([]string, 0)}

	start := -1
	k := -1
	var v byte
	for k, v = range buffer {
		if k != 0 && v == 0 {
			if start != -1 {
				cmd.Params = append(cmd.Params, string(buffer[start:k]))
			}
			break
		}

		if k != 0 && v == 10 { //LF ASCII
			if start != -1 {
				cmd.Params = append(cmd.Params, string(buffer[start:k]))
			}
			start = k + 1
		}
	}

	return cmd
}
