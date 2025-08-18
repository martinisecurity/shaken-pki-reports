# STIR/SHAKEN Certificate Repository Compliance

## FracTEL LLC

Name: `https://ssc.fractel.net/ssc/fractelssc.pem`\
Tested At: 18 Aug 25 21:05 UTC\
Time: 67ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [e_tls_transport](../../ISSUES/e_tls_transport/README.md) | error | System | Get "https://ssc.fractel.net/ssc/fractelssc.pem": x509: certificate has expired or is not yet valid: current time 2025-08-18T21:05:42Z is after 2025-08-16T03:05:32Z |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC