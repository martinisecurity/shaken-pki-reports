# STIR/SHAKEN Certificate Repository Compliance

## FracTEL LLC

Name: `https://ssc.fractel.net/ssc/fractelssc.pem`\
Tested At: 26 Aug 24 18:42 UTC\
Time: 182ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://ssc.fractel.net/ssc/fractelssc.pem": x509: certificate has expired or is not yet valid: current time 2024-08-26T18:42:30Z is after 2024-08-20T19:29:08Z |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:49 UTC