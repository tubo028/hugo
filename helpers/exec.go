package helpers

import (
	"bytes"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"os/exec"
)

func Exec(command, argsStr string) string {
	allowed := false
	for _, elem := range viper.GetStringSlice("ExecWhitelist") {
		if elem == command {
			allowed = true
			break
		}
	}
	if !allowed {
		jww.ERROR.Println("Executing `" + command + "` is not allowed")
		return ""
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	var cmd *exec.Cmd
	if argsStr == "" {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command(command, argsStr)
	}
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		jww.ERROR.Print("Error while executing `"+command+"`: ", stderr.String())
		return ""
	}
	return out.String()
}
