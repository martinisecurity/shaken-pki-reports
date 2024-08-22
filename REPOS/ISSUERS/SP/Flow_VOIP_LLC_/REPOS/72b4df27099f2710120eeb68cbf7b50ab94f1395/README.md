# STIR/SHAKEN Certificate Repository Compliance

## Flow VOIP LLC.

Name: `https://cdn.cnxcdn.com/shaken/2cee92bfa4.crt`\
Tested At: 22 Aug 24 15:16 UTC\
Time: 31ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 15:44 UTC