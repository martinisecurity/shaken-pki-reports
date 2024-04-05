# STIR/SHAKEN Certificate Repository Compliance

## Yelcot

Name: `https://cdn-cr.cgah.tnsi.com/certs/02f31fd9be31f3b5cc9fa396dc9fab04f1dc93e1`\
Tested At: 05 Apr 24 18:39 UTC\
Time: 132ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 05 Apr 24 19:04 UTC