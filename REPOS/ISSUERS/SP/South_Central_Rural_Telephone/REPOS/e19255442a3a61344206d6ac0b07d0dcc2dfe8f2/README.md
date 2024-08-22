# STIR/SHAKEN Certificate Repository Compliance

## South Central Rural Telephone

Name: `https://cdn-cr.cgah.tnsi.com/certs/448db51c0154c1b4aa55612c8779cede93963e92`\
Tested At: 22 Aug 24 15:18 UTC\
Time: 25ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 16:06 UTC