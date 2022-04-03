## Install Fonts
```bash
sudo su
```

- #### ***Hack Nerd Font***
	---
	Download the font ***Hack Nerd Font*** from the page `https://www.nerdfonts.com/font-downloads` and install with the following commands
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/Hack
	cd !$
	mv /home/alejo/Downloads/Firefox/Hack.zip .
	unzip Hack.zip
	rm Hack.zip
	```

- #### ***Meslo Nerd Font***
	---
	Download the font ***Meslo Nerd Font*** from the page `https://www.nerdfonts.com/font-downloads` and install with the following commands
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/Meslo
	cd !$
	mv /home/alejo/Downloads/Firefox/Meslo.zip .
	unzip Meslo.zip
	rm Meslo.zip
	```

Reload fonts
```bash
fc-cache -vf
```