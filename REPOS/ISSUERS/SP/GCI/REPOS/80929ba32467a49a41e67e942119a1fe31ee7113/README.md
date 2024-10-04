# STIR/SHAKEN Certificate Repository Compliance

## GCI

Name: `https://cdn-cr.cgah.tnsi.com/certs/69fae8e574f06fb796e08c7b6ea847f2563cea5b`\
Tested At: 04 Oct 24 15:31 UTC\
Time: 76ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC