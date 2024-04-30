package main

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"
	"time"

	"github.com/0xAX/notificator"
)

func get_image(image string) string {
	usr, err := user.Current()

	if err != nil {
        log.Fatal(err)
    }

	var file = filepath.Join(usr.HomeDir, ".config", "eye-break-go", "assets", image)

	fmt.Println(file)

	return file
}

func user_notify(title string, description string, image string) {
	var notify = notificator.New(notificator.Options{
		DefaultIcon: get_image("icon.ico"),
		AppName: "Eye Break",
	})

	var image_path = get_image(image)

	notify.Push(title, description, image_path, notificator.UR_NORMAL)
}

func main() {

	for {
		user_notify(
			"Time to take an eye break!",
			"Look at least 20 feet away for 20 seconds",
			"away.gif",
		)

		time.Sleep(time.Second * 20)

		user_notify(
			"You can continue now :)",
			"20 seconds up you may continue.",
			"can_look.png",
		)

		time.Sleep(time.Minute * 25)
	}
}