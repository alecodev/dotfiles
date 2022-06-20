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

- #### ***JetBrainsMono Nerd Font***
	---
	Download the font ***JetBrainsMono Nerd Font*** from the page `https://www.nerdfonts.com/font-downloads` and install with the following commands
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/JetBrainsMono
	cd !$
	mv /home/alejo/Downloads/Firefox/JetBrainsMono.zip .
	unzip JetBrainsMono.zip
	rm JetBrainsMono.zip
	```

- #### ***Rofi***
	---
	Download the font ***Rofi*** from the page `https://github.com/adi1090x/rofi/tree/master/fonts` and install with the following commands
	```bash
	mkdir -p /usr/local/share/fonts/rofi
	cd !$
	mv -r /home/alejo/Downloads/Firefox/* .
	```

- #### ***Material Design Font***
	---
	Download the font ***Material Design Font*** from the page `https://github.com/Templarian/MaterialDesign-Font` and install with the following commands
	```bash
	mkdir -p /usr/local/share/fonts/MaterialDesign-Font
	cd !$
	wget https://github.com/Templarian/MaterialDesign-Font/raw/master/MaterialDesignIconsDesktop.ttf
	```

Reload fonts
```bash
fc-cache -vf
```