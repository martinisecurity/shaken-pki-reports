# STIR/SHAKEN Certificate Repository Compliance

## Vonage

Name: `https://sticr-cstga.ccid.neustar/api/v1/certificate/3f7be0c667b2f3c24546cf26629eb906.crt`\
Tested At: 01 Dec 22 19:21 UTC\
Time: 671ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 01 Dec 22 19:22 UTC