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

- #### ***Material Design Font***
	---
	Descargue la fuente ***Material Design Font*** de la página `https://github.com/Templarian/MaterialDesign-Font` e instálela con los siguientes comandos
	```bash
	mkdir -p /usr/local/share/fonts/MaterialDesign-Font
	cd !$
	wget https://github.com/Templarian/MaterialDesign-Font/raw/master/MaterialDesignIconsDesktop.ttf
	```

Recargue las fuentes
```bash
fc-cache -vf
```