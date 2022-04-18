### Install Google Chrome

```zsh
cd !$
git clone https://aur.archlinux.org/google-chrome.git
cd google-chrome
makepkg -si
cd ..
rm -r google-chrome
sudo su
```

>O instala con el siguiente comando
>```zsh
>yay -S google-chrome
>```

Cree y cambie el directorio de descarga en la configuración de Chrome a `~/Downloads/Chrome`
```bash
mkdir -p ~/Downloads/Chrome
```

Ejecuta los siguientes comandos
```zsh
echo "--force-dark-mode" >> ~/.config/chrome-flags.conf
sudo pacman -S noto-fonts-emoji
gsettings set org.gnome.desktop.interface gtk-theme "Adwaita-dark"
```

[1]:../../README.es.md#editor-de-texto