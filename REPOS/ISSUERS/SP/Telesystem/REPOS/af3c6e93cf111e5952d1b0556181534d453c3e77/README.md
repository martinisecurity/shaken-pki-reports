# STIR/SHAKEN Certificate Repository Compliance

## Telesystem

Name: `https://cdn-cr.cgah.tnsi.com/certs/9a9d047943fd910a6df566d57af59a46a867276e`\
Tested At: 18 Aug 25 20:05 UTC\
Time: 21ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC