# STIR/SHAKEN Certificate Repository Compliance

## Neustar

Name: `https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/146.147`\
Tested At: 06 Nov 22 17:26 UTC\
Time: 21ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 06 Nov 22 17:28 UTC