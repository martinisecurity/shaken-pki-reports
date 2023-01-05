# STIR/SHAKEN Certificate Repository Compliance

## Microtalk USA, Inc.

Name: `https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem`\
Tested At: 05 Jan 23 20:55 UTC\
Time: 5325ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 05 Jan 23 21:05 UTC