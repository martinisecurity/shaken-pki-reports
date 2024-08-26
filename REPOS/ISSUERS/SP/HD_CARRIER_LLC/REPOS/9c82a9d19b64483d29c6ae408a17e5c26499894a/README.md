# STIR/SHAKEN Certificate Repository Compliance

## HD CARRIER LLC

Name: `https://d64db847f381fcb974ab1b6150b49a91.cxstatic.com/f/13cac9d6-ba52-4064-9a8e-e36c9fe52208`\
Tested At: 26 Aug 24 17:49 UTC\
Time: 80ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header doesn't have 'max-age' directive |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:03 UTC