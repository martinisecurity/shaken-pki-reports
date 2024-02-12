# STIR/SHAKEN Certificate Repository Compliance

## AcmeTelecom, Inc.

Name: `https://65.108.80.93/cert.pem`\
Tested At: 12 Feb 24 18:53 UTC\
Time: 348ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://65.108.80.93/cert.pem": x509: cannot validate certificate for 65.108.80.93 because it doesn't contain any IP SANs |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 19:26 UTC