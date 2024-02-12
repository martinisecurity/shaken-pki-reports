# STIR/SHAKEN Certificate Repository Compliance

## Commio

Name: `https://storage.googleapis.com/stirshaken/ShakeNBakeCert.pem`\
Tested At: 12 Feb 24 16:58 UTC\
Time: 58ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The expiration time is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 17:02 UTC