# STIR/SHAKEN Certificate Repository Compliance

## Telcast Networks

Name: `https://ssc.getsipnav.com/certs/724dcfb47a877f9ba8cf306376dfde026b11b227`\
Tested At: 26 Aug 24 17:57 UTC\
Time: 37ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:03 UTC