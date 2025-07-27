#!/bin/bash

# SSH Connection Diagnostic Script
# This script helps diagnose SSH connection issues for GitHub Actions deployment

set -e

echo "🔍 SSH Connection Diagnostic Tool"
echo "=================================="
echo ""

# Check if required parameters are provided
if [ $# -lt 3 ]; then
    echo "Usage: $0 <host> <username> <private_key_path>"
    echo ""
    echo "Example:"
    echo "  $0 your-server.com ubuntu ~/.ssh/id_rsa"
    echo "  $0 192.168.1.100 ec2-user ~/.ssh/deploy_key"
    exit 1
fi

HOST=$1
USERNAME=$2
PRIVATE_KEY_PATH=$3

echo "Testing connection to: $USERNAME@$HOST"
echo "Using private key: $PRIVATE_KEY_PATH"
echo ""

# Check if private key file exists
if [ ! -f "$PRIVATE_KEY_PATH" ]; then
    echo "❌ Private key file not found: $PRIVATE_KEY_PATH"
    exit 1
fi

# Check private key format
echo "🔑 Checking private key format..."
if head -n 1 "$PRIVATE_KEY_PATH" | grep -q "BEGIN.*PRIVATE KEY"; then
    echo "✅ Private key format looks correct"
    head -n 1 "$PRIVATE_KEY_PATH"
else
    echo "❌ Private key format may be incorrect"
    echo "First line of key file:"
    head -n 1 "$PRIVATE_KEY_PATH"
    echo ""
    echo "Expected format should start with something like:"
    echo "-----BEGIN OPENSSH PRIVATE KEY-----"
    echo "or"
    echo "-----BEGIN RSA PRIVATE KEY-----"
fi

# Check key permissions
echo ""
echo "🔒 Checking key permissions..."
KEY_PERMS=$(stat -c %a "$PRIVATE_KEY_PATH")
if [ "$KEY_PERMS" = "600" ]; then
    echo "✅ Key permissions are correct (600)"
else
    echo "⚠️  Key permissions are $KEY_PERMS (should be 600)"
    echo "Fix with: chmod 600 $PRIVATE_KEY_PATH"
fi

# Test basic connectivity
echo ""
echo "🌐 Testing basic connectivity..."
if ping -c 1 -W 5 "$HOST" > /dev/null 2>&1; then
    echo "✅ Host is reachable"
else
    echo "❌ Host is not reachable"
    echo "Check if the hostname/IP is correct and accessible"
fi

# Test SSH connection with verbose output
echo ""
echo "🔌 Testing SSH connection..."
echo "Running: ssh -o ConnectTimeout=10 -o BatchMode=yes -i $PRIVATE_KEY_PATH $USERNAME@$HOST 'echo \"Connection successful\"'"
echo ""

if timeout 30 ssh -o ConnectTimeout=10 -o BatchMode=yes -i "$PRIVATE_KEY_PATH" "$USERNAME@$HOST" 'echo "Connection successful"' 2>/dev/null; then
    echo "✅ SSH connection successful!"
    echo ""
    echo "Your SSH configuration is working correctly."
    echo "The issue might be with how the key is stored in GitHub Secrets."
else
    echo "❌ SSH connection failed"
    echo ""
    echo "🔍 Running verbose SSH test..."
    timeout 30 ssh -vvv -o ConnectTimeout=10 -o BatchMode=yes -i "$PRIVATE_KEY_PATH" "$USERNAME@$HOST" 'echo "Connection test"' 2>&1 | head -n 50
fi

echo ""
echo "📝 Next steps:"
echo "1. If the connection worked locally, the issue is likely with GitHub Secrets"
echo "2. Make sure you're copying the ENTIRE private key to GitHub Secrets"
echo "3. Check that the username in GitHub Secrets matches what you tested here"
echo "4. Verify the public key is in ~/.ssh/authorized_keys on the server"
echo ""
echo "To check the public key on the server:"
echo "  ssh -i $PRIVATE_KEY_PATH $USERNAME@$HOST 'cat ~/.ssh/authorized_keys'"
echo ""
echo "To generate the public key from this private key:"
echo "  ssh-keygen -y -f $PRIVATE_KEY_PATH" 