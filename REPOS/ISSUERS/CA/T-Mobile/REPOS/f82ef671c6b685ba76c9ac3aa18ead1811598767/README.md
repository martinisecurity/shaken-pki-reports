# STIR/SHAKEN Certificate Repository Compliance

## T-Mobile

Name: `https://t-mobile-sticr.fosrvt.com/f59530d22b839d3b50a91a6279e8b1e0510cade2b8913ff3011c89e3f3725300.pem`\
Tested At: 21 Nov 23 17:11 UTC\
Time: 17ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Nov 23 17:16 UTC