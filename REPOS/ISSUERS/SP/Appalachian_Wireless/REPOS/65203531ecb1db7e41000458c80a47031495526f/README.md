# STIR/SHAKEN Certificate Repository Compliance

## Appalachian Wireless

Name: `https://cdn-cr.cgah.tnsi.com/certs/590f3a17396e9e874d158eb60337f24c3ac4812b`\
Tested At: 28 Apr 23 02:03 UTC\
Time: 1ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 28 Apr 23 02:17 UTC