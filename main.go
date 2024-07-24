package main

//#include<conio.h>
import "C"
import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"golang.org/x/sys/windows"
)

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, nil)))

	_, err := windows.CreateMutex(nil, true, windows.StringToUTF16Ptr("ROBLOX_singletonMutex"))

	if err != nil {
		if errors.Is(err, windows.ERROR_ALREADY_EXISTS) {
			slog.Error("make sure roblox isnt open before running this.")
		} else {
			slog.Error("mutex creation error:", tint.Err(err))
		}

		time.Sleep(2 * time.Second)
		os.Exit(1)
	}

	slog.Info("you can now launch multiple instances of roblox! press any key to exit.")
	windows.MessageBox(0, windows.StringToUTF16Ptr("You can now launch multiple instances of roblox. Press any key in the terminal to exit!"), windows.StringToUTF16Ptr("Mutex Unlocked"), 0)

	C.getch()
}
