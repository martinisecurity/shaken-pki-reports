# STIR/SHAKEN Certificate Repository Compliance

## AcmeTelecom, Inc.

Name: `https://65.108.80.93/cert.pem`\
Tested At: 18 Aug 25 20:04 UTC\
Time: 356ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://65.108.80.93/cert.pem": x509: certificate has expired or is not yet valid: current time 2025-08-18T20:04:39Z is after 2024-09-12T19:30:51Z |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC