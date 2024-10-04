# STIR/SHAKEN Certificate Repository Compliance

## Altice USA

Name: `https://cr.ccid.alticeusa.com/ccid/authn/v2/certs/11011.10003`\
Tested At: 04 Oct 24 15:51 UTC\
Time: 165ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://cr.ccid.alticeusa.com/ccid/authn/v2/certs/11011.10003": x509: certificate has expired or is not yet valid: current time 2024-10-04T15:51:15Z is after 2024-09-21T23:59:59Z |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC