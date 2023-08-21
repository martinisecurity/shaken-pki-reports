# STIR/SHAKEN Certificate Repository Compliance

## Avid Telecom LLC

Name: `https://ssc.getsipnav.com/certs/04f6fca4ed2cb2ba3395ab5c4de7d4573e3e6c54`\
Tested At: 21 Aug 23 20:16 UTC\
Time: 161ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Aug 23 20:18 UTC