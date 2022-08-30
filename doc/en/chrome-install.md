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

>Or install with the following command
>```zsh
>yay -S google-chrome
>```

Create and change the download directory in Chrome settings to `~/Downloads/Chrome`
```bash
mkdir -p ~/Downloads/Chrome
```

Execute the following commands
```zsh
echo "--force-dark-mode" >> ~/.config/chrome-flags.conf
sudo pacman -S noto-fonts-emoji
gsettings set org.gnome.desktop.interface gtk-theme 'Adwaita-dark'
gsettings set org.gnome.desktop.interface color-scheme 'prefer-dark'
```

[1]:../../README.md#text-editor