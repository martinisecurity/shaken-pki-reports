# STIR/SHAKEN Certificate Repository Compliance

## Latam Telecommunications LLC

Name: `https://187.174.67.118:8080/0e4158a72f7ae0b175762c8b52a0b357.cer`\
Tested At: 04 Oct 24 15:30 UTC\
Time: 135ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://187.174.67.118:8080/0e4158a72f7ae0b175762c8b52a0b357.cer": x509: certificate has expired or is not yet valid: current time 2024-10-04T15:30:56Z is after 2024-04-01T12:18:54Z |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |
| [w_atis_protocol](../../ISSUES/w_atis_protocol/README.md) | warn | ATIS-1000080 | The verifier should not dereference any protocol other than https or a port other than 443 or 8443 |

Generated: 04 Oct 24 16:29 UTC