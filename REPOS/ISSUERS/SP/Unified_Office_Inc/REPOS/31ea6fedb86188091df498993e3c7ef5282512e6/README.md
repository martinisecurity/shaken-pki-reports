# STIR/SHAKEN Certificate Repository Compliance

## Unified Office Inc

Name: `https://downloads.uotcn.net/certs/uo-shaken-cert-20230222.pem`\
Tested At: 28 Nov 23 20:12 UTC\
Time: 184ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The expiration time is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 28 Nov 23 20:21 UTC