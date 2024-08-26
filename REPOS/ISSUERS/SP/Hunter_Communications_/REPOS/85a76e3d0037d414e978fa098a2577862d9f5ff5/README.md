# STIR/SHAKEN Certificate Repository Compliance

## Hunter Communications 

Name: `https://cdn-cr.cgah.tnsi.com/certs/7ac98ce362be7e0d6650ea43e17a8320e956b66a`\
Tested At: 26 Aug 24 17:42 UTC\
Time: 9ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 26 Aug 24 18:03 UTC