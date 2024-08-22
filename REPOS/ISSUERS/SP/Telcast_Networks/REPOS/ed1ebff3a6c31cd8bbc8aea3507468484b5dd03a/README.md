# STIR/SHAKEN Certificate Repository Compliance

## Telcast Networks

Name: `https://ssc.getsipnav.com/certs/724dcfb47a877f9ba8cf306376dfde026b11b227`\
Tested At: 22 Aug 24 15:37 UTC\
Time: 53ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 15:44 UTC