package sys

import (
	"os/exec"
	"path/filepath"

	"github.com/golang/glog"
)

// Exec executes an executable on a new terminal. path can be absolute or $PATH combined with path
func Exec(path string) (*exec.Cmd, error) {
	if path, err := exec.LookPath(path); err != nil {
		glog.Error("can't find " + path)
		return nil, err
	}
	cmd := exec.Command("cmd", "/C", "start", "/D"+filepath.Dir(path), path)
	if err := cmd.Start(); err != nil {
		glog.Error("can't exec " + path + ": " + err.Error())
		return nil, err
	}
	return cmd, nil
}

func Kill(pid string) error {
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", pid)
	//kill.Stderr = os.Stderr
	//kill.Stdout = os.Stdout
	return kill.Run()
}

func KillByName(name string) error {
	kill := exec.Command("TASKKILL", "/T", "/F", "/IM", name)
	//kill.Stderr = os.Stderr
	//kill.Stdout = os.Stdout
	return kill.Run()
}
