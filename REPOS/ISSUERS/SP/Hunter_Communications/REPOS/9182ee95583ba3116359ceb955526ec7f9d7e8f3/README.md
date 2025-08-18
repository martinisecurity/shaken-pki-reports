# STIR/SHAKEN Certificate Repository Compliance

## Hunter Communications

Name: `https://cdn-cr.cgah.tnsi.com/certs/7fba9a7fc8b3131f6fcea50e668939fd26bbd4a3`\
Tested At: 18 Aug 25 20:05 UTC\
Time: 23ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC