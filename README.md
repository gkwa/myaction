# myaction
Purpose: A tool to manage Gmail filters using the gmailctl configuration.

## Example Usage
```bash
# Add a new email filter from clipboard
myaction email-from-clipboard

# Add a new email filter from command line
myaction email-from-cli user@example.com

# Add a new domain filter from command line
myaction email-from-cli @example.com

# Use a custom gmailctl path
myaction email-from-cli --gmailctl-path /path/to/gmailctl user@example.com

# Enable verbose logging
myaction -v email-from-clipboard

# Use JSON logging format
myaction --log-format json email-from-cli user@example.com

# Show version
myaction version

# Get help
myaction --help
```

## Install myaction
On macOS/Linux:
```bash
brew install gkwa/homebrew-tools/myaction
```

On Windows:
```powershell
TBD
```

## Configuration
By default, myaction looks for the gmailctl configuration file at `$HOME/.gmailctl/config.jsonnet`. 
It uses `/usr/local/bin/gmailctl` as the default path for the gmailctl executable.

You can override these defaults using command-line flags or by setting up a configuration file at `$HOME/.myaction.yaml`.

Example `.myaction.yaml`:
```yaml
gmailctl-path: /custom/path/to/gmailctl
verbose: true
log-format: json
```
