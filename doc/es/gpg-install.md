## Instalación de GPG

```bash
pacman -S gnupg

# Generando una nueva clave GPG
gpg --full-generate-key
git config --global commit.gpgsign true
git config --global credential.helper /usr/lib/git-core/git-credential-libsecret
```

La clave GPG se agrega a git
```zsh
gpg --list-secret-keys --keyid-format=long
git config --global user.signingkey YOURKEY

# Generar clave pública para Github
gpg --armor --export YOURKEY
```

Edite el archivo `~/.bashrc` con [editor de texto][1] y agregue las siguientes líneas
```zsh
#  GPG key
export GPG_TTY=$(tty)
```

Edite el archivo `~/.zshrc` con [editor de texto][1] y agregue las siguientes líneas
```zsh
#  GPG key
export GPG_TTY=$TTY
```

[1]:../../README.es.md#editor-de-texto