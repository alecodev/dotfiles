## Install SSH

```bash
sudo pacman -S openssh
mkdir -p ~/.ssh
chmod 700 ~/.ssh
echo "# host-specific options\n" >> ~/.ssh/config
chmod 600 ~/.ssh/config

# Generate a new ssh key
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
mkdir -p ~/.config/systemd/user
```

Download file `~/.config/systemd/user/ssh-agent.service` with the following command
```bash
wget --quiet https://raw.githubusercontent.com/alecodev/dotfiles/main/.config/systemd/user/ssh-agent.service --output-document=~/.config/systemd/user/ssh-agent.service
```

Edit file `~/.bashrc` and `~/.zshrc` with [text editor][1] and add the following lines
```zsh
# Service ssh-agent
export SSH_AUTH_SOCK="$XDG_RUNTIME_DIR/ssh-agent.socket"
```

Enable the service
```zsh
systemctl --user enable ssh-agent
systemctl --user start ssh-agent
ssh-add ~/.ssh/id_rsa
```

[1]:../../README.md#text-editor