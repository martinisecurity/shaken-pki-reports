# STIR/SHAKEN Certificate Repository Compliance

## Five9 Inc

Name: `https://cr-uat.ccid.neustar.biz/ccid/authn/v2/certs/11363.10160.pem`\
Tested At: 22 Nov 23 03:23 UTC\
Time: 95ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Nov 23 03:38 UTC