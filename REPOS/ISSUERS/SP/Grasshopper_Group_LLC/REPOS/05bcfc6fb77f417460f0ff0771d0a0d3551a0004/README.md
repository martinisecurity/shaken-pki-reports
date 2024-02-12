# STIR/SHAKEN Certificate Repository Compliance

## Grasshopper Group LLC

Name: `https://pstn-cdn.live.gtc.goto.com/certs/stirshaken/gh-2023-11`\
Tested At: 12 Feb 24 16:55 UTC\
Time: 14ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 17:02 UTC