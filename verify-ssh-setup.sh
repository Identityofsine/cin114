#!/bin/bash

# SSH Setup Verification Script
# Run this ON YOUR SERVER to verify SSH configuration

echo "=== SSH Setup Verification ==="
echo ""

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "⚠️  Not running as root. Some checks may fail."
    echo "   Run with: sudo ./verify-ssh-setup.sh"
    echo ""
fi

# 1. Check if .ssh directory exists and has correct permissions
echo "1. Checking SSH directory..."
if [ -d "/root/.ssh" ]; then
    echo "✅ /root/.ssh directory exists"
    ls -la /root/.ssh/
    echo ""
else
    echo "❌ /root/.ssh directory not found"
    echo "   Creating directory..."
    mkdir -p /root/.ssh
    chmod 700 /root/.ssh
    echo "✅ Created /root/.ssh with correct permissions"
    echo ""
fi

# 2. Check authorized_keys file
echo "2. Checking authorized_keys file..."
if [ -f "/root/.ssh/authorized_keys" ]; then
    echo "✅ authorized_keys file exists"
    echo "   File permissions: $(ls -l /root/.ssh/authorized_keys)"
    echo "   Number of keys: $(wc -l < /root/.ssh/authorized_keys)"
    echo ""
    echo "   Keys in file:"
    cat /root/.ssh/authorized_keys | while read line; do
        if [[ $line == ssh-* ]]; then
            echo "   - ${line:0:50}..."
        fi
    done
    echo ""
else
    echo "❌ authorized_keys file not found"
    echo "   You need to add your public key to: /root/.ssh/authorized_keys"
    echo ""
fi

# 3. Check SSH daemon configuration
echo "3. Checking SSH daemon configuration..."
if [ -f "/etc/ssh/sshd_config" ]; then
    echo "✅ SSH daemon config found"
    echo ""
    echo "   PubkeyAuthentication: $(grep -E '^PubkeyAuthentication' /etc/ssh/sshd_config || echo 'default (yes)')"
    echo "   PasswordAuthentication: $(grep -E '^PasswordAuthentication' /etc/ssh/sshd_config || echo 'default (yes)')"
    echo "   PermitRootLogin: $(grep -E '^PermitRootLogin' /etc/ssh/sshd_config || echo 'default (yes)')"
    echo "   AuthorizedKeysFile: $(grep -E '^AuthorizedKeysFile' /etc/ssh/sshd_config || echo 'default (.ssh/authorized_keys)')"
    echo ""
else
    echo "❌ SSH daemon config not found"
    echo ""
fi

# 4. Check SSH service status
echo "4. Checking SSH service status..."
if systemctl is-active --quiet sshd; then
    echo "✅ SSH service is running"
elif systemctl is-active --quiet ssh; then
    echo "✅ SSH service is running"
else
    echo "❌ SSH service is not running"
    echo "   Try: sudo systemctl start sshd"
fi
echo ""

# 5. Check firewall
echo "5. Checking firewall..."
if command -v ufw &> /dev/null; then
    echo "   UFW status: $(ufw status | head -1)"
    if ufw status | grep -q "22/tcp"; then
        echo "✅ SSH port 22 is allowed in UFW"
    else
        echo "⚠️  SSH port 22 might be blocked by UFW"
    fi
elif command -v iptables &> /dev/null; then
    echo "   Checking iptables for SSH rules..."
    if iptables -L | grep -q "ssh\|22"; then
        echo "✅ SSH rules found in iptables"
    else
        echo "⚠️  No SSH rules found in iptables"
    fi
else
    echo "   No firewall configuration found"
fi
echo ""

# 6. Show recent SSH connection attempts
echo "6. Recent SSH connection attempts:"
if [ -f "/var/log/auth.log" ]; then
    echo "   Last 10 SSH-related log entries:"
    tail -10 /var/log/auth.log | grep -i ssh
elif [ -f "/var/log/secure" ]; then
    echo "   Last 10 SSH-related log entries:"
    tail -10 /var/log/secure | grep -i ssh
else
    echo "   No SSH log files found"
fi
echo ""

echo "=== Verification Complete ==="
echo ""
echo "If you need to add your public key, run:"
echo "echo 'ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDe0IvBKlcPue/VgOO9sIecq62PosnYbV9uydR2hzFMLHQChHQQkkNpUL0olI3p0OOOBBAXNG44E0gDU/UP+zU48Mx7ct+Bv+83YrFTwM3Tx26rfRtIvhMZmHgJJjuAp/OtFPG3d2oM00XZ7wdyg+Bx0YRORY/hQ5+cal+LUkCjg0vvYWkk53th9n/VGpIp65Vxn7SiayGOCFrTaj93ZSXruhHfdi5kGUNYo/FJ7q8xKTexo/ccl5wQIA2hdOZfcmdsbmcDEpMnBZRu8EoeHngRrlzAxGiOHIDFZKJKuSxWTRd5zOOSEbX2YEQaQ72PucKwIrvMpYQhUFI6amgAie4VuU/DiqGUgxtIVOdK4VBxcDdSS8oBpZXtUjZVojeiaKICyFP4S8841VYeJ8H+llyMNHlztTOFs2Xm90QupSn2i1jHJvrbQCM2uAR0yjyY8S6pRwqqaBvYzZAASXDPO8p4rki5l82SMzGFU/bWs1OAx/eEGzrUgWojSnG/YGNKurQCyy/l/21Z/0P+hzAJs0k499S74/1X+3BL4W9wIvwyPzNrX6Xs5PD8wsA9gJm/14qCdc5cQ7jjtGcsiagKPmQMKoMI3b+caWY0OvijpQvralJ4IIX/iaC5qzEFhCd4vS5tsc0hM/Dkoexabkzwk5DQcn2SvnBe6Lh6OijJihDq+Q== root@ubuntu-s-1vcpu-1gb-35gb-intel-nyc1-01' >> /root/.ssh/authorized_keys"
echo ""
echo "Then set correct permissions:"
echo "chmod 600 /root/.ssh/authorized_keys"
echo "chmod 700 /root/.ssh" 