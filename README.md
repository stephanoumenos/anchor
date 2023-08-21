# anchor

Anchor is a CLI designed to streamline your navigation in the terminal. The main purpose of this tool is to allow you to go to your saved directory anchor automatically whenever you open your terminal. This is particularly helpful for project management, as you can easily navigate to different project directories saved as anchors.

## Key Features

- **Automatically Navigate to Default Anchor:** Open your terminal to your saved anchor automatically, every time.
- **Project Management:** Manage your various project directories by saving them as named anchors.
- **Easy Navigation:** Use simple commands to navigate between your anchors, save new ones, and more.

## Usage

Below are the available commands for the Anchor CLI.

### Set the default directory

Set the current directory or a named anchor as the default directory to automatically navigate to.

```bash
anchor down [anchor_name]
```

### Unset Default Directory

Unset the current default directory.

```bash
anchor up
```

### Create a saved Anchor

Save the current directory as a named anchor for easy project navigation.

```bash
anchor save [anchor_name]
```

### Navigate to an anchor

Go to your saved anchor whenever you open your terminal, or manually navigate to one of your saved named anchors.

```bash
anchor go [anchor_name]
```

Use the `-f` flag to enable fuzzy finding mode, allowing you to easily search and select an anchor from your saved anchors:

```bash
anchor go -f
```

### Delete a saved anchor

Delete a saved anchor directory.

```bash
anchor remove [anchor_name]
```

### Get Current Anchor Path

Get the path of the current anchor.

```bash
anchor get
```

### List Saved Anchors

List the current saved project directories.

```bash
anchor list
```

### Generate Completion Script

Generate the completion script for different shell environments.

```bash
anchor completion [bash|zsh|fish|powershell]
```

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
    elif [[ -n $$2 ]]; then\
      echo "⛵️ No saved anchor named '\''$$2'\'' found.";\
    else\
      echo "⛵️ No default anchor set. Use '\''anchor down'\'' to set your current directory as the default.";\
    fi;\
  else\
    command anchor "$$@";\
  fi\
}\
\
if [[ $$PWD == $$HOME ]]; then\
  eval "cd $$(anchor get)";\
fi\
source <(anchor completion zsh)' >> ~/.zshrc
```

#### For bash:

Run this one-liner to integrate `anchor` into your bash shell:

```bash
echo 'anchor() {\
  if [[ $$1 == "go" ]]; then\
    local anchor_path="$$(command anchor go $$2 2>/dev/null)";\
    if [[ -n $$anchor_path ]]; then\
      cd "$$anchor_path";\
    elif [[ -n $$2 ]]; then\
      echo "⛵️ No saved anchor named '\''$$2'\'' found.";\
    else\
      echo "⛵️ No default anchor set. Use '\''anchor down'\'' to set your current directory as the default.";\
    fi;\
  else\
    command anchor "$$@";\
  fi\
}\
\
if [[ $$PWD == $$HOME ]]; then\
  eval "cd $$(anchor get)";\
fi\
source <(anchor completion bash)' >> ~/.bashrc
```

#### For fish:

Run this one-liner to integrate `anchor` into your fish shell:

```bash
echo 'function anchor\
  if test $$argv[1] = "go";\
    set anchor_path (command anchor go $$argv[2] 2>/dev/null);\
    if test -n $$anchor_path;\
      cd $$anchor_path;\
    else if test -n $$argv[2];\
      echo "⛵️ No saved anchor named '\''$$argv[2]'\'' found.";\
    else;\
      echo "⛵️ No default anchor set. Use '\''anchor down'\'' to set your current directory as the default.";\
    end;\
  else;\
    command anchor $$argv;\
  end;\
end\
\
if test $$PWD = $$HOME;\
  eval "cd (anchor get)";\
end\
anchor completion fish | source' >> ~/.config/fish/config.fish
```

### 5. Apply Changes

Now, restart your shell or run one of the following to apply the changes:

- For zsh: `source ~/.zshrc`
- For bash: `source ~/.bashrc`
- For fish: `source ~/.config/fish/config.fish`

Enjoy using `anchor` to enhance your navigation between directories!
