## Instalación de SSH

```bash
sudo pacman -S openssh
mkdir -p ~/.ssh
chmod 700 ~/.ssh
echo "# host-specific options\n" >> ~/.ssh/config
chmod 600 ~/.ssh/config

# Generar una nueva clave ssh
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
mkdir -p ~/.config/systemd/user
```

Descargue el archivo `~/.config/systemd/user/ssh-agent.service` con el siguiente comando
```bash
wget --quiet https://raw.githubusercontent.com/alecodev/dotfiles/main/.config/systemd/user/ssh-agent.service --output-document=~/.config/systemd/user/ssh-agent.service
```

Edite el archivo `~/.bashrc` y `~/.zshrc` con [editor de texto][1] y agregue las siguientes líneas
```zsh
# Service ssh-agent
export SSH_AUTH_SOCK="$XDG_RUNTIME_DIR/ssh-agent.socket"
```

Habilite el servicio
```zsh
systemctl --user enable ssh-agent
systemctl --user start ssh-agent
ssh-add ~/.ssh/id_rsa
```

[1]:../../README.es.md#editor-de-texto