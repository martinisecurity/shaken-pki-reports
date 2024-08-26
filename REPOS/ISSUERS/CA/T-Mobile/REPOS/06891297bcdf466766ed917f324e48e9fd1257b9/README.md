# STIR/SHAKEN Certificate Repository Compliance

## T-Mobile

Name: `https://t-mobile-sticr.fosrvt.com/8814f66226a3d07edcffb9cfb333c423bfe5dd0f76dfda2839df9004f46fe86e.pem`\
Tested At: 26 Aug 24 18:46 UTC\
Time: 64ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:49 UTC