# STIR/SHAKEN Certificate Repository Compliance

## Highland Telephone Cooperative

Name: `https://cdn-cr.cgah.tnsi.com/certs/2ea4f70250c45d8d12df4852263f77d007152f19`\
Tested At: 04 Oct 24 15:31 UTC\
Time: 61ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC