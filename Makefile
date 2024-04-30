build:
	go build -o eye-break-go ./src/eye-break.go

install: install-assets
	cp eye-break-go ~/.local/bin
	cp ./eye-break-go.service ~/.config/systemd/user
	systemctl --user daemon-reload

install-assets:
	mkdir -p $(HOME)/.config/eye-break-go
	cp ./assets $(HOME)/.config/eye-break-go -r

uninstall: remove-assets
	systemctl --user stop eye-break-go
	systemctl --user disable eye-break-go
	rm ~/.local/bin/eye-break-go
	rm ~/.config/systemd/user/eye-break-go.service
	systemctl --user daemon-reload

remove-assets:
	sudo rm -r $(HOME)/.config/eye-break-go
