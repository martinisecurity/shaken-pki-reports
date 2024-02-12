# STIR/SHAKEN Certificate Repository Compliance

## HD CARRIER LLC

Name: `https://d64db847f381fcb974ab1b6150b49a91.cxstatic.com/f/c09d4283-3f57-407d-bbfd-e909046890ed`\
Tested At: 12 Feb 24 16:53 UTC\
Time: 277ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header doesn't have 'max-age' directive |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 17:02 UTC