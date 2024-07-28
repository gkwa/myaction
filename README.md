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

# Run gmailctl after successful validation
myaction email-from-cli --run-gmailctl user@example.com

# Show version
myaction version

# Get help
myaction --help
