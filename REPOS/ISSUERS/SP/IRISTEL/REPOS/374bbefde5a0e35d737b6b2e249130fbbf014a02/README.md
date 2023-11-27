# STIR/SHAKEN Certificate Repository Compliance

## IRISTEL

Name: `https://sticr-cstga.ccid.neustar/api/v1/certificate/a93fb04b9d4d2e9384352d32f7830cb2.crt`\
Tested At: 27 Nov 23 23:24 UTC\
Time: 513ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 27 Nov 23 23:28 UTC