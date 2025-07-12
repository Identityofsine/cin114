# Secrets Management Guide

This document outlines how to properly manage secrets and sensitive configuration in the CIN114 project.

## Required Secrets

### Stripe Configuration

| Environment Variable | Description | Development | Production |
|---------------------|-------------|-------------|------------|
| `STRIPE_SECRET_KEY` | Stripe API secret key | `sk_test_*` | `sk_live_*` |
| `STRIPE_WEBHOOK_SECRET` | Stripe webhook endpoint secret | `whsec_*` | `whsec_*` |
| `STRIPE_REDIRECT_URL` | URL to redirect after payment | `https://dev.cin114.net` | `https://cin114.net` |

### Authentication Secrets

| Environment Variable | Description | Required |
|---------------------|-------------|----------|
| `JWT_SECRET` | JWT signing secret | Yes |
| `GOOGLE_CLIENT_ID` | Google OAuth client ID | Optional |
| `GOOGLE_CLIENT_SECRET` | Google OAuth client secret | Optional |

## Environment-Specific Configuration

### Development Environment

For local development, you can use placeholder values or test keys:

```bash
# .env file for development
STRIPE_SECRET_KEY=sk_test_placeholder
STRIPE_WEBHOOK_SECRET=whsec_placeholder
STRIPE_REDIRECT_URL=https://dev.cin114.net
JWT_SECRET=dev-jwt-secret
```

### Production Environment

For production, you must use real secrets:

```bash
# Set these as environment variables or in your deployment system
STRIPE_SECRET_KEY=sk_live_your_actual_live_key
STRIPE_WEBHOOK_SECRET=whsec_your_actual_webhook_secret
STRIPE_REDIRECT_URL=https://cin114.net
JWT_SECRET=your_secure_jwt_secret
```

## GitHub Actions Secrets

Add these secrets to your GitHub repository:

1. Go to your repository settings
2. Navigate to "Secrets and variables" → "Actions"
3. Add the following repository secrets:

| Secret Name | Description |
|-------------|-------------|
| `STRIPE_SECRET_KEY` | Your Stripe live secret key |
| `STRIPE_WEBHOOK_SECRET` | Your Stripe webhook secret |
| `SERVER_IP` | Your deployment server IP |
| `SERVER_USER` | SSH username for deployment |
| `SSH_PRIVATE_KEY` | SSH private key for deployment |

## Local Development Setup

1. Create a `.env` file in the project root:

```bash
# Copy the example and modify as needed
cp .env.example .env
```

2. Set your development secrets:

```bash
# Development secrets
STRIPE_SECRET_KEY=sk_test_your_test_key
STRIPE_WEBHOOK_SECRET=whsec_your_test_webhook_secret
STRIPE_REDIRECT_URL=https://dev.cin114.net
JWT_SECRET=dev-jwt-secret-change-this
```

## Validation

The application includes secret validation that will:

- Check that Stripe keys have the correct prefix (`sk_test_` for dev, `sk_live_` for prod)
- Verify webhook secrets have the correct prefix (`whsec_`)
- Warn about missing required secrets in production
- Allow placeholder values in development

## Security Best Practices

1. **Never commit secrets to version control**
   - Use `.env` files (already in `.gitignore`)
   - Use environment variables in production
   - Use GitHub Secrets for CI/CD

2. **Use different keys for different environments**
   - Test keys for development
   - Live keys for production

3. **Rotate secrets regularly**
   - Update JWT secrets periodically
   - Rotate Stripe keys if compromised

4. **Monitor secret usage**
   - Check application logs for secret validation warnings
   - Monitor Stripe dashboard for unusual activity

## Troubleshooting

### Common Issues

1. **"Stripe configuration errors" warning**
   - Check that your secrets are properly set
   - Verify key prefixes are correct
   - Ensure required secrets are present in production

2. **Webhook signature verification fails**
   - Verify `STRIPE_WEBHOOK_SECRET` is correct
   - Check that webhook endpoint URL matches your configuration

3. **Payment processing fails**
   - Verify `STRIPE_SECRET_KEY` is correct
   - Check that you're using the right keys for your environment

### Debugging

To debug secret configuration:

```bash
# Check environment variables
echo $STRIPE_SECRET_KEY
echo $STRIPE_WEBHOOK_SECRET

# Run with debug logging
GO_ENV=development go run cmd/server/main.go
```

## Getting Stripe Keys

1. **Test Keys**: Available in your Stripe Dashboard → Developers → API Keys
2. **Live Keys**: Available in your Stripe Dashboard → Developers → API Keys (switch to "Live" mode)
3. **Webhook Secrets**: Create a webhook endpoint in Stripe Dashboard → Developers → Webhooks

## Support

If you encounter issues with secret configuration:

1. Check the application logs for validation warnings
2. Verify your environment variables are set correctly
3. Ensure you're using the right keys for your environment
 