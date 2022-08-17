## Instalación de fuentes
```bash
sudo su
```

- #### ***Hack Nerd Font***
	---
	Descargue la fuente ***Hack Nerd Font*** de la página `https://www.nerdfonts.com/font-downloads` e instálela con los siguientes comandos
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/Hack
	cd !$
	mv /home/alejo/Downloads/Firefox/Hack.zip .
	unzip Hack.zip
	rm Hack.zip
	```

- #### ***Meslo Nerd Font***
	---
	Descargue la fuente ***Meslo Nerd Font*** de la página `https://www.nerdfonts.com/font-downloads` e instálela con los siguientes comandos
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/Meslo
	cd !$
	mv /home/alejo/Downloads/Firefox/Meslo.zip .
	unzip Meslo.zip
	rm Meslo.zip
	```

- #### ***JetBrainsMono Nerd Font***
	---
	Descargue la fuente ***JetBrainsMono Nerd Font*** de la página `https://www.nerdfonts.com/font-downloads` e instálela con los siguientes comandos
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/JetBrainsMono
	cd !$
	mv /home/alejo/Downloads/Firefox/JetBrainsMono.zip .
	unzip JetBrainsMono.zip
	rm JetBrainsMono.zip
	```

- #### ***Rofi***
	---
	Descargue la fuente ***Rofi*** de la página `https://github.com/adi1090x/rofi/tree/master/fonts` e instálela con los siguientes comandos
	```bash
	mkdir -p /usr/local/share/fonts/Rofi
	cd !$
	mv -r /home/alejo/Downloads/Firefox/* .
	```

- #### ***Material Design Font***
	---
	Descargue la fuente ***Material Design Font*** de la página `https://github.com/Templarian/MaterialDesign-Font` e instálela con los siguientes comandos
	```bash
	mkdir -p /usr/local/share/fonts/MaterialDesign
	cd !$
	wget --quiet https://github.com/Templarian/MaterialDesign-Font/raw/master/MaterialDesignIconsDesktop.ttf
	```

Recargue las fuentes
```bash
fc-cache -vf
```