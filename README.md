# Dotfiles


***Languages***
- [🇪🇸 - Español](./README.es.md)
- **🇺🇸 - English**
---


#### ***Linux Distribution***
- [Arch](doc/en/arch-install.md)


#### ***Text Editor***
- neovim
>but feel free to use your preferred text editor (vim, nano, ...)


---
## Install Window Manager

Login with the root user and run the following commands
```bash
pacman -Syu
pacman -S --needed gcc make git base-devel \
numlockx \
nmap wget curl \
xclip \
p7zip unzip \
zsh tmux \
dunst ranger htop locate \
flameshot \
libsecret gnome-keyring gnome-themes-extra \
xorg-server xorg-xev \
bspwm sxhkd \
alacritty kitty \
rofi polybar picom \
lightdm lightdm-gtk-greeter \
bat lsd fzf jq \
feh \
neofetch \
pipewire pipewire-pulse \
udisks2 udiskie \
openssh
```

Set keyboard layout
```bash
localectl set-x11-keymap es
```

Edit the file `/etc/lightdm/lightdm.conf` with [text editor][1] and modify the following lines
```diff
-#greeter-session=example-gtk-gnome
+greeter-session=lightdm-gtk-greeter
...
-#display-setup-script=
+display-setup-script=/usr/bin/setxkbmap -layout latam,es
...
-#greeter-setup-script=
+greeter-setup-script=/usr/bin/numlockx on
```

Activate the lightdm service
```bash
systemctl enable lightdm
```

Change user
```bash
su alejo
```

Create or edit the file `~/.xprofile` with [text editor][1] and set the following lines
>**In case of running in a virtual machine like VirtualBox add the following line**
>
>VBoxClient-all &
```text
dbus-update-activation-environment --systemd DISPLAY &
sxhkd &
exec bspwm
```

Create the directories and set the configuration
```bash
# Create directories
mkdir -p ~/{.config,.ssh,Images}

# Clone repo
git clone https://github.com/alecodev/dotfiles.git
cd dotfiles

# Add permissions to files
sudo chmod +x $PWD/.config/bspwm/bspwmrc
sudo chmod +x $PWD/.config/bspwm/scripts/{bspwm_layout,bspwm_resize,bspwm_smart_move}

# Create symbolic links of the files
ln -sf $PWD/.config/{alacritty,bspwm,gtk-3.0,kitty,picom,polybar,sxhkd,systemd} ~/.config/
ln -sf $PWD/{.aliases,.bashrc,.p10k.zsh,.xprofile,.zshrc} ~/

# Change default shell per user
sudo usermod --shell /usr/bin/zsh $(whoami)
sudo usermod --shell /usr/bin/zsh root

# Install Oh My Zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# Add plugins and themes
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/themes/powerlevel10k

# Install Oh My Zsh for root user
sudo sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# Create symbolic links of the files to the root user
sudo ln -sf ~/.oh-my-zsh/custom/plugins/{zsh-autosuggestions,zsh-syntax-highlighting} /root/.oh-my-zsh/custom/plugins/
sudo ln -sf ~/.oh-my-zsh/custom/themes/powerlevel10k /root/.oh-my-zsh/custom/themes/
sudo ln -sf ~/{.aliases,.bashrc,.p10k.zsh,.zshrc} /root/

# Set Wallpaper
wget --quiet https://wallpaperaccess.com/full/2098223.png --output-document=~/Images/wallpaper.png

# Change power button behavior
sudo sed -i 's/#HandlePowerKey=poweroff/HandlePowerKey=ignore/' /etc/systemd/logind.conf
sudo systemctl restart systemd-logind
```

Reboot and log in with the other user
```bash
reboot
```

Ready now you can log in with the other user and use bspwm by pressing `Super + Enter`

---
## Setting up the work environment

### Programs & Applications

- [Yay](doc/en/yay-install.md)
- [Git](doc/en/git-install.md)
- [Bluetooth](doc/en/bluetooth-install.md)
- [SSH](doc/en/ssh-install.md)
- [GPG](doc/en/gpg-install.md)
- [Fonts](doc/en/fonts-install.md)
- [Visual Studio Code](doc/en/vscode-install.md)
- [Firefox](doc/en/firefox-install.md)
- [Google Chrome](doc/en/chrome-install.md)
- [Docker](doc/en/docker-install.md)
- [DBeaver](doc/en/dbeaver-install.md)
- [DroidCam](doc/en/droidcam-install.md)

[1]:#text-editor