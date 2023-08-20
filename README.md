# anchor

# Installation

### Pre-requisites

Before proceeding with the installation, please ensure that you have the Go programming language installed on your system. Having Go properly set up is essential for compiling the anchor binary from source. If you don't have it installed, you can follow the official installation guide to set up Go in your environment.

### 1. Clone the Repository

```bash
git clone git@github.com:stephanoumenos/anchor.git
```

### 2. Compile the Binary

```bash
cd anchor && go build
```

### 3. Copy the Binary

```bash
sudo cp anchor /usr/local/bin/
```

### 4. Shell Integration

#### For zsh:

Run this one-liner to integrate `anchor` into your zsh shell:

```bash
echo 'anchor() {\
  if [[ $$1 == "go" ]]; then\
    local anchor_path="$$(command anchor go $$2 2>/dev/null)";\
    if [[ -n $$anchor_path ]]; then\
      cd "$$anchor_path";\
    else\
      echo "⛵️ No default anchor set. Use 'anchor down' to set your current directory as the default.";\
    fi;\
  else\
    command anchor "$$@";\
  fi\
}\
\
if [[ $$PWD == $$HOME ]]; then\
  eval "cd $$(anchor get)";\
fi' >> ~/.zshrc
```

#### For bash:

Run this one-liner to integrate `anchor` into your bash shell:

```bash
echo 'anchor() {\
  if [[ $$1 == "go" ]]; then\
    local anchor_path="$$(command anchor go $$2 2>/dev/null)";\
    if [[ -n $$anchor_path ]]; then\
      cd "$$anchor_path";\
    else\
      echo "⛵️ No default anchor set. Use 'anchor down' to set your current directory as the default.";\
    fi;\
  else\
    command anchor "$$@";\
  fi\
}\
\
if [[ $$PWD == $$HOME ]]; then\
  eval "cd $$(anchor get)";\
fi' >> ~/.bashrc
```

#### For fish:

Run this one-liner to integrate `anchor` into your fish shell:

```bash
echo 'function anchor\
  if test $$argv[1] = "go";\
    set anchor_path (command anchor go $$argv[2] 2>/dev/null);\
    if test -n $$anchor_path;\
      cd $$anchor_path;\
    else;\
      echo "⛵️ No default anchor set. Use 'anchor down' to set your current directory as the default.";\
    end;\
  else;\
    command anchor $$argv;\
  end;\
end\
\
if test $$PWD = $$HOME;\
  eval "cd (anchor get)";\
end' >> ~/.config/fish/config.fish
```

### 5. Apply Changes

Now, restart your shell or run one of the following to apply the changes:

- For zsh: `source ~/.zshrc`
- For bash: `source ~/.bashrc`
- For fish: `source ~/.config/fish/config.fish`

Enjoy using `anchor` to enhance your navigation between directories!
