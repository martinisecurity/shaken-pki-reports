# STIR/SHAKEN Certificate Repository Compliance

## XCast Labs

Name: `https://cr.xcastlabs.net/1668234600/xclsshaken.pem`\
Tested At: 01 Nov 22 16:04 UTC\
Time: 2ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 01/11/2022 at 16:11:23