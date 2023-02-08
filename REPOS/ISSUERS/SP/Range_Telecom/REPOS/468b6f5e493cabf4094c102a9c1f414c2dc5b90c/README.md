# STIR/SHAKEN Certificate Repository Compliance

## Range Telecom

Name: `https://ssc.getsipnav.com/certs/907a99e8008aa590695697cee73d401d4ef33451`\
Tested At: 08 Feb 23 19:36 UTC\
Time: 27ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 08 Feb 23 19:45 UTC