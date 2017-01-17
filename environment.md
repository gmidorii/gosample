# Golang environment establishment

### HowTo
##### 1. Install golang
```
% brew install go
```
search `GOROOT`
```
% go env GOROOT
```
##### 2. Setting $GOPATH  
Setting $GOPATH is necessary.  
you develop in $GOPATH/.  
But $GOPATH is everywhere, is ok.  
write .bashrc or .zshrc or etc.  
```
export GOPATH=$HOME/dev
export PATH=$PATH:$GOPATH/bin
```

##### 3. Install ghq
```
% brew install ghq
```
set directory.
```
% git config --global ghq.root $GOPATH/src
```

##### 4. Install peco
```
brew install peco
```
write .bashrc or .zshrc or etc.  
```
bindkey '^]' peco-src

function peco-src() {
  local src=$(ghq list --full-path | peco --query "$LBUFFER")
  if [ -n "$src" ]; then
    BUFFER="cd $src"
    zle accept-line
  fi
  zle -R -c
}
zle -N peco-src
```
short cut `Ctr + ]`
