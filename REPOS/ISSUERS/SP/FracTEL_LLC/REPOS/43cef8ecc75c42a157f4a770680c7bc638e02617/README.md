# STIR/SHAKEN Certificate Repository Compliance

## FracTEL LLC

Name: `https://ssc.fractel.net/ssc/fractelssc.pem`\
Tested At: 04 Oct 24 16:23 UTC\
Time: 39ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://ssc.fractel.net/ssc/fractelssc.pem": x509: certificate has expired or is not yet valid: current time 2024-10-04T16:23:37Z is after 2024-08-20T19:29:08Z |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC