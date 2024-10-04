# STIR/SHAKEN Certificate Repository Compliance

## WOW Internet Cable and Phone

Name: `https://cdn-cr.cgah.tnsi.com/certs/27c187b748655ad5d9d3a570aefb8db66d57f3eb`\
Tested At: 04 Oct 24 15:31 UTC\
Time: 60ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC