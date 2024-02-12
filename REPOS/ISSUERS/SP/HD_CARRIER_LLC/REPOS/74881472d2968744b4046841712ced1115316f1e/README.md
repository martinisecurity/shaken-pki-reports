# STIR/SHAKEN Certificate Repository Compliance

## HD CARRIER LLC

Name: `https://d64db847f381fcb974ab1b6150b49a91.cxstatic.com/f/20f8bca4-d82b-49d7-b419-ff8a42d2dfbb`\
Tested At: 12 Feb 24 19:17 UTC\
Time: 326ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header doesn't have 'max-age' directive |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 19:26 UTC